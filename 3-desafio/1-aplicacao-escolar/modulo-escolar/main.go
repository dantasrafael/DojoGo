package main

import (
	"fmt"
	"log"
	"net/http"

	"modulo-escolar/src/application/consumers"
	app_routes "modulo-escolar/src/core/routes"

	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/config"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/router"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/router/routes"
)

func init() {
	config.Load()

	routes.AddRoutes(app_routes.CourseRoutes)
	routes.AddRoutes(app_routes.EnrollmentRoutes)
	routes.AddRoutes(app_routes.StudentRoutes)
}

func main() {
	sess := messaging.CreateLocalstackSession()
	consumers.StartFinantialInstallmentConsumer(sess)

	appRouter := router.Create()

	log.Printf("Rodando o modulo-escolar na porta %d\n", config.BackendPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.BackendPort), appRouter))
}
