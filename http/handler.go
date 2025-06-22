package http

import "fmt"

// HandleRequest returns a response based on the given path.
func HandleRequest(method, path string) []byte {
	switch {
	case method == "GET" && path == "/":
		return BuildTextResponse("Welcome (GET /)")
	case method == "POST" && path == "/submit":
		return BuildTextResponse("Submitted!")
	case method == "PUT" && path == "/update":
		return BuildTextResponse("Updated!")
	case method == "DELETE" && path == "/delete":
		return BuildTextResponse("Deleted!")
	default:
		return BuildTextResponse("404 Not Found")
	}

}

// BuildTextResponse builds a 200 OK response with plain text body
func BuildTextResponse(body string) []byte {
	header := fmt.Sprintf(
		"HTTP/1.1 200 OK\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n",
		len(body),
	)
	return []byte(header + body)
}
