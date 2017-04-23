package main

import "fmt"

const (
	_ = iota
	// Mb Megabytes
	Mb = 1 << (iota * 10)
	//Kb Kilobytes
	Kb = 1 << (iota * 10)
	//Gb Gigabytes
	Gb = 1 << (iota * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", Mb)
	fmt.Printf("%d\n", Mb)
	fmt.Printf("%b\t", Kb)
	fmt.Printf("%d\n", Kb)
	fmt.Printf("%b\t", Gb)
	fmt.Printf("%d\n", Gb)
}
