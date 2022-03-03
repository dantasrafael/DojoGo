package router

import (
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/router/routes"
	"github.com/gorilla/mux"
)

func Create() *mux.Router {
	router := mux.NewRouter()

	return routes.Config(router)
}
