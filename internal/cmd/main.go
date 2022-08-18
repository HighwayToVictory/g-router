package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/HighwayToVictory/g-router/internal/configs"
	"github.com/HighwayToVictory/g-router/internal/metrics"
	"github.com/HighwayToVictory/g-router/internal/router"
)

func Execute() {
	wg := sync.WaitGroup{}
	cfg := configs.Load()

	go putOnRouter(cfg.Router)

	metrics.Register(&wg)

	wg.Wait()
}

func randomIp() string {
	min := 0
	max := 255

	var parts []int

	for i := 0; i < 4; i++ {
		parts = append(parts, rand.Intn(max-min)+min)
	}

	return strings.Trim(strings.Replace(fmt.Sprint(parts), " ", ".", -1), "[]")
}

func putOnRouter(cfg router.Config) {
	// register
	rt := router.New(cfg)

	for i := 1; i < cfg.NumberOfInterface; i++ {
		rt.RegisterIp(strconv.Itoa(i), randomIp())
	}

	for {
		ip := randomIp()

		rt.Next(ip)

		time.Sleep(1 * time.Second)
	}
}
