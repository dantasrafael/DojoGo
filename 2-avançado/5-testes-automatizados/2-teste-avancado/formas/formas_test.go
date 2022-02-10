package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}

		areaEsperada := float64(120)
		areaRecebida := ret.area()

		if areaEsperada != areaRecebida {
			t.Fatalf("À área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		crc := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := crc.area()

		if areaEsperada != areaRecebida {
			t.Fatalf("À área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}
