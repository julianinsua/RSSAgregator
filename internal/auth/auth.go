package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API key from the headers of an http request
// Example:
// Authorization: "ApiKey <API_KEY_VALUE>"
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization header set")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed Authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed Authorization header")
	}
	return vals[1], nil
}
