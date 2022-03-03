package main

import (
	"financial/app/consumer"
	"financial/app/controller"
	"fmt"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/config"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/db"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	config.Load()

	router := mux.NewRouter()
	router.HandleFunc("/invoices", controller.NewInvoiceController().Find)

	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("could not create database connection: %v", err)
	}

	sess := messaging.CreateLocalstackSession()
	consumer.StartSchoolEnrollmentConsumer(sess, conn)

	port := config.BackendPort
	log.Println("starting application on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
