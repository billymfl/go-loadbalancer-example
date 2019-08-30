package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/billymfl/go-loadbalancer-example/config"
	"github.com/gorilla/mux"
)

// Root returns app name and version
// Get /
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s", config.Getname(), config.Getversion())
}

// HealthCheck heartbeat
// GET /heathcheck
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

// Register registers a chat server
// PUT /register/{name}/{cpus}/{rooms}/{verson}
func Register(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	cpus := vars["cpus"]
	rooms := vars["rooms"]
	version := vars["version"]

	fmt.Fprintf(w, "%s %s %s %s", name, cpus, rooms, version)
}
