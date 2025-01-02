package health_test

import (
	"bootstrap-go-api/health"
	"bootstrap-go-api/server"
	"io"
	"net/http"
	"testing"
)

func TestHealthController(t *testing.T) {
	s := server.New(server.Controllers{
		Health: health.Controller,
	})
	go s.Run("8080")

	t.Cleanup(func() {
		s.Stop()
	})

	t.Run("get body with success", func(t *testing.T) {
		wantStatus := 200
		wantBody := "ok"

		body, status := bodyAndStatusFromUrl(t, "http://localhost:8080/health")

		if status != wantStatus {
			t.Errorf("invalid status, got %d, want %d", status, wantStatus)
		}

		if body != wantBody {
			t.Errorf("invalid body, got %s, want %s", body, wantBody)
		}
	})
}

func bodyAndStatusFromUrl(t *testing.T, url string) (string, int) {
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

	return string(b), r.StatusCode
}
