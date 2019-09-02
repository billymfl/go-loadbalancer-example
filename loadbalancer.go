package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/billymfl/go-loadbalancer-example/config"
	"github.com/billymfl/go-loadbalancer-example/router"
	_ "github.com/billymfl/go-loadbalancer-example/util"
)

func main() {
	fmt.Printf("%s %s is starting in %s mode...\n", config.Name(), config.Version(), config.Env())

	hostPort := fmt.Sprintf("%s:%s", config.Host(), config.Port())
	if err := http.ListenAndServe(hostPort, router.Router); err != nil {
		log.Fatal(err)
	}
}
