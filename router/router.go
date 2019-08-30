package router

import (
	"net/http"
	"strings"

	"github.com/billymfl/go-loadbalancer-example/handlers"
	"github.com/gorilla/mux"
)

type routerConfig struct {
	path    string                                   // path
	handler func(http.ResponseWriter, *http.Request) // handler func for path
	methods []string                                 // methods allowed for path
	secure  bool                                     //if this path requires authenticated use
}

// RouterConfig configs the paths, handlers and methods for the paths
var routerconfig = []*routerConfig{
	{"/", handlers.Root, []string{"GET"}, false},
	{"/healthcheck", handlers.HealthCheck, []string{"GET"}, false},
	{"/register/{name}/{cpus}/{rooms}/{version}", handlers.Register, []string{"PUT"}, true},
}

// Router the mux router
var Router = mux.NewRouter()

// configure the Router's HandleFunc with our routerconfig
func init() {
	for _, route := range routerconfig {
		var handler http.HandlerFunc
		if route.secure {
			handler = Chain(route.handler, Authenticated())
		} else {
			handler = Chain(route.handler)
		}
		Router.HandleFunc(route.path, handler).Methods(strings.Join(route.methods, ","))
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
