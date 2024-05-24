package server

import "net/http"

func (s *Server) registerRoutes() {
	http.HandleFunc("GET /health", s.controllers.Health)
}
