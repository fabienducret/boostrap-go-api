package server

import (
	"context"
	"fmt"
	"net/http"
)

type Controllers struct {
	Health http.HandlerFunc
}

type Server struct {
	httpServer  *http.Server
	controllers Controllers
}

func New(c Controllers) *Server {
	return &Server{
		httpServer:  &http.Server{},
		controllers: c,
	}
}

func (s *Server) Run(p string) {
	fmt.Printf("start server on port %s\n", p)

	s.httpServer.Addr = fmt.Sprintf(":%s", p)
	loadRoutes(s.controllers)

	s.httpServer.ListenAndServe()
}

func (s *Server) Stop() {
	s.httpServer.Shutdown(context.Background())
}
