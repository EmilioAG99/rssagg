package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts an API key from the headers of an HHTP request
// Example:
// Authorization: Apikey {insert api key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("no authentication info found")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first parth of auth header")
	}

	return vals[1], nil

}
