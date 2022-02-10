package tipos

import (
	print "tipos-de-dados/impressao"
)

func ImprimirValorZero() {
	var texto string
	var numero int16
	var real float32
	var booleano bool
	var erro error

	print.ImprimirCabecalho("VALOR ZERO")
	print.ImprimirValor("texto", texto)
	print.ImprimirValor("numero", numero)
	print.ImprimirValor("real", real)
	print.ImprimirValor("booleano", booleano)
	print.ImprimirValor("erro", erro)
}
