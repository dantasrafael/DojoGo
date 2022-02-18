package routes

import (
	"modulo-escolar/src/application/controllers"
	"net/http"
)

var courseRoutes = []Route{
	{
		URI:      "/courses",
		Method:   http.MethodPost,
		Function: controllers.CreateCourse,
	},
	{
		URI:      "/courses",
		Method:   http.MethodGet,
		Function: controllers.FindAllCourses,
	},
	{
		URI:      "/courses/{id}",
		Method:   http.MethodGet,
		Function: controllers.FindStudentById,
	},
	{
		URI:      "/courses/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdateCourse,
	},
	{
		URI:      "/courses/{id}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteCourse,
	},
}
