// CHALLENGE #1:
// -- Use goroutines and channels to calculate 100 factorial computations conncurrently and in parellel

package main

import "fmt"

func main() {
	in := gen()

	f := factorial(in)

	for n := range f {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i < 101; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
