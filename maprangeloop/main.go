package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Bonjour!",
		1: "Good day!",
		2: "你好!",
		3: "早安！",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
