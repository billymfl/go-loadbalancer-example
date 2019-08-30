// Package config configures app defaults and environment variables
package config

import (
	"os"
	"strings"
)

// Config struct
type Config struct {
	name    string // app name
	version string // app version
	key     string // our api key
	host    string // host domain
	port    string // port to listen on
	env     string // the environment (development, staging, production)
}

// Config data
var config *Config

func init() {
	config = &Config{
		name:    "Loadbalancer",
		version: "1.0",
		key:     os.Getenv("KEY"),
		host:    os.Getenv("HOST"),
		port:    os.Getenv("PORT")}

	var errors []string

	if config.port == "" {
		config.port = "80"
	}

	if config.key == "" {
		errors = append(errors, "KEY env var must be set and cannot be empty")
	}

	if len(errors) > 0 {
		panic(strings.Join(errors, "\n"))
	}
}

// Getname returns app name
func Getname() string {
	return config.name
}

// Getversion returns version
func Getversion() string {
	return config.version
}

// Getenv returns env
func Getenv() string {
	return config.env
}

// Getkey returns key
func Getkey() string {
	return config.key
}

// Getport returns env
func Getport() string {
	return config.port
}

// Gethost returns host
func Gethost() string {
	return config.host
}
