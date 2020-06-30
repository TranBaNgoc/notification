package controllers

import (
	"github.com/appleboy/go-fcm"
	"github.com/sideshow/apns2"
)

var (
	//golbal config
	PushConf *Config
	// QueueNotification is chan type
	QueueNotification chan PushNotification
	// ApnsClient is apns client
	ApnsClient *apns2.Client
	// FCMClient is apns client
	FCMClient *fcm.Client
	// MaxConcurrentIOSPushes pool to limit the number of concurrent iOS pushes
	MaxConcurrentIOSPushes chan struct{}
)
