package cmd

import (
	"sync"

	"github.com/HighwayToVictory/g-router/internal/metrics"
)

func Execute() {
	wg := sync.WaitGroup{}

	metrics.NewMetrics(&wg)

	wg.Wait()
}
