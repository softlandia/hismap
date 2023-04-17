package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	namespace = "hismap"
	subsystem = "server"
)

var (
	HttpCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   subsystem,
			Name:        "http_requests_total",
			Help:        "The total number of http requests",
			ConstLabels: nil,
		},
		[]string{
			"code",
			"method",
			"handler",
		},
	)
	HttpDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "http_duration",
			Help:      "Duration of http requests",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{
			"code",
			"method",
			"handler",
		},
	)
	ErrorCollector = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   namespace,
			Subsystem:   subsystem,
			Name:        "err_collector",
			Help:        "The total number of err while service is working",
			ConstLabels: nil,
		},
		[]string{"method", "call"},
	)
)
