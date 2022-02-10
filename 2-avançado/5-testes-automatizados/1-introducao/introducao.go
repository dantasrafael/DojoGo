package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	fmt.Println(enderecos.TipoDeEndereco("Rua dos Imigrantes"))
	fmt.Println(enderecos.TipoDeEndereco("Estrada dos Imigrantes"))
	fmt.Println(enderecos.TipoDeEndereco("Avenida dos Imigrantes"))
	fmt.Println(enderecos.TipoDeEndereco("Rodovia dos Imigrantes"))
	fmt.Println(enderecos.TipoDeEndereco("Travessa dos Imigrantes"))
}
