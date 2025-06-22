package http_test

import (
	"testing"

	"go-http-server-tcp/internal/http"

	"github.com/stretchr/testify/require"
)

func TestHandleRequest(t *testing.T) {
	tests := []struct {
		method   string
		path     string
		expected string
	}{
		{"GET", "/", "Welcome (GET /)"},
		{"POST", "/submit", "Submitted!"},
		{"PUT", "/update", "Updated!"},
		{"DELETE", "/delete", "Deleted!"},
		{"GET", "/unknown", "404 Not Found"},
	}

	for _, tt := range tests {
		req := http.Request{
			Method: tt.method,
			Path:   tt.path,
		}

		res := http.HandleRequest(req)
		require.Equal(t, tt.expected, res.Body)

	}
}
