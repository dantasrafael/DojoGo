package main

import "fmt"

func somar(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}

}

func main() {
	fmt.Println(somar(1, 2, 3))
	escrever("Número: ", 10, 250, 60)
}
