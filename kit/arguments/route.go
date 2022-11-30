package arguments

type Route struct {
	r      *Router
	routes []string
}

func (r *Route) Route(routeNames string) *Route {
	r.routes = append(r.routes, routeNames)
	return r
}

func (r *Route) Handle(handler RouteHandler) {
	for _, routeName := range r.routes {
		r.r.routeMap[routeName] = handler
	}
}
