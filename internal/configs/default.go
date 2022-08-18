package configs

import "github.com/HighwayToVictory/g-router/internal/router"

func Default() Config {
	return Config{
		Router: router.Config{
			DefaultGateWay:    "0",
			NumberOfInterface: 10,
		},
	}
}
