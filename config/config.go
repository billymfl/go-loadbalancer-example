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
		name:    "loadbalancer",
		version: "1.0",
		key:     os.Getenv("KEY"),
		host:    os.Getenv("HOST"),
		port:    os.Getenv("PORT"),
		env:     os.Getenv("ENV")}

	var errors []string

	if config.env == "" {
		config.env = "development"
	}

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

// Name returns app name
func Name() string {
	return config.name
}

// Version returns version
func Version() string {
	return config.version
}

// Env returns env
func Env() string {
	return config.env
}

// Key returns key
func Key() string {
	return config.key
}

// Port returns env
func Port() string {
	return config.port
}

// Host returns host
func Host() string {
	return config.host
}
