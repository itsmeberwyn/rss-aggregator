package helper

import (
	"fmt"
	"strings"
)

func ExtractAPIKey(apiKey string) (string, error) {
	authorizationHeader := strings.Split(apiKey, " ")
	if len(authorizationHeader) == 1 {
		return "", fmt.Errorf("invalid header")
	}
	return authorizationHeader[1], nil
}
