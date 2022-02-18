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

func Config(r *mux.Router) *mux.Router {
	var routes []Route
	routes = append(routes, studentRoutes...)
	routes = append(routes, courseRoutes...)
	routes = append(routes, enrollmentRoutes...)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
