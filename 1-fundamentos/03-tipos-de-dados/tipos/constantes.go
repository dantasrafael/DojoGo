package tipos

import (
	print "tipos-de-dados/impressao"
)

const (
	ANO        = uint16(2022)
	PI         = float32(3.14)
	TEXTO      = "Esta é uma constante de texto"
	VERDADEIRO = true
)

func ImprimirConstantes() {
	print.ImprimirCabecalho("CONSTANTES")
	print.ImprimirValor("ano", ANO)
	print.ImprimirValor("pi", PI)
	print.ImprimirValor("texto", TEXTO)
	print.ImprimirValor("verdadeiro", VERDADEIRO)
}
