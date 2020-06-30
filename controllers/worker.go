package controllers

import (
	"context"
	"github.com/TranBaNgoc/notification-liveshopping/logging"
	"sync"
)

// InitWorkers for initialize all workers.
func InitWorkers(ctx context.Context, wg *sync.WaitGroup, workerNum int64, queueNum int64) {
	logging.Logger.Info("worker number is ", workerNum, ", queue number is ", queueNum)
	QueueNotification = make(chan PushNotification, queueNum)
	for i := int64(0); i < workerNum; i++ {
		go startWorker(wg, i)
	}
}

// SendNotification is send message to iOS or Android
func SendNotification(req PushNotification) {
	if PushConf.Sync {
		defer req.WaitDone()
	}

	switch req.Platform {
	case PlatFormIos:
		PushToIOS(req)
	case PlatFormAndroid:
		PushToAndroid(req)
	}
}

func startWorker(wg *sync.WaitGroup, num int64) {
	defer wg.Done()
	for notification := range QueueNotification {
		SendNotification(notification)
	}
	logging.Logger.Info("closed the worker num ", num)
}

// markFailedNotification adds failure logs for all tokens in push notification
func markFailedNotification(notification *PushNotification, reason string) {
	logging.Logger.Error(reason)
	/*for _, token := range notification.Tokens {
		notification.AddLog(getLogPushEntry(FailedPush, token, *notification, errors.New(reason)))
	}*/
	notification.WaitDone()
}

// queueNotification add notification to queue list.
func AddNotification(notification PushNotification) int {
	var count int
	wg := sync.WaitGroup{}

	//log := make([]LogPushEntry, 0, count)

	if PushConf.Sync {
		notification.wg = &wg
		//notification.log = &log
		notification.AddWaitCount()
	}
	if !tryEnqueue(notification, QueueNotification) {
		markFailedNotification(&notification, "max capacity reached")
	}
	count += len(notification.Tokens)
	// Count topic message
	if notification.To != "" {
		count++
	}

	if PushConf.Sync {
		wg.Wait()
	}

	//StatStorage.AddTotalCount(int64(count))

	return count
}

// tryEnqueue tries to enqueue a job to the given job channel. Returns true if
// the operation was successful, and false if enqueuing would not have been
// possible without blocking. Job is not enqueued in the latter case.
func tryEnqueue(job PushNotification, jobChan chan<- PushNotification) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}
