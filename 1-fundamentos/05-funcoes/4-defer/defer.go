package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
}

func funcao2() {
	fmt.Println("Função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("### Média calculada. O retorno será retornado")
	fmt.Println("### Entra na função aprovador")

	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	defer funcao1() // DEFER = adiar
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
