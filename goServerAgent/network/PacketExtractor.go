package network

import (
	"fmt"
	"strings"
)

// Extract HTTP request details from payload
func ExtractHTTPDetails(payload string) (string, string, int, uint64) {
	fmt.Println(payload)
	lines := strings.Split(payload, "\r\n")
	if len(lines) < 1 {
		return "", "", 0, 0
	}

	// Extract method and URL from the first line
	requestLine := lines[0]
	parts := strings.Split(requestLine, " ")
	if len(parts) < 2 {
		return "", "", 0, 0
	}
	method := parts[0]
	url := parts[1]

	// Extract status code from the response (if available)
	statusCode := 0
	if strings.HasPrefix(method, "HTTP") {
		// Response line for HTTP responses
		statusParts := strings.Split(requestLine, " ")
		if len(statusParts) > 1 {
			fmt.Sscanf(statusParts[1], "%d", &statusCode)
		}
	}

	// Calculate payload size (excluding headers)
	payloadSize := uint64(len(payload))

	return method, url, statusCode, payloadSize
}
