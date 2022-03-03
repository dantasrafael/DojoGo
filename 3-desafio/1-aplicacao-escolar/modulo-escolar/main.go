package main

import (
	"fmt"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/config"
	"log"
	"modulo-escolar/src/core/router"
	"net/http"
)

func main() {
	config.Load()
	r := router.Create()

	log.Printf("Rodando o modulo-escolar na porta %d\n", config.BackendPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.BackendPort), r))
}
