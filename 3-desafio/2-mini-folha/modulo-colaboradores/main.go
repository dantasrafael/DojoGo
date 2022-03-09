package main

import (
	"fmt"
	"log"
	"modulo-colaboradores/controller"
	"net/http"

	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/config"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/router"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/router/routes"
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

func init() {
	config.Load()
	routes.AddRoutes(ColaboradoresRoutes)
}

func main() {
	log.Println("Inicializando módulo colaboradores")
	approuter := router.Create()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.BackendPort), approuter))
}
