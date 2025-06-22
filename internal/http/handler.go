package http

import "fmt"

type Request struct {
	Method string
	Path   string
}

type Response struct {
	StatusCode int
	Body       string
}

// HandleRequest returns a response based on the given path.
func HandleRequest(req Request) Response {
	switch {
	case req.Method == "GET" && req.Path == "/":
		return Response{StatusCode: 200, Body: "Welcome to the home page!"}
	case req.Method == "POST" && req.Path == "/submit":
		return Response{StatusCode: 200, Body: "Submitted!"}
	case req.Method == "PUT" && req.Path == "/update":
		return Response{StatusCode: 200, Body: "Updated!"}
	case req.Method == "DELETE" && req.Path == "/delete":
		return Response{StatusCode: 200, Body: "Deleted!"}
	default:
		return Response{StatusCode: 404, Body: "Not Found"}
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
