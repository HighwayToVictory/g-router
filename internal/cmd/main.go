package cmd

import (
	"fmt"

	"github.com/HighwayToVictory/g-router/internal/metrics"
)

func Execute() {
	metrics.NewMetrics()

	fmt.Println("vim-go")
}
