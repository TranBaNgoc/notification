package controllers

import (
	"runtime"
	"time"
)

func init() {

}

func DefaultConfig() *Config {
	return &Config{
		Host:        "0.0.0.0",
		Port:        "8081",
		LogLevel:    "info",
		Cache:       "memory",
		Android:     map[string]string{"APIKey": "AAAAwud4BT8:APA91bEVAUa9ksGKw3E2NAYVzCJEngCcWAAyNnqydzhfB_FnZPgVUaagpHy0Oog4dxUrklvmKtHvA7-VoDEENhGzMJeGWOwBddLM7bKF6u2ORCCtNMzJ5aIUrQaamlqktQ4s0lGhIV3E"},
		Ios:         map[string]string{"certificateValidP12": "", },
		WorkerNum:   runtime.NumCPU(),
		QueueNum:    8192,
		Sync:        false,
		GracePeriod: 5 * time.Second,
	}
}

const envPrefix = "notify-config"

// Config for the loginsrv handler
type Config struct {
	Host        string
	Port        string
	LogLevel    string
	Cache       string //"memory" # support memory, redis
	Redis       map[string]string
	Android     map[string]string
	Ios         map[string]string
	WorkerNum   int
	QueueNum    uint
	Sync        bool
	GracePeriod time.Duration
}
