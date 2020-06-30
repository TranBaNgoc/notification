package controllers

import (
	"github.com/micro/go-micro/v2/api"
	apirpc "github.com/micro/go-micro/v2/api/handler/rpc"
	"github.com/micro/go-micro/v2/api/router"
	rstatic "github.com/micro/go-micro/v2/api/router/static"
	"github.com/micro/go-micro/v2/registry"
	"net/http"
	"github.com/TranBaNgoc/notification-liveshopping/logging"
)

const EndpointService = "/notify"

type ApiRouter interface {
	Register(ep *api.Endpoint) error
	Deregister(ep *api.Endpoint) error
	Options() router.Options
	Close() error
	Endpoint(req *http.Request) (*api.Service, error)
	Route(req *http.Request) (*api.Service, error)
}

func NewRouter(serviceName string, registry registry.Registry) ApiRouter {
	apiRouter := rstatic.NewRouter(
		router.WithHandler(apirpc.Handler),
		router.WithRegistry(registry),
	)
	//add endpoint
	err := apiRouter.Register(&api.Endpoint{
		// rpc dinh nghia trong file .proto
		Name:    serviceName + ".Notify.Send",
		Method:  []string{"POST"},
		Path:    []string{EndpointService + "/sent"},
		Handler: "rpc",
	})

	if err != nil {
		logging.Logger.Fatal(err)
	}

	err = apiRouter.Register(&api.Endpoint{
		// rpc dinh nghia trong file .proto
		Name:    serviceName + ".Health.Check",
		Method:  []string{"GET"},
		Path:    []string{EndpointService + "/check"},
		Handler: "rpc",
	})

	if err != nil {
		logging.Logger.Fatal(err)
	}
	return apiRouter
}
