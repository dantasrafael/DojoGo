package tipos

import (
	print "tipos-de-dados/impressao"
)

func ImprimirReais() {
	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 123000000000.45
	numeroReal3 := 123456.7

	print.ImprimirCabecalho("NUMÉROS REAIS")
	print.ImprimirValor("numeroReal1", numeroReal1)
	print.ImprimirValor("numeroReal2", numeroReal2)
	print.ImprimirValor("numeroReal3", numeroReal3)
}
