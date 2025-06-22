package http

import (
	"bufio"
	"fmt"
	"net"
)

// Start launches a raw TCP server on port 8080.
// It accepts one connection and reads the first line of the HTTP request.
func Start() error {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return fmt.Errorf("failed to listen on :8080: %w", err)
	}

	fmt.Println("Listening on port 8080...")

	conn, err := listener.Accept()
	if err != nil {
		return fmt.Errorf("failed to accept connection: %w", err)
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Read the request line (e.g., "GET / HTTP/1.1\r\n")
	line, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read request: %w", err)
	}

	fmt.Printf("Raw request line: %s", line)

	return nil

}
