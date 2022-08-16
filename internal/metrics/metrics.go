package metrics

import (
	"net/http"
	"sync"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewMetrics(wg *sync.WaitGroup) {
	go func() {
		wg.Add(1)

		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":4040", nil); err != nil {
			panic(err)
		}

		wg.Done()
	}()
}
