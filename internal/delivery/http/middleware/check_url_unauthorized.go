package middleware

import "strings"

// function ini mengembalikan true jika URL tidak butuh JWT
func IsUnauthorizedURL(requestPath string, whitelist []string) bool {
	for _, url := range whitelist {
		// exact match
		if strings.EqualFold(requestPath, url) {
			return true
		}
	}
	return false
}
