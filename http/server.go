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

	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("failed to accept connection: %w", err)
		}

		reader := bufio.NewReader(conn)

		// Read the request line ("GET / HTTP/1.1\r\n")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read request: %v\n", err)
			conn.Close()
			continue
		}

		fmt.Printf("Raw request line: %s", line)

		// Parse the request line into method, path, and version
		method, path, version, err := ParseRequestLine(line)
		if err != nil {
			fmt.Printf("Failed to parse request line: %v\n", err)
			conn.Close()
			continue
		}
		fmt.Printf("Parsed request - Method: %s, Path: %s, Version: %s\n", method, path, version)

		// Read headers line by line until an empty line indicates end of headers
		for {
			headerLine, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Failed to read header line: %v\n", err)
				break
			}

			if headerLine == "\r\n" || headerLine == "\n" {
				break
			}

			fmt.Printf("Header line: %s", headerLine)
		}

		response := HandleRequest(method, path)
		_, err = conn.Write(response)
		if err != nil {
			fmt.Printf("Failed to write response: %v\n", err)
		}

		conn.Close()
	}
}
