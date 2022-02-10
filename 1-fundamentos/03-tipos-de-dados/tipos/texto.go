package tipos

import (
	print "tipos-de-dados/impressao"
)

func ImprimirTexto() {
	var str1 string = "Texto"
	str2 := "Texto 2"
	char := 'B'

	print.ImprimirCabecalho("TEXTO")
	print.ImprimirValor("str1", str1)
	print.ImprimirValor("str2", str2)
	print.ImprimirValor("char", char)
}
