package handler

import (
	"net/http"
	"github.com/gorilla/context"
	"github.com/DanielFrag/widgets-spa-rv/repository"
)

//TokenCheckerMiddleware performs a user's token validation, extracting the payload, and generate a new token for user
func TokenCheckerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	})
}

func UserRepositoryInjector(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "UserRepository", repository.GetUserRepository())
		next(w, r)
	})
}

func WidgetRepositoryInjector(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	})
}