package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewMetrics() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
	}()
}
