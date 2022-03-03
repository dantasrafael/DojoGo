package routes

import (
	"modulo-escolar/src/application/controllers"
	"net/http"

	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/router/routes"
)

var EnrollmentRoutes = []routes.Route{
	{
		URI:      "/enrollments",
		Method:   http.MethodPost,
		Function: controllers.CreateEnrollment,
	},
	{
		URI:      "/enrollments",
		Method:   http.MethodGet,
		Function: controllers.FindAllEnrollments,
	},
	{
		URI:      "/enrollments/{id}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteEnrollment,
	},
}
