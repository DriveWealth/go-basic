package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	result := sum(5, 7)
	fmt.Println("Sum=", result)

	//result1, err := sqrt(16)
	result1, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result1)
	}

}

// Function Example
func sum(x int, y int) int {
	return x + y
}

// Function with multiple return type
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for -ve numbers")
	}
	return math.Sqrt(x), nil
}
