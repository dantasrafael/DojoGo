package main

import "fmt"

func main() {

	func(text string) {
		fmt.Println(text)
	}("Hello World!")

	func() {
		fmt.Println("anonymous function without parameter...")
	}()

	result := func(text string) string {
		return fmt.Sprintf("Param => %s\n", text)
	}

	fmt.Println(result("param value"))
}
