package auth

import (
	"fmt"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from the header of
// an HTTP request.
//
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", fmt.Errorf("no authentication info found")
	}

	vals := strings.Fields(val)
	if len(vals) != 2 {
		return "", fmt.Errorf("malformed authentication header")
	}

	if vals[0] != "ApiKey" {
		return "", fmt.Errorf("malformed authentication header")
	}

	return vals[1], nil
}
