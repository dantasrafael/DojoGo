package routes

import (
	"modulo-escolar/src/application/controllers"
	"net/http"
)

var studentRoutes = []Route{
	{
		URI:      "/students",
		Method:   http.MethodPost,
		Function: controllers.CreateStudent,
	},
	{
		URI:      "/students",
		Method:   http.MethodGet,
		Function: controllers.FindAllStudents,
	},
	{
		URI:      "/students/{id}",
		Method:   http.MethodGet,
		Function: controllers.FindStudentById,
	},
	{
		URI:      "/students/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdateStudent,
	},
	{
		URI:      "/students/{id}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteStudent,
	},
}
