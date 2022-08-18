package configs

import "github.com/HighwayToVictory/g-router/internal/router"

type Config struct {
	Router router.Config
}

func Load() Config {
	return Default()
}
