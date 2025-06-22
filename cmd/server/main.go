package main

import (
	"go-http-server-tcp/internal/http"

	"log"
)

func main() {
	if err := http.Start(); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
