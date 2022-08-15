package router

const (
	numberOfInterface = 10
)

type Router struct {
	defaultGateWay string

	routingTable map[string]string
}

func New() *Router {
	return &Router{
		routingTable: make(map[string]string, numberOfInterface),
	}
}

func (r *Router) RegisterIp(routerInterface string, ip string) {
	r.routingTable[ip] = routerInterface
}

func (r *Router) Next(ip string) string {
	return r.match(ip)
}

func (r *Router) match(ip string) string {
	for key := range r.routingTable {
		if key == ip {
			return r.routingTable[key]
		}
	}

	return r.defaultGateWay
}
