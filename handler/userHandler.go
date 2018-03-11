package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"encoding/json"
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

func UserLogin(w http.ResponseWriter, r *http.Request) {
	body, bodyReadError := ioutil.ReadAll(r.Body)
	if bodyReadError != nil {
		http.Error(w, "Error reading body request: " + bodyReadError.Error(), http.StatusInternalServerError)
		return
	}
	userRepository, userRepositoryError := extractUserRepository(r)
	if userRepositoryError != nil {
		http.Error(w, userRepositoryError.Error(), http.StatusInternalServerError)
		return
	}
	var jsonUser map[string]string
	jsonError := json.Unmarshal(body, &jsonUser)
	if jsonError != nil {
		http.Error(w, "Json error: " + jsonError.Error(), http.StatusInternalServerError)
		return
	}
	user, userError := userRepository.GetUserByLogin(jsonUser["login"], jsonUser["password"])
	if userError != nil {
		http.Error(w, "Can't find the requested user: " + userError.Error(), http.StatusNotFound)
		return
	}
	session := utils.GenerateRandomAlphaNumericString(10)
	updateSessionError := userRepository.UpdateUserSession(user.ID.Hex(), session)
	if updateSessionError != nil {
		http.Error(w, "Can't update the user session: " + updateSessionError.Error(), http.StatusInternalServerError)
		return
	}
	m := map[string]string {
		"userID": user.ID.Hex(),
		"userSession": session,
	}
	token, tokenError := utils.EncodeToken(m)
	if tokenError != nil {
		http.Error(w, "Can't sign the token: " + tokenError.Error(), http.StatusInternalServerError)
		return
	}
	m2 := map[string]string {
		"token": token,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(utils.FormatJSON(m2))
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