package handler

import (
	"fmt"
	"net/http"
)

//CorsSetup middleware to allow cross domain origin
func CorsSetup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

//RecoverFunc is used by a controller's defer statement. It will close the request's body, check for general errors and format an error response
func RecoverFunc(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
	recoverError := recover()
	if recoverError != nil {
		http.Error(w, fmt.Sprint("Error: ", recoverError), http.StatusInternalServerError)
	}
}
