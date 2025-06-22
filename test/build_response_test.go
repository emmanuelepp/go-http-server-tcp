package http_test

import (
	"go-http-server-tcp/internal/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildResponse(t *testing.T) {
	resp := http.Response{
		StatusCode: 200,
		Body:       "Hellooooooo, Test!",
	}

	result := http.BuildResponse(resp)

	expected := "HTTP/1.1 200 OK\r\nContent-Length: 12\r\nContent-Type: text/plain\r\n\r\nHellooooooo, Test!"

	require.Equal(t, expected, string(result))
}
