package server

import "net/http"

func loadRoutes(c Controllers) {
	http.HandleFunc("GET /health", c.Health)
}
