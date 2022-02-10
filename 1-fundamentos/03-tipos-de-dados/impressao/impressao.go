package impressao

import "fmt"

func ImprimirCabecalho(nome string) {
	fmt.Printf("\n ----- %s (variável | tipo | valor) -----\n", nome)
}

func ImprimirValor(nome string, variavel interface{}) {
	fmt.Printf("|%15s", nome)
	fmt.Printf("|%20T", variavel)
	fmt.Printf("|%v\t\n", variavel)
}
