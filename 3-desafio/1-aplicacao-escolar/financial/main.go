package main

import (
	"financial/app/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/invoices", controller.NewInvoiceController().Find)

	log.Fatal(http.ListenAndServe(":8181", nil))
}
