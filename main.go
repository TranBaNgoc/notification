package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	PushConf = DefaultConfig()
	ct := context.Background()
	wg := &sync.WaitGroup{}
	wg.Add(PushConf.WorkerNum)
	InitWorkers(ct, wg, int64(PushConf.WorkerNum), int64(PushConf.QueueNum))

	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	mux.HandleFunc("/follow", FollowNotify)
	fmt.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(":"+port, mux))
}

func FollowNotify(writer http.ResponseWriter, request *http.Request) {
	user_id, ok := request.URL.Query()["user_id"]
	if !ok {
		log.Println("user_id is not param query")
	} else {
		req := PushNotification{}
		req.Tokens = GetUserTokens(user_id[0])
		req.Message = "This is a follow"
		req.Platform = 2
		success := PushToAndroid(req)
		if success {
			log.Println("Send message fail")
			writer.WriteHeader(400)
			writer.Write([]byte("Send message fail"))
			return
		}
		log.Println("Send message success")
		writer.WriteHeader(200)
		writer.Write([]byte("Send message success"))

	}
}
