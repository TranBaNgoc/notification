package controllers

import (
	"github.com/prometheus/client_golang/prometheus"
)

const namespace = "gorush_"

// Metrics implements the prometheus.Metrics interface and
// exposes gorush metrics for prometheus
type Metrics struct {
	TotalPushCount *prometheus.Desc
	IosSuccess     *prometheus.Desc
	IosError       *prometheus.Desc
	AndroidSuccess *prometheus.Desc
	AndroidError   *prometheus.Desc
	QueueUsage     *prometheus.Desc
}

// NewMetrics returns a new Metrics with all prometheus.Desc initialized
func NewMetrics() Metrics {

	return Metrics{
		TotalPushCount: prometheus.NewDesc(
			namespace+"total_push_count",
			"Number of push count",
			nil, nil,
		),
		IosSuccess: prometheus.NewDesc(
			namespace+"ios_success",
			"Number of iOS success count",
			nil, nil,
		),
		IosError: prometheus.NewDesc(
			namespace+"ios_error",
			"Number of iOS fail count",
			nil, nil,
		),
		AndroidSuccess: prometheus.NewDesc(
			namespace+"android_success",
			"Number of android success count",
			nil, nil,
		),
		AndroidError: prometheus.NewDesc(
			namespace+"android_fail",
			"Number of android fail count",
			nil, nil,
		),
		QueueUsage: prometheus.NewDesc(
			namespace+"queue_usage",
			"Length of internal queue",
			nil, nil,
		),
	}
}

// Describe returns all possible prometheus.Desc
func (c Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.TotalPushCount
	ch <- c.IosSuccess
	ch <- c.IosError
	ch <- c.AndroidSuccess
	ch <- c.AndroidError
	ch <- c.QueueUsage
}

// Collect returns the metrics with values
func (c Metrics) Collect(ch chan<- prometheus.Metric) {

}
