package prometheus

import (
	"orchid-starter/internal/common"

	prometheusApplication "github.com/prometheus/client_golang/prometheus"
)

var (
	HttpRequestsTotal = prometheusApplication.NewCounterVec(
		prometheusApplication.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
			ConstLabels: prometheusApplication.Labels{
				"app_name": common.GetEnvWithDefault("APP_NAME", "orchid-starter"),
			},
		},
		[]string{"path", "method", "status"},
	)

	RequestDuration = prometheusApplication.NewHistogramVec(
		prometheusApplication.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of request durations",
			Buckets: prometheusApplication.DefBuckets,
			ConstLabels: prometheusApplication.Labels{
				"app_name": common.GetEnvWithDefault("APP_NAME", "orchid-starter"),
			},
		},
		[]string{"path", "method"},
	)

	InProgressRequests = prometheusApplication.NewGauge(
		prometheusApplication.GaugeOpts{
			Name: "http_in_progress_requests",
			Help: "Number of HTTP requests currently being processed",
			ConstLabels: prometheusApplication.Labels{
				"app_name": common.GetEnvWithDefault("APP_NAME", "orchid-starter"),
			},
		},
	)
)

func InitPrometheus() {
	if !common.GetBoolEnv("PROMETHEUS_ACTIVE", false) {
		return
	}

	prometheusApplication.MustRegister(HttpRequestsTotal)
	prometheusApplication.MustRegister(RequestDuration)
	prometheusApplication.MustRegister(InProgressRequests)
}
