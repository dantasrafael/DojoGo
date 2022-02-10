package tipos

import (
	print "tipos-de-dados/impressao"
)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func ImprimirStruct() {
	var usuario1 usuario

	var usuario2 usuario
	usuario2.nome = "Usuário 2"
	usuario2.idade = 10

	usuario3 := usuario{"Usuário 3", 5, endereco{"Rua do Centro", 1}}

	usuario4 := usuario{
		nome:  "Usuário 4",
		idade: 37,
	}

	usuario5 := usuario{
		nome: "Usuário 5",
	}

	usuario6 := usuario{
		idade: 11,
	}

	print.ImprimirCabecalho("STRUCT")
	print.ImprimirValor("usuario1", usuario1)
	print.ImprimirValor("usuario2", usuario2)
	print.ImprimirValor("usuario3", usuario3)
	print.ImprimirValor("usuario4", usuario4)
	print.ImprimirValor("usuario5", usuario5)
	print.ImprimirValor("usuario6", usuario6)
}
