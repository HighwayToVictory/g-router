package router

import "github.com/HighwayToVictory/g-router/internal/metrics"

type Router struct {
	defaultGateWay string

	routingTable map[string]string

	metrics metrics.Metrics
}

func New(cfg Config) Router {
	return Router{
		defaultGateWay: cfg.DefaultGateWay,

		routingTable: make(map[string]string, cfg.NumberOfInterface),

		metrics: metrics.NewMetrics(),
	}
}

func (r *Router) RegisterIp(routerInterface string, ip string) {
	r.routingTable[ip] = routerInterface
}

func (r *Router) Next(ip string) string {
	r.metrics.RequestCounter.Add(1)

	return r.match(ip)
}

func (r *Router) match(ip string) string {
	for key := range r.routingTable {
		if key == ip {
			return r.routingTable[key]
		}
	}

	r.metrics.GatewayCounter.Add(1)

	return r.defaultGateWay
}
