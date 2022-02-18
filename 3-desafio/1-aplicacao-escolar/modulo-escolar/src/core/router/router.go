package router

import (
	"modulo-escolar/src/core/router/routes"

	"github.com/gorilla/mux"
)

func Create() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
