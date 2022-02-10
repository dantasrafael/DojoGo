package main

import "fmt"

func calculosMatematicos(n1, n2 int8) (soma int8, subtracao int8) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Println(calculosMatematicos(10, 5))

	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)

	x, y := calculosMatematicos(1, 2)
	fmt.Println(x, y)
}
