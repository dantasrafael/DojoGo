package main

import "fmt"

func closure() func() {
	texto := "Dentro da função Closure"

	return func() {
		fmt.Println(texto)
	}
}

func main() {
	texto := "Dentro da função Main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
