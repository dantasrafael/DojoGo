package main

// go get (instala todas as dependencias)
// go get URL (instala a dependencia)

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("dantasrafael")
	fmt.Println(erro)
}
