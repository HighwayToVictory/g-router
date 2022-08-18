package cmd

import (
	"github.com/HighwayToVictory/g-router/internal/configs"
	"github.com/HighwayToVictory/g-router/internal/metrics"
	"github.com/HighwayToVictory/g-router/internal/router"
)

func Execute() {
	cfg := configs.Load()

	go router.NewProvider(cfg.Router)

	metrics.Register()
}
