package main

import (
	"fmt"
	"net/http"

	"github.com/billymfl/go-loadbalancer-example/config"
	"github.com/billymfl/go-loadbalancer-example/router"
)

// Middleware type
//type Middleware func(http.HandlerFunc) http.HandlerFunc

func main() {
	fmt.Printf("%s %s is starting in %s mode...\n", config.Getname(), config.Getversion(), config.Getenv())

	hostPort := fmt.Sprintf("%s:%s", config.Gethost(), config.Getport())
	if err := http.ListenAndServe(hostPort, router.Router); err != nil {
		panic(err)
	}
}
