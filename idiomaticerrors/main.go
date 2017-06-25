package main

import (
	"fmt"
	"log"
	"math"
)

// ErrImaginaryNumber generates an error when the output would be imaginary
type ErrImaginaryNumber struct {
	val string
	err error
}

func (n *ErrImaginaryNumber) Error() string {
	return fmt.Sprintf("An Imaginary Number error occured: %v %v", n.val, n.err)
}

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

// Sqrt function returns a square root or error
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		err := fmt.Errorf("square root of negative number: %v", f)
		return 0, &ErrImaginaryNumber{"value", err}
	}
	return math.Sqrt(f), nil
}
