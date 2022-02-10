package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Nome  string
	Email string
}

func main() {
	user := usuario{"Pedro", "pedro@email.com"}
	pages := template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pages.ExecuteTemplate(w, "home.html", user)
	})

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
