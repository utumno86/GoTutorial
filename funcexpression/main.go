package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hello there!")
	}
	greeting()
}
