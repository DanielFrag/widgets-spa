package handler

import (
	"net/http"
	"github.com/gorilla/context"
	"github.com/DanielFrag/widgets-spa-rv/repository"
	"github.com/DanielFrag/widgets-spa-rv/utils"
)

//TokenCheckerMiddleware performs a user's token validation, extracting the payload, and generate a new token for user
func TokenCheckerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userToken := r.Header.Get("authorization")
		if userToken == "" {
			http.Error(w, "Error: no token provided", http.StatusBadRequest)
			return
		}
		payload, tokenError := utils.TokenChecker(userToken)
		if tokenError != nil {
			http.Error(w, "Error reading access token: " + tokenError.Error(), http.StatusForbidden)
			return
		}
		validatedToken, _ := utils.EncodeToken(payload)
		context.Set(r, "TokenPayload", payload)
		context.Set(r, "Token", validatedToken)
		next(w, r)
		return
	})
}

func UserRepositoryInjector(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "UserRepository", repository.GetUserRepository())
		next(w, r)
		return
	})
}

func UserSessionChecker(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRepository, userRepositoryError := extractUserRepository(r)
		if userRepositoryError != nil {
			http.Error(w, userRepositoryError.Error(), http.StatusInternalServerError)
			return
		}
		tokenPayload := context.Get(r, "TokenPayload")
		if tokenPayload == nil {
			http.Error(w, "Can't access the user's token payload", http.StatusInternalServerError)
			return
		}
		userData, ok := tokenPayload.(map[string]string)
		if !ok {
			http.Error(w, "Can't reconize the user's token payload", http.StatusInternalServerError)
			return
		}
		user, userError := userRepository.GetUserByID(userData["userID"])
		if userError != nil {
			http.Error(w, "Can't find the requested user: " + userError.Error(), http.StatusUnauthorized)
			return
		}
		if userData["userSession"] == "" || user.Session != userData["userSession"] {
			http.Error(w, "Invalid user session", http.StatusForbidden)
			return
		}
		next(w, r)
		return
	})
}

func WidgetRepositoryInjector(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "WidgetRepository", repository.GetWidgetRepository())
		next(w, r)
		return
	})
}
