package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/DanielFrag/widgets-spa-rv/utils"
	"github.com/DanielFrag/widgets-spa-rv/model"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func TestGetUsers(t *testing.T) {
	userMock := UserDBMock{}
	userMock.InitializeUserDB()
	dbInjector := func (next http.HandlerFunc) http.HandlerFunc {
		return func (w http.ResponseWriter, r *http.Request) {
			context.Set(r, "UserRepository", &userMock)
			next(w, r)
			return
		}
	}
	t.Run("GetAllUsers", func(t *testing.T) {
		hfi := utils.HandlerFuncInjector{
			Dependencies: []func (http.HandlerFunc) http.HandlerFunc {
				dbInjector,
			},
			Handler: GetUsers,
		}
		hfi.InjectDependencies()
		req, reqError := http.NewRequest("GET", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
		}
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		var result []model.User
		jsonError := json.Unmarshal(reqRecorder.Body.Bytes(), &result)
		if jsonError != nil {
			t.Error("Json error: " + jsonError.Error())
		}
		usersDB, _ := userMock.GetUsers()
		for i, user := range usersDB {
			if user.ID.Hex() != result[i].ID.Hex() {
				t.Error("Inconsistent data")
			}
		}
	})
	t.Run("GetUserByID", func(t *testing.T) {
		users, _ := userMock.GetUsers()
		userID := users[0].ID.Hex()
		hfi := utils.HandlerFuncInjector{
			Dependencies: []func (http.HandlerFunc) http.HandlerFunc {
				dbInjector,
			},
			Handler: GetUserByID,
		}
		hfi.InjectDependencies()
		req, reqError := http.NewRequest("GET", "/" + userID, nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
		}
		reqRecorder := httptest.NewRecorder()
		r := mux.NewRouter()
		r.StrictSlash(true).HandleFunc("/{id}", hfi.Handler).Methods("GET")
		r.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode == 400 || result.StatusCode == 500 {
			t.Error("Wrong response")
		}
		var user model.User
		jsonError := json.Unmarshal(reqRecorder.Body.Bytes(), &user)
		if jsonError != nil {
			t.Error("Json error: " + jsonError.Error())
		}
		if userID != user.ID.Hex() {
			t.Error("Wrong user")
		}
	})
}