package tipos

import (
	print "tipos-de-dados/impressao"
)

func ImprimirBooleano() {
	var booleano1 bool = true
	booleano2 := false

	print.ImprimirCabecalho("BOOLEANO")
	print.ImprimirValor("booleano1", booleano1)
	print.ImprimirValor("booleano2", booleano2)
}
