package main

import (
	"folha/src/app/consumer"
	"folha/src/app/controller"
	"fmt"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/config"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/db"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting module mini-folha")
	config.Load()
	
	conn, err := db.Connect()
	if err != nil {
		log.Fatalf("Could not create database connection: %v", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/employees", controller.EmployeeFindAll).Methods(http.MethodGet)
	router.HandleFunc("/employees/{id}", controller.EmployeeFindById).Methods(http.MethodGet)

	sess := messaging.CreateLocalstackSession()
	consumer.StartSchoolEnrollmentConsumer(sess, conn)

	port := config.BackendPort
	log.Println("Application available on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
