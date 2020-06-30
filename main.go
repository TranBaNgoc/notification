package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	ahandler "github.com/micro/go-micro/v2/api/handler"
	apirpc "github.com/micro/go-micro/v2/api/handler/rpc"
	"github.com/micro/go-micro/v2/server"
	gsrv "github.com/micro/go-micro/v2/server/grpc"
	tgrpc "github.com/micro/go-micro/v2/transport/grpc"
	"net/http"
	"notification-liveshopping/controllers"
	"notification-liveshopping/logging"
	notify "notification-liveshopping/proto/notification"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const serviceName = "notifysrv"

func main() {
	controllers.PushConf = controllers.DefaultConfig()
	if err := logging.Set(controllers.PushConf.LogLevel, true); err != nil {
		exit(nil, err)
	}
	controllers.PushConf.Port = os.Getenv("PORT")

	//Load config OK --> start services
	logging.LifecycleStart(serviceName, controllers.PushConf)

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// khoi tao Transport
	tr := tgrpc.NewTransport()
	// Cau hinh Registry
	// tao server gprc
	svr := gsrv.NewServer(
		server.Name(serviceName),
		server.Version("latest"),
		server.Transport(tr),
		server.Address("0.0.0.0:8080"),
	)

	// khoi tao Service su dung server RPC
	service := micro.NewService(
		micro.Server(svr),
		micro.Transport(tr),
		micro.WrapHandler(),
	)

	// Register Handler rpc
	h, err := controllers.NewHandler(controllers.PushConf)
	if err != nil {
		exit(nil, err)
	}

	// Init worker
	ctx := context.Background()
	wg := &sync.WaitGroup{}
	wg.Add(int(controllers.PushConf.WorkerNum))
	controllers.InitWorkers(ctx, wg, int64(controllers.PushConf.WorkerNum), int64(controllers.PushConf.QueueNum))

	notify.RegisterNotifyHandler(svr, h)
	notify.RegisterHealthHandler(svr, h)
	defer svr.Stop()

	// Run service tren goroutine
	go func() {
		if err := service.Run(); err != nil {
			exit(nil, err)
		}

	}()

	// check registration ok ?
	time.Sleep(1 * time.Second)
	// khoi tao Router
	apiRouter := controllers.NewRouter(serviceName, service.Server().Options().Registry)
	// khoi tao API Handler
	hrpc := apirpc.NewHandler(ahandler.WithRouter(apiRouter))

	//setup HTTP server wrap len API handler
	logging.Logger.Info("Listen HTTP in port:", controllers.PushConf.Port)
	hsrv := &http.Server{
		Addr:           "0.0.0.0:" + controllers.PushConf.Port,
		Handler:        hrpc,
		WriteTimeout:   150 * time.Second,
		ReadTimeout:    150 * time.Second,
		IdleTimeout:    200 * time.Second,
		MaxHeaderBytes: 1024 * 1024 * 1, // 1Mb
	}
	//Listen API
	go func() {
		if err := hsrv.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				logging.ServerClosed(serviceName)
			} else {
				exit(nil, err)
			}
		}
	}()
	//Stop
	logging.LifecycleStop(serviceName, <-stop, nil)

	ctx, cancel := context.WithTimeout(context.Background(), controllers.PushConf.GracePeriod)
	hsrv.Shutdown(ctx)
	cancel()
}

var exit = func(signal os.Signal, err error) {
	logging.LifecycleStop(serviceName, signal, err)
	if err == nil {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}