package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"João", "Santos", 20, 1.78}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)

	e2 := estudante{pessoa{"Tiago", "Barreto", 18, 1.70}, "Medicina", "UCSal"}
	fmt.Println(e2)
	fmt.Println(e2.nome)
}
