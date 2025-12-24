package middleware

import (
	"orchid-starter/internal/common"
	"orchid-starter/observability/prometheus"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

func Prometheus(irisCtx iris.Context) {
	if !common.GetBoolEnv("PROMETHEUS_ACTIVE", false) {
		irisCtx.Next()
		return
	}

	if irisCtx.Path() == "/metrics" {
		irisCtx.Next()
		return
	}

	start := time.Now()
	prometheus.InProgressRequests.Inc()
	defer prometheus.InProgressRequests.Dec()

	irisCtx.Next()
	// Increment total request counter
	prometheus.HttpRequestsTotal.WithLabelValues(irisCtx.Path(), irisCtx.Method(), strconv.Itoa(irisCtx.GetStatusCode())).Inc()
	// Observe duration
	prometheus.RequestDuration.WithLabelValues(irisCtx.Path(), irisCtx.Method()).Observe(time.Since(start).Seconds())
}
