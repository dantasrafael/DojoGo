package main

import "fmt"

func function1() {
	fmt.Println("Function 1")
}

func function2() {
	fmt.Println("Function 2")
}

func isApprovedStudent(n1, n2 float32) bool {
	defer fmt.Printf("### Returning average = ")
	fmt.Println("### Calculating approved student")

	average := (n1 + n2) / 2
	if average >= 6 {
		return true
	}
	return false
}

func main() {
	defer function1()
	function2()

	fmt.Println(isApprovedStudent(7, 8))
}
