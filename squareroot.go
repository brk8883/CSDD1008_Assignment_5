package main

import (
	"errors"
	"fmt"
	"math"
)

// Sqrt function
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot calculate square root of negative number")
	}
	return math.Sqrt(a), nil
}

func main() {

	var a float64
	fmt.Println("Enter a number:")
	fmt.Scanln(&a)
	result, err := Sqrt(a)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Square root: %.2f\n", result)
	}

}
