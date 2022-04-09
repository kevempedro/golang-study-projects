// Teste de unidade

package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua Oito", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"AVENIDA Paulista", "Avenida"},
		{"Estrada 66", "Estrada"},
		{"Rodovia Oito", "Rodovia"},
		{"Praça da Estação", "Tipo Inválido"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if cenario.retornoEsperado != tipoDeEnderecoRecebido {
			t.Errorf("O tipo de endereço %s é diferente do tipo de endereço esperado %s",
				tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}
