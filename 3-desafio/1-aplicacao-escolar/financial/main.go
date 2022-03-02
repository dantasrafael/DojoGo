package main

import (
	"financial/app/consumer"
	"financial/app/controller"
	"github.com/dantasrafael/DojoGo/tree/master/3-desafio/starters/messaging"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/invoices", controller.NewInvoiceController().Find)

	sess := messaging.CreateLocalstackSession()
	consumer.StartSchoolEnrollmentConsumer(sess)

	log.Println("starting application on port 8181")
	log.Fatal(http.ListenAndServe(":8181", router))
}
