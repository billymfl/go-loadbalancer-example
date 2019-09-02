package router

import (
	"net/http"
	"strings"

	"github.com/billymfl/go-loadbalancer-example/handlers"
	"github.com/gorilla/mux"
)

// Route is a path, handler and allowed methods for an endpoint
type Route struct {
	path    string                                   // endpoint path
	handler func(http.ResponseWriter, *http.Request) // handler func for this path
	methods []string                                 // methods allowed for this path
	secure  bool                                     //if this path requires authenticated use
}

// Routes configs the paths, handlers and methods for endpoints
var Routes = []*Route{
	{"/", handlers.Root, []string{"GET"}, false},
	{"/healthcheck", handlers.HealthCheck, []string{"GET"}, false},
	{"/register/{name}/{cpus}/{rooms}/{version}", handlers.Register, []string{"PUT"}, true},
	{"/register/{name}", handlers.Unregister, []string{"DELETE"}, true},
	{"/list", handlers.List, []string{"GET"}, true},
	{"/find/{version}", handlers.Find, []string{"GET"}, true},
}

// Router the mux router
var Router = mux.NewRouter()

// configure the Router's HandleFunc with our Routes
func init() {
	Router.Use(defaults)

	var handler http.HandlerFunc
	for _, route := range Routes {
		if route.secure {
			handler = chain(route.handler, Authenticated())
		} else {
			handler = chain(route.handler)
		}
		Router.HandleFunc(route.path, handler).Methods(strings.Join(route.methods, ","))
	}
}

// chain applies middlewares to a http.HandlerFunc
func chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func defaults(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
