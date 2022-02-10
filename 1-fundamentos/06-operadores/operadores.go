package main

import "fmt"

func main() {
	// ARITMETICOS (+ - / * %)
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// var numero1 int16 = 10
	// var numero2 int32 = 20
	// soma2 := numero1 + numero2
	// fmt.Println(soma)

	// ATRIBUIÇÃO
	fmt.Println("-------------")
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// RELACIONAIS (true/false)
	fmt.Println("-------------")
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// LÓGICOS (true/false)
	fmt.Println("-------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(falso && verdadeiro)
	fmt.Println(verdadeiro || falso)
	fmt.Println(falso || verdadeiro)
	fmt.Println(!verdadeiro && falso)
	fmt.Println(!verdadeiro || falso)

	// UNÁRIOS
	fmt.Println("-------------")
	numero := 10
	fmt.Println(numero)

	numero++
	fmt.Println(numero)

	numero--
	fmt.Println(numero)

	numero += 15
	fmt.Println(numero)

	numero -= 20
	fmt.Println(numero)

	numero *= 3
	fmt.Println(numero)

	numero /= 10
	fmt.Println(numero)

	numero %= 3
	fmt.Println(numero)

	// TERNÁRIOS (não existe em Go)
	// texto := numero > 5 ? "Maior que 5" : "Menor que 5"
}
