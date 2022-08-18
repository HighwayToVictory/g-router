package metrics

import (
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "g-router"
)

type Metrics struct {
	RequestCounter prometheus.Counter
	GatewayCounter prometheus.Counter
}

func NewMetrics() Metrics {
	return Metrics{
		RequestCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "request_counter",
			Help:      "number of requests to router",
		}),
		GatewayCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "gateway_counter",
			Help:      "number of routes that hit gateway",
		}),
	}
}

func Register(wg *sync.WaitGroup) {
	go func() {
		wg.Add(1)

		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":4040", nil); err != nil {
			panic(err)
		}

		wg.Done()
	}()
}
