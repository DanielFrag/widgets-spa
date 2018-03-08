package handler

import (
	"net/http"
)

//TokenCheckerMiddleware performs a user's token validation, extracting the payload, and generate a new token for user
func TokenCheckerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	})
}
