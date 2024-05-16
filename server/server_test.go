package server_test

import (
	"bootstrap-go-api/health"
	"bootstrap-go-api/server"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	client := http.Client{}

	c := server.Controllers{
		Health: health.Controller,
	}

	s := server.New(c)
	go s.Run("8080")

	t.Cleanup(func() {
		s.Stop()
	})

	t.Run("get health", func(t *testing.T) {
		r, _ := client.Get("http://localhost:8080/health")

		if r.StatusCode != 200 {
			t.Errorf("invalid status code, got %d, want %d", r.StatusCode, 200)
		}
	})
}
