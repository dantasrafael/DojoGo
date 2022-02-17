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
	curso  string
	escola string
}

func (p *pessoa) salvar() {
	fmt.Printf("Salvando os dados da pessoa '%s' no SGBD\n", p.nome)
}

func (p *pessoa) maiorDeIdade() bool {
	return p.idade >= 18
}

func (p *pessoa) fazerAniversario() {
	p.idade++
}

func main() {
	p1 := pessoa{"João", "Santos", 17, 1.58}
	fmt.Println(p1)
	fmt.Printf("%s é maior de idade? %v\n", p1.nome, p1.maiorDeIdade())
	p1.salvar()

	p1.fazerAniversario()
	fmt.Println("João fez aniversário")
	fmt.Printf("%s é maior de idade? %v\n\n", p1.nome, p1.maiorDeIdade())

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.maiorDeIdade())
	e1.salvar()
}
