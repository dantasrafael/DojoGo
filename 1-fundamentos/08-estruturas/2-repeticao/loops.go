package main

import (
	"fmt"
	"time"
)

func main() {
	numero := 0
	for numero < 5 {
		numero++
		fmt.Println("Incrementando número")
		time.Sleep(time.Second)
	}
	fmt.Println(numero)

	for numero2 := 0; numero2 < 10; numero2 += 2 {
		fmt.Println("Incrementando número2", numero2)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	for indice := range nomes {
		fmt.Println(indice)
	}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leo",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
