package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Página Raiz!"))
	})

	http.HandleFunc("/nutela", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Página nutela!"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
