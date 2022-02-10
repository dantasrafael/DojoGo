package tipos

import (
	"errors"
	print "tipos-de-dados/impressao"
)

func ImprimirErro() {
	var error1 error = errors.New("Erro interno")

	print.ImprimirCabecalho("ERROR")
	print.ImprimirValor("error1", error1)
}
