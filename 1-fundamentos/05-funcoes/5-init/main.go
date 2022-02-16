package main

import "fmt"

var n int

func init() {
	fmt.Println("Starting when loading package")
	n = 10
}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(n)
}
