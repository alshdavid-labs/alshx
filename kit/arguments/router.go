package arguments

import "fmt"

type RouteHandler func(args *Arguments)

type Router struct {
	routeMap map[string]RouteHandler
}

func NewRouter() *Router {
	return &Router{
		routeMap: map[string]RouteHandler{},
	}
}

func (r *Router) Route(routeNames ...string) *Route {
	route := &Route{r, routeNames}
	return route
}

func (r *Router) Run(args []string) {
	if len(args) == 0 {
		return
	}
	routeName := args[1]
	route := r.routeMap[routeName]
	if route == nil || route == nil && routeName == "help" {
		for key, _ := range r.routeMap {
			fmt.Println("Argument", key)
		}
		return
	}
	a := newArguments(args[2:])
	route(a)
}
