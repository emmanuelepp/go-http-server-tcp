package http

import "fmt"

// BuildResponse constructs a simple HTTP response with a "Hello, World!" message.
func BuildResponse(resp Response) []byte {
	statusText := map[int]string{
		200: "OK",
		404: "Not Found",
	}[resp.StatusCode]

	header := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n",
		resp.StatusCode, statusText, len(resp.Body),
	)

	return []byte(header + resp.Body)
}
