package router

type Router struct {
	routingTable map[string]string
}

func New() *Router {
	return &Router{
		routingTable: make(map[string]string),
	}
}

func (r *Router) match(ip string) string {
	for key := range r.routingTable {
		if key == ip {
			return r.routingTable[key]
		}
	}

	return "default gateway"
}

func (r *Router) Register(routerInterface string, ip string) {
	r.routingTable[ip] = routerInterface
}

func (r *Router) Next(ip string) string {
	return r.match(ip)
}
