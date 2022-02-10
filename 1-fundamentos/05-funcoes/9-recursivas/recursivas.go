package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	// 1 1 2 3 5 8 13
	posicao := uint(10)
	fmt.Println(fibonacci(posicao))

	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(6))

	for i := uint(1); i < posicao; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}

}
