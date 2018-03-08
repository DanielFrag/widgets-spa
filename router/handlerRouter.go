package router

import (
	"net/http"

	"github.com/DanielFrag/widgets-spa/handler"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	ContentType string
}

var openRoutes = []Route{
	Route{
		Name:        "Login",
		Method:      "POST",
		Pattern:     "/login",
		HandlerFunc: Login,
	},
}

var apiRoutes = []Route{
	Route{
		Name:        "GetUsers",
		Method:      "GET",
		Pattern:     "/users",
		HandlerFunc: UserLogin,
	},
	Route{
		Name:        "GetUserByID",
		Method:      "GET",
		Pattern:     "/users/{id}",
		HandlerFunc: TokenCheckerMiddleware(),
	},
	Route{
		Name:        "GetWidgets",
		Method:      "GET",
		Pattern:     "/widgets",
		HandlerFunc: TokenCheckerMiddleware(),
	},
	Route{
		Name:        "GetWidgetById",
		Method:      "GET",
		Pattern:     "/widgets/{id}",
		HandlerFunc: TokenCheckerMiddleware(),
	},
	Route{
		Name:        "CreateWidgets",
		Method:      "POST",
		Pattern:     "/widgets",
		HandlerFunc: TokenCheckerMiddleware(),
	},
	Route{
		Name:        "ChangeWidgets",
		Method:      "PUT",
		Pattern:     "/widgets/{id}",
		HandlerFunc: TokenCheckerMiddleware(),
	},
}

func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range openRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	for _, route := range apiRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return handler.CorsSetup(router)
}
