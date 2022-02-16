package main

import (
	"fmt"
)

func main() {
	fmt.Println("ARRAYS")

	var array1 [5]int
	array1[0] = 1
	fmt.Printf("Tipo da array1: %T\n", array1)
	fmt.Println("array1", array1)

	array2 := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Tipo da array2: %T\n", array2)
	fmt.Println("array2", array2)

	array3 := [...]int{100, 200, 300, 400, 500}
	fmt.Printf("Tipo da array3: %T\n", array3)
	fmt.Println("array3", array3)
}
