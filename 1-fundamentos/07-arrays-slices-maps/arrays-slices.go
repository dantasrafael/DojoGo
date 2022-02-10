package main

import "fmt"

func main() {
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(array2)

	array3 := [...]int{100, 200, 300}
	fmt.Println(array3)

	slice1 := []int{10, 11, 12, 13}
	fmt.Println(slice1)

	slice1 = append(slice1, 14)
	fmt.Println(slice1)

	fmt.Printf("%T, %T\n", array1, slice1)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = 999
	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("-------")
	slice3 := make([]float32, 10, 11) // Criou um array interno de 11 posições com um slice de 10 posições
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // tamanho
	fmt.Println(cap(slice4)) // capacidade

	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // tamanho
	fmt.Println(cap(slice4)) // capacidade
}
