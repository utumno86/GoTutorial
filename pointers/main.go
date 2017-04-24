package main

import "fmt"

func main() {
	var a int
	fmt.Print("What is a: ")
	fmt.Scan(&a)
	fmt.Println(a)
	fmt.Println(&a)

	var b = &a
	fmt.Println(b)
	fmt.Println(*b)
}
