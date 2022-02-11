package main

import (
	"fmt"
)

func main() {
	fmt.Println("MAPS")

	m := map[string]interface{}{"k1": "v1", "k2": "v2"}

	m["k3"] = "v3"

	fmt.Println(m)
}
