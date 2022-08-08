package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(3, 2)
	fmt.Println(soma)
	fmt.Println(somar(10, 2))

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	fmt.Println(f("Texto da função"))

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma2, _ := calculosMatematicos(5, 15)
	fmt.Println(resultadoSoma2)

	_, resultadoSubtracao2 := calculosMatematicos(5, 15)
	fmt.Println(resultadoSubtracao2)
}
