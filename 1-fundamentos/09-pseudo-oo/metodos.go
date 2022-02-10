package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário '%s' no SGBD\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Manoel", 32}
	usuario1.salvar()

	usuario2 := usuario{"Maria", 17}
	usuario2.salvar()

	fmt.Println(usuario2.maiorDeIdade())

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
	fmt.Println(usuario2.maiorDeIdade())
}
