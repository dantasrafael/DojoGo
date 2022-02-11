package main

import (
	"fmt"
	ps "github.com/inancgumus/prettyslice"
)

func main() {
	fmt.Println("ARRAYS")

	var array1 [5]int
	array1[0] = 1
	ps.Show("array1", array1)

	array2 := [5]int{10, 20, 30, 40, 50}
	ps.Show("array2", array2)

	array3 := [...]int{100, 200, 300}
	ps.Show("array3", array3)
}
