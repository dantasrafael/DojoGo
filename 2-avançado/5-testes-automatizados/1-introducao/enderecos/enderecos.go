package enderecos

import "strings"

// TipoDeEndereco verifica se um endereco tem um tipo válido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		return strings.Title(primeiraPalavraEndereco)
	}
	return "Tipo inválido"
}
