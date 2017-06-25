package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

// ErrImaginaryNumber generates an error when the output would be imaginary
var ErrImaginaryNumber = errors.New("Imaginary Number: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrImaginaryNumber)
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

// Sqrt function returns a square root or error
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrImaginaryNumber
	}
	return math.Sqrt(f), nil
}
