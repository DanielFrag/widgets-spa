package router

import (
	"net/http"

	"github.com/DanielFrag/widgets-spa-rv/handler"
	"github.com/DanielFrag/widgets-spa-rv/utils"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc utils.HandlerFuncInjector
	ContentType string
}

var openRoutes = []Route{
	Route{
		Method:  "POST",
		Pattern: "/login",
		HandlerFunc: utils.HandlerFuncInjector{
			Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
				handler.UserRepositoryInjector,
			},
			Handler: handler.UserLogin,
		},
	},
}

var apiRoutes = []Route{
	Route{
		Name:    "GetUsers",
		Method:  "GET",
		Pattern: "/users",
		HandlerFunc: utils.HandlerFuncInjector{
			Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
				handler.TokenCheckerMiddleware,
				handler.UserRepositoryInjector,
				handler.UserSessionChecker,
			},
			Handler: handler.GetUsers,
		},
	},
	Route{
		Name:    "GetUserByID",
		Method:  "GET",
		Pattern: "/users/{id}",
		HandlerFunc: utils.HandlerFuncInjector{
			Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
				handler.TokenCheckerMiddleware,
				handler.UserRepositoryInjector,
				handler.UserSessionChecker,
			},
			Handler: handler.GetUserByID,
		},
	},
	Route{
		Name:    "GetWidgets",
		Method:  "GET",
		Pattern: "/widgets",
		HandlerFunc: utils.HandlerFuncInjector{
			Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
				handler.TokenCheckerMiddleware,
				handler.UserRepositoryInjector,
				handler.UserSessionChecker,
				handler.WidgetRepositoryInjector,
			},
			Handler: handler.GetWidgets,
		},
	},
	Route{
		Name:    "GetWidgetById",
		Method:  "GET",
		Pattern: "/widgets/{id}",
		HandlerFunc: utils.HandlerFuncInjector{
			Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
				handler.TokenCheckerMiddleware,
				handler.UserRepositoryInjector,
				handler.UserSessionChecker,
				handler.WidgetRepositoryInjector,
			},
			Handler: handler.GetWidgetById,
		},
	},
	Route{
		Name:    "CreateWidget",
		Method:  "POST",
		Pattern: "/widgets",
		HandlerFunc: utils.HandlerFuncInjector{
			Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
				handler.TokenCheckerMiddleware,
				handler.UserRepositoryInjector,
				handler.UserSessionChecker,
				handler.WidgetRepositoryInjector,
			},
			Handler: handler.CreateWidget,
		},
	},
	Route{
		Name:    "ChangeWidget",
		Method:  "PUT",
		Pattern: "/widgets/{id}",
		HandlerFunc: utils.HandlerFuncInjector{
			Dependencies: []func(http.HandlerFunc) http.HandlerFunc{
				handler.TokenCheckerMiddleware,
				handler.UserRepositoryInjector,
				handler.UserSessionChecker,
				handler.WidgetRepositoryInjector,
			},
			Handler: handler.ChangeWidget,
		},
	},
}

func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range openRoutes {
		route.HandlerFunc.InjectDependencies()
		router.
			HandleFunc(route.Pattern, route.HandlerFunc.Handler).
			Methods(route.Method)
	}
	for _, route := range apiRoutes {
		route.HandlerFunc.InjectDependencies()
		router.
			HandleFunc(route.Pattern, route.HandlerFunc.Handler).
			Methods(route.Method)
	}
	return handler.CorsSetup(router)
}
