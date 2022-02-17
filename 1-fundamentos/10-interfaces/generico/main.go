package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Teste")
	generica(10)
	generica(false)
	generica(10.4)

	mapa := map[string]interface{}{
		"nome":  "Pedro",
		"idade": 20,
		"peso":  70.5,
	}
	fmt.Println(mapa)
}
