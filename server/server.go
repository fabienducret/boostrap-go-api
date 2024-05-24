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

func (s *Server) Run(port string) {
	fmt.Printf("start server on port %s\n", port)

	s.httpServer.Addr = fmt.Sprintf(":%s", port)
	s.registerRoutes()

	s.httpServer.ListenAndServe()
}

func (s *Server) Stop() {
	s.httpServer.Shutdown(context.TODO())
}
