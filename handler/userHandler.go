package handler

import (
	"errors"
	"net/http"
	"github.com/DanielFrag/widgets-spa-rv/repository"
	"github.com/DanielFrag/widgets-spa-rv/utils"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userRepository, userRepositoryError := extractUserRepository(r)
	if userRepositoryError != nil {
		http.Error(w, userRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	users, usersError := userRepository.GetUsers()
	if usersError != nil {
		http.Error(w, "Error: " + usersError.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(users))
	return
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		http.Error(w, "Error id not provided", http.StatusBadRequest)
		return
	}
	userRepository, userRepositoryError := extractUserRepository(r)
	if userRepositoryError != nil {
		http.Error(w, userRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	user, userError := userRepository.GetUserByID(vars["id"])
	if userError != nil {
		http.Error(w, "Error: " + userError.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(user))
	return
}

func extractUserRepository(r *http.Request) (repository.UserRepository, error) {
	contextUserRepository := context.Get(r, "UserRepository")
	if contextUserRepository == nil {
		return nil, errors.New("Can't access the context user repository")
	}
	userRepository, userRepositoryOk := contextUserRepository.(repository.UserRepository)
	if !userRepositoryOk {
		return nil, errors.New("Can't access the user repository")
	}
	return userRepository, nil
}