package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Página Raiz!"))
	})

	http.HandleFunc("/users", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Página usuários!"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
