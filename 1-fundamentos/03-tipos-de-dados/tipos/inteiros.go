package tipos

import (
	print "tipos-de-dados/impressao"
)

func ImprimirInteiros() {
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64 (SEM SINAL)
	// int = usa a arquitetura da máquina como base
	// alias = rune (int32) | byte (uint8)
	var numero1 int16 = 1000
	var numero2 uint8 = 100
	var numero3 rune = 123456
	var numero4 byte = 123
	numero5 := 100

	print.ImprimirCabecalho("NUMÉROS INTEIROS")
	print.ImprimirValor("numero1", numero1)
	print.ImprimirValor("numero2", numero2)
	print.ImprimirValor("numero3", numero3)
	print.ImprimirValor("numero4", numero4)
	print.ImprimirValor("numero5", numero5)
}
