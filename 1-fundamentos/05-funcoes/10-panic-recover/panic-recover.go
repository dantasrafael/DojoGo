package main

import (
	"fmt"
	"log"
	"regexp"
)

func recuperarExecucao() {
	fmt.Println("Tentando recuperar a execução!")
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float32) (bool, error) {
	defer recuperarExecucao()

	media := (n1 + n2) / 2

	if media > 6 {
		return true, nil
	} else if media < 6 {
		return false, nil
	}

	//return false
	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	result, err := alunoEstaAprovado(6, 6)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)
	fmt.Println("Pós execução")

	r, err := regexp.Compile("")
	fmt.Println(err)
	r = regexp.MustCompile(`^https?: `)
	fmt.Println(r)
}
