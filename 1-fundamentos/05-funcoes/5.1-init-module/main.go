package main

import (
	"example/pkg"
	"fmt"
)

func init() {
	fmt.Println("init in main package")
}

func main() {
	fmt.Println("5.1-init-module")
	pkg.PrintSomething()
}
