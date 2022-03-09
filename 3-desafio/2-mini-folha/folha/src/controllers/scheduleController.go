package controllers

import (
	"fmt"
	"modulo-folha/src/repositories"
	"net/http"
)

func CalcularFolha(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calculando a folha...")
	
	if err := repositories.CalcularFolha(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Calculo da folha finalizado...")
	
	w.WriteHeader(http.StatusOK)
}
