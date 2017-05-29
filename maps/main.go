package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	// delete(m, "k2")
	fmt.Println("map:", m)

	_, ok := m["k2"]
	fmt.Println("Ok?: ", ok)
}
