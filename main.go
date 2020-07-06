package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	PushConf = DefaultConfig()
	ct := context.Background()
	wg := &sync.WaitGroup{}
	wg.Add(PushConf.WorkerNum)
	InitWorkers(ct, wg, int64(PushConf.WorkerNum), int64(PushConf.QueueNum))


	mux := http.NewServeMux()
	mux.HandleFunc("/send", SendNotify)

	log.Fatalln(http.ListenAndServe(":8899", mux))
}

func SendNotify(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
		return
	}

	req := PushNotification{}
	json.Unmarshal(body, &req)


	success := PushToAndroid(req)
	if !success {
		log.Println("Send message fail")
		writer.WriteHeader(400)
		writer.Write([]byte("Send message fail"))
	} else {
		log.Println("Send message success")
		writer.WriteHeader(200)
		writer.Write([]byte("Send message success"))
	}

}
