package router

type Router struct {
	routingTable map[string][]string
}

func New() *Router {
	return &Router{
		routingTable: make(map[string][]string),
	}
}

func (r *Router) match(ip string) string {
	return ""
}

func (r *Router) Next(ip string) {

}
