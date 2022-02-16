package main

import (
	"fmt"
	ps "github.com/inancgumus/prettyslice"
)

func main() {
	fmt.Println("SLICES")

	ps.PrintBacking = true
	ps.PrintElementAddr = true
	ps.MaxPerLine = 10

	slice1 := []int{35, 15, 25}
	slice2 := slice1[:1]

	fmt.Printf("Tipo slice1: %T\n", slice1)
	ps.Show("slice1", slice1)
	fmt.Printf("Tipo slice2: %T\n", slice2)
	ps.Show("slice2", slice2)

	slice1 = append(slice1, 1)
	//slice2 = append(slice2, 99)

	ps.Show("slice1", slice1)
	//ps.Show("slice2", slice2)

	slice1 = append(slice1, 2, 3)
	ps.Show("slice1", slice1)
	slice1 = append(slice1, 4)
	ps.Show("slice1", slice1)
}
