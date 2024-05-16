package main

import (
	"bootstrap-go-api/config"
	"bootstrap-go-api/health"
	"bootstrap-go-api/server"
)

func main() {
	config := config.MustInit()
	runServerWith(config)
}

func runServerWith(config *config.Config) {
	c := controllers()
	s := server.New(c)

	s.Run(config.Port)
}

func controllers() server.Controllers {
	return server.Controllers{
		Health: health.Controller,
	}
}
