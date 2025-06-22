package http

import "fmt"

// BuildResponse constructs a simple HTTP response with a "Hello, World!" message.
func BuildResponse() []byte {
	body := "Hello, World!"
	header := fmt.Sprintf(
		"HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n",
		len(body),
	)
	return []byte(header + body)
}
