package main

import (
	"fmt"
	"log"
	"net/http"

	"crud-basico/server"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.FindAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.FindUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Servidor subiu na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
