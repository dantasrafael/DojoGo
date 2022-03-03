package routes

import (
	"net/http"

	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/middlewares"
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
		router.HandleFunc(route.URI, middlewares.HttpLogger(route.Function)).Methods(route.Method)
	}

	return router
}
