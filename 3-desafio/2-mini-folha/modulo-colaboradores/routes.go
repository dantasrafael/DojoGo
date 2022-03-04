package main

import (
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/router/routes"
	"modulo-colaboradores/controller"
	"net/http"
)

var ColaboradoresRoutes = []routes.Route{
	{
		URI:      "/colaboradores",
		Method:   http.MethodPost,
		Function: controller.SalvarColaborador,
	},
	{
		URI:      "/colaboradores",
		Method:   http.MethodGet,
		Function: controller.BuscarTodosColaboradores,
	},
	{
		URI:      "/colaboradores/{id}",
		Method:   http.MethodDelete,
		Function: controller.DeletarColaborador,
	},
}
