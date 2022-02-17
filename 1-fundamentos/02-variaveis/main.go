package main

import "fmt"

func main() {
	// const NOME tipo = VALOR
	// var NOME tipo "sem '=' -> (valor zero)"
	// var NOME tipo = VALOR"
	// var ( NOME tipo = VALOR ) -> agrupar"
	// NOME := VALOR"
	// NOME1, NOME2 := VALOR1, VALOR2"

	const constante1 string = "Constante 1"
	var variavel1 string
	var variavel2 string = "Variável 2"
	var (
		variavel3 string = "Variável 3"
		variavel4 int    = 4
		variavel8        = 4
	)

	variavel5 := "Variável 5"
	variavel6, variavel7 := "Variável 6", "Variável 7"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Printf("%T - %d\n", variavel8, variavel8)
	fmt.Println(variavel5)
	fmt.Println(variavel6)
	fmt.Println(variavel7)
	fmt.Println(constante1)
}
