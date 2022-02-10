package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Printf("variavel1: %d (%v)\n", variavel1, &variavel1)
	fmt.Printf("variavel2: %d (%v)\n", variavel2, &variavel2)

	fmt.Println("\nSomando +1 na variavel 1")
	variavel1++
	fmt.Printf("variavel1: %d (%v)\n", variavel1, &variavel1)
	fmt.Printf("variavel2: %d (%v)\n", variavel2, &variavel2)

	fmt.Println("\n------------------------------")

	var variavel3 int
	var ponteiro *int // Ponteiro é uma referência de memória
	fmt.Printf("variavel3: %d (%v)\n", variavel3, &variavel3)
	fmt.Printf("ponteiro: %v\n", ponteiro)

	fmt.Println("\nAtribuindo o endereço da variavel3 ao ponteiro")
	ponteiro = &variavel3
	fmt.Printf("variavel3: %d (%v)\n", variavel3, &variavel3)
	fmt.Printf("ponteiro: %d (%v)\n", *ponteiro, ponteiro) // (*ponteiro) desreferenciação

	fmt.Println("\nAtribuindo o valor 150 a variavel3")
	variavel3 = 150
	fmt.Printf("variavel3: %d (%v)\n", variavel3, &variavel3)
	fmt.Printf("ponteiro: %d (%v)\n", *ponteiro, ponteiro) // (*ponteiro) desreferenciação
}
