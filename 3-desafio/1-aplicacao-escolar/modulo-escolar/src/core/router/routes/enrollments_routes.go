package routes

import (
	"modulo-escolar/src/application/controllers"
	"net/http"
)

var enrollmentRoutes = []Route{
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
