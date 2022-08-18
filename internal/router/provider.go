package router

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func NewProvider(cfg Config) {
	putOnRouter(cfg)
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

func putOnRouter(cfg Config) {
	// register
	rt := New(cfg)

	for i := 1; i < cfg.NumberOfInterface; i++ {
		rt.RegisterIp(strconv.Itoa(i), randomIp())
	}

	for {
		ip := randomIp()

		rt.Next(ip)

		time.Sleep(1 * time.Second)
	}
}
