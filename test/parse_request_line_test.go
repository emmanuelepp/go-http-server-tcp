package http_test

import (
    "testing"
    "go-http-server-tcp/internal/http"
    "github.com/stretchr/testify/require"
)


func TestParseRequestLine_Valid(t *testing.T) {
	method, path, version, err := http.ParseRequestLine("GET /index.html HTTP/1.1")
	require.NoError(t, err)
	require.Equal(t, "GET", method)
	require.Equal(t, "/index.html", path)
	require.Equal(t, "HTTP/1.1", version)

}
