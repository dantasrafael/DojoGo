package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

var appRoutes []Route

func AddRoutes(routes []Route) {
	appRoutes = append(appRoutes, routes...)
}

func Config(router *mux.Router) *mux.Router {
	for _, route := range appRoutes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return router
}
