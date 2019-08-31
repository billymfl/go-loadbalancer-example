package router

import (
	"net/http"

	"github.com/billymfl/go-loadbalancer-example/config"
)

// Middleware type
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Authenticated ensures that url can only be requested with an api key, else returns a 403 Forbidden
func Authenticated() Middleware {

	// Create a new Middleware
	return func(next http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get("X-API-Key")

			if config.Key() != key {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			} else {
				next(w, r)
			}
		}
	}
}
