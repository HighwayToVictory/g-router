package cmd

import (
	"sync"

	"github.com/HighwayToVictory/g-router/internal/configs"
	"github.com/HighwayToVictory/g-router/internal/metrics"
	"github.com/HighwayToVictory/g-router/internal/router"
)

func Execute() {
	wg := sync.WaitGroup{}
	cfg := configs.Load()

	rt := router.New(cfg.Router)

	metrics.Register(&wg)

	wg.Wait()
}
