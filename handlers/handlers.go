package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/billymfl/go-loadbalancer-example/config"
	"github.com/billymfl/go-loadbalancer-example/util"
	"github.com/gorilla/mux"
)

// Root returns app name and version
// Get /
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"app": "%s", "version": "%s"}`, config.Name(), config.Version())
}

// HealthCheck heartbeat
// GET /heathcheck
// Does not require API key
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"alive": true}`)
}

// Register registers a chat server
// PUT /register/{name}/{cpus}/{rooms}/{verson}
func Register(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	cpus, _ := strconv.Atoi(vars["cpus"])
	rooms, _ := strconv.Atoi(vars["rooms"])
	version := vars["version"]

	msg := util.Register(name, cpus, rooms, version)
	fmt.Fprintf(w, `{"message": "%s"}`, msg)
}

// Unregister removes server with key name
// DELETE /register/{name}
func Unregister(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	util.Unregister(name)
	fmt.Fprint(w, `{"message": "unregistered"}`)
}

// List returns a json data of all registered servers
// GET /list
// TODO: pagination
func List(w http.ResponseWriter, r *http.Request) {
	data, err := util.List()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		fmt.Fprint(w, data)
	}
}
