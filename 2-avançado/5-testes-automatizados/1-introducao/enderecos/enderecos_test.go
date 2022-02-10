package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça dos 46", "Tipo inválido"},
		{"Estrada do Trabalhador", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido '%s' é diferente do esperado '%s'", tipoDeEnderecoRecebido, cenario.enderecoEsperado)
		}
	}
}
