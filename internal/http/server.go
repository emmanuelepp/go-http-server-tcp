package http

import (
	"bufio"
	"fmt"
	"net"
)

const asciiArt = `
  _    _ _______ _______ _____     _____ ______ _______      ________ _____  
 | |  | |__   __|__   __|  __ \   / ____|  ____|  __ \ \    / /  ____|  __ \ 
 | |__| |  | |     | |  | |__) | | (___ | |__  | |__) \ \  / /| |__  | |__) |
 |  __  |  | |     | |  |  ___/   \___ \|  __| |  _  / \ \/ / |  __| |  _  / 
 | |  | |  | |     | |  | |       ____) | |____| | \ \  \  /  | |____| | \ \ 
 |_|  |_|  |_|     |_|  |_|      |_____/|______|_|  \_\  \/   |______|_|  \_\

Go HTTP Server - Powered by TCP
Created by: emmanuelepp
`

// Start launches a TCP server on port 8080 that handles HTTP requests.
func Start() error {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return fmt.Errorf("failed to listen on :8080: %w", err)
	}

	fmt.Print(asciiArt)
	fmt.Println("Listening on port 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Read the request line ("GET / HTTP/1.1")
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read request: %v\n", err)
		return
	}
	fmt.Printf("Raw request line (DEBUG): %q\n", line)

	// Parse method, path, version
	method, path, version, err := ParseRequestLine(line)
	if err != nil {
		fmt.Printf("Failed to parse request line: %v\n", err)
		return
	}
	fmt.Printf("Parsed request - Method: %s, Path: %s, Version: %s\n", method, path, version)

	// Read headers
	for {
		headerLine, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read header line: %v\n", err)
			return
		}
		if headerLine == "\r\n" || headerLine == "\n" {
			break
		}
		fmt.Printf("Header line: %s", headerLine)
	}

	// Build request and response
	req := Request{Method: method, Path: path}
	resp := HandleRequest(req)
	raw := BuildResponse(resp)

	_, err = conn.Write(raw)
	if err != nil {
		fmt.Printf("Failed to write response: %v\n", err)
	}
}
