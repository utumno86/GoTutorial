package main

import "fmt"

func main() {
	a := 10
	b := "string"
	c := 4.2
	d := false

	fmt.Printf("%T %v \n", a, a)
	fmt.Printf("%T %v \n", b, b)
	fmt.Printf("%T %v \n", c, c)
	fmt.Printf("%T %v \n", d, d)
}
