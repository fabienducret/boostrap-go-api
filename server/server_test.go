package server_test

import (
	"bootstrap-go-api/health"
	"bootstrap-go-api/server"
	"io"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	s := server.New(controllers())
	go s.Run("8080")

	t.Cleanup(func() {
		s.Stop()
	})

	t.Run("get health body with success", func(t *testing.T) {
		want := "ok"

		got := bodyFromUrl(t, "http://localhost:8080/health")

		if got != want {
			t.Errorf("invalid body, got %s, want %s", got, want)
		}
	})
}

func controllers() server.Controllers {
	return server.Controllers{
		Health: health.Controller,
	}
}

func bodyFromUrl(t *testing.T, url string) string {
	client := http.Client{}

	r, err := client.Get(url)
	if err != nil {
		t.Fatalf("error on get url %v", err)
	}

	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("error on parsing body %v", err)
	}

	return string(b)
}
