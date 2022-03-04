package main

import (
	"fmt"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/config"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/router"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/router/routes"
	"log"
	"net/http"
)

func init() {
	config.Load()
	routes.AddRoutes(ColaboradoresRoutes)
}

func main() {
	log.Println("Inicializando módulo colaboradores")
	approuter := router.Create()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.BackendPort), approuter))
}
