package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/DanielFrag/widgets-spa-rv/utils"
	"github.com/DanielFrag/widgets-spa-rv/repository"
	"github.com/DanielFrag/widgets-spa-rv/model"
	"github.com/gorilla/context"
)

func TestTokenChecker(t *testing.T) {
	fA := func (w http.ResponseWriter, r *http.Request) {}
	hfi := utils.HandlerFuncInjector{
		Dependencies: []func (http.HandlerFunc) http.HandlerFunc {
			TokenCheckerMiddleware,
		},
		Handler: fA,
	}
	hfi.InjectDependencies()
	t.Run("ValidToken", func (t *testing.T) {
		req, reqError := http.NewRequest("GET", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
		}
		req.Header["Authorization"] = []string{`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJJRCI6IjEyMzQiLCJzZXNzaW9uIjoiMTIzNDU2Nzg5MCJ9LCJpYXQiOjE1MjA0NzAwNDcsImlzIjowfQ.G4TWM1MFNbHSZHLogG3OFsxNnwpmYt9iwpaaJqJGTM0`}
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode == 400 || result.StatusCode == 403 {
			t.Error("Can't validate the token: " + reqRecorder.Body.String())
		}
		tokenPayload := context.Get(req, "TokenPayload")
		newToken := context.Get(req, "Token")
		if tokenPayload == nil || newToken == nil {
			t.Error("New token not seted")
		}
	})
	t.Run("NoToken", func (t *testing.T) {
		req, reqError := http.NewRequest("GET", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
		}
		req.Header["Authorization"] = []string{""}
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 400 {
			t.Error("Wrong status code. Expected 400 and got " + string(result.StatusCode))
		}
	})
}

func TestRepositoryInjection(t *testing.T) {
	fA := func (w http.ResponseWriter, r *http.Request) {}
	hfi := utils.HandlerFuncInjector{
		Dependencies: []func (http.HandlerFunc) http.HandlerFunc {
			UserRepositoryInjector,
			WidgetRepositoryInjector,
		},
		Handler: fA,
	}
	hfi.InjectDependencies()
	req, reqError := http.NewRequest("GET", "/", nil)
	if reqError != nil {
		t.Error("Error to create the request: " + reqError.Error())
	}
	reqRecorder := httptest.NewRecorder()
	hfi.Handler.ServeHTTP(reqRecorder, req)
	userRepositoryContext := context.Get(req, "UserRepository")
	widgetRepositoryContext := context.Get(req, "WidgetRepository")
	if userRepositoryContext == nil || widgetRepositoryContext == nil {
		t.Error("Repository not seted in requisiton context")
	}
	_, userOk := userRepositoryContext.(repository.UserRepository)
	_, widgetOk := widgetRepositoryContext.(repository.WidgetRepository)
	if !userOk || !widgetOk {
		t.Error("Context seted with an invalid interface")
	}
}

func TestUserSessionChecker(t *testing.T) {
	var firstUser model.User
	userMock := UserDBMock{}
	userMock.InitializeUserDB()
	users, _ := userMock.GetUsers()
	firstUser = users[0]
	fA := func (w http.ResponseWriter, r *http.Request) {}
	dbInjector := func (next http.HandlerFunc) http.HandlerFunc {
		return func (w http.ResponseWriter, r *http.Request) {
			context.Set(r, "TokenPayload", map[string]string {
				"userID": firstUser.ID.Hex(),
				"userSession": firstUser.Session,
			})
			context.Set(r, "UserRepository", &userMock)
			next(w, r)
			return
		}
	}
	hfi := utils.HandlerFuncInjector{
		Dependencies: []func (http.HandlerFunc) http.HandlerFunc {
			dbInjector,
			UserSessionChecker,
		},
		Handler: fA,
	}
	hfi.InjectDependencies()
	t.Run("ValidUser", func(t *testing.T) {
		req, reqError := http.NewRequest("GET", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
		}
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode == 500 || result.StatusCode == 401 {
			t.Error("Can't found the user with the session")
		}
	})
	t.Run("InvalidUser", func(t *testing.T) {
		firstUser = model.User{
			ID: firstUser.ID,
			Session: firstUser.Session + "2",
		}
		req, reqError := http.NewRequest("GET", "/", nil)
		if reqError != nil {
			t.Error("Error to create the request: " + reqError.Error())
		}
		reqRecorder := httptest.NewRecorder()
		hfi.Handler.ServeHTTP(reqRecorder, req)
		result := reqRecorder.Result()
		if result.StatusCode != 403 {
			t.Error("Wrong status code. Expected 403 and got " + string(result.StatusCode))
		}
	})
}