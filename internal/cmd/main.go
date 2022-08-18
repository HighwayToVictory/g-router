package cmd

import (
	"sync"

	"github.com/HighwayToVictory/g-router/internal/metrics"
)

func Execute() {
	wg := sync.WaitGroup{}

	metrics.Register(&wg)

	wg.Wait()
}
