package main

import (
	"fmt"
	"log"
	"modulo-folha/src/consumers"
	"modulo-folha/src/controllers"
	"net/http"

	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/config"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/router"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/router/routes"
)

var AppRoutes = []routes.Route{
	{
		URI:      "/schedule",
		Method:   "POST",
		Function: controllers.CalcularFolha,
	},
	{
		URI:      "/folhas",
		Method:   "GET",
		Function: controllers.ListarFolhas,
	},
}

func init() {
	config.Load()
	routes.AddRoutes(AppRoutes)
}

func main() {
	fmt.Println("Iniciando o modulo de folha...")

	ses := messaging.CreateLocalstackSession()
	consumers.IniciarColaboradorConsumer(ses)

	appRouter := router.Create()

	fmt.Printf("Rodando o modulo folha na porta %d", config.BackendPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.BackendPort), appRouter))
}
