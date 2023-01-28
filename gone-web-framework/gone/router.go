package gone

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandleFunc),
	}
}

func (r *router) createRouteKey(method string, pattern string) string {
	return method + "-" + pattern
}

func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := r.createRouteKey(method, pattern)
	r.handlers[key] = handler
}

func (r *router) handle(context *Context) {
	key := r.createRouteKey(context.Method, context.Path)
	if handler, ok := r.handlers[key]; ok {
		handler(context)
	} else {
		context.WriteString(
			http.StatusNotFound,
			"404 NOT FOUND: %s\n",
			context.Path)
	}

}
