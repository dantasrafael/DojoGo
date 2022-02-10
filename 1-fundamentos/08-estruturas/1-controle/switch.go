package main

import "fmt"

func diaDaSemana(numero uint8) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sabado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero uint8) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana(2))
	fmt.Println(diaDaSemana(3))
	fmt.Println(diaDaSemana(4))
	fmt.Println(diaDaSemana(5))
	fmt.Println(diaDaSemana(6))
	fmt.Println(diaDaSemana(7))
	fmt.Println(diaDaSemana(0))
	fmt.Println(diaDaSemana(8))

	fmt.Println("------------------")

	fmt.Println(diaDaSemana2(1))
	fmt.Println(diaDaSemana2(2))
	fmt.Println(diaDaSemana2(3))
	fmt.Println(diaDaSemana2(4))
	fmt.Println(diaDaSemana2(5))
	fmt.Println(diaDaSemana2(6))
	fmt.Println(diaDaSemana2(7))
	fmt.Println(diaDaSemana2(0))
	fmt.Println(diaDaSemana2(8))
}
