package main

import "fmt"

func main() {
	n := greatest(43, 56, 87, 45, 57)
	fmt.Println(n)
}

func greatest(num ...int) int {
	highNum := 0
	for _, v := range num {
		if v > highNum {
			highNum = v
		}
	}
	return highNum
}
