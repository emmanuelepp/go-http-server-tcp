package http

import (
	"fmt"
	"strings"
)

// ParseRequestLine splits an HTTP request line into method, path, and version.
func ParseRequestLine(line string) (method, path, version string, err error) {
	line = strings.TrimSpace(line)

	parts := strings.SplitN(line, " ", 3)
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("invalid request line: %s", line)
	}

	method = parts[0]
	path = parts[1]
	version = parts[2]

	return method, path, version, nil
}
