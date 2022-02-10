package tipos

import (
	print "tipos-de-dados/impressao"
)

type Regioes int

const (
	Norte = iota
	Sul
	Leste
	Oeste
)

func (r Regioes) String() string {
	return [...]string{"Norte", "Sul", "Leste", "Oeste"}[r]
}

func ImprimirEnumerado() {
	var regiao1 Regioes = Norte
	var regiao2 Regioes = Sul
	var regiao3 Regioes = Leste
	var regiao4 Regioes = Oeste

	print.ImprimirCabecalho("ENUMERADO")
	print.ImprimirValor("regiao1", regiao1)
	print.ImprimirValor("regiao2", regiao2)
	print.ImprimirValor("regiao3", regiao3)
	print.ImprimirValor("regiao4", regiao4)
}
