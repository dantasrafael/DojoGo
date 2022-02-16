package main

import "fmt"

func main() {
	usuario1 := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario1)
	fmt.Println(usuario1["nome"])
	fmt.Println(usuario1["sobrenome"])

	delete(usuario1, "sobrenome")
	fmt.Println(usuario1)
	fmt.Println(usuario1["nome"])
	fmt.Println(usuario1["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nomes": {
			"primeiro": "João",
			"último":   "Pedro",
		},
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nomes"])
	fmt.Println(usuario2["nomes"]["primeiro"])

	usuario2["cursos"] = map[string]string{
		"nome":   "Medicina",
		"campus": "Centro",
	}
	fmt.Println(usuario2)
	fmt.Println(usuario2["nomes"])
	fmt.Println(usuario2["nomes"]["primeiro"])
}
