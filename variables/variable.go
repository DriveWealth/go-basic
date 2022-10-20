package main

import (
	"fmt"
)

func main() {
	//Sample Variables
	var x int = 5
	var y int = 7
	var sum int = x + y
	fmt.Println(sum)

	//other way to define Variables
	a := 5
	b := 7
	sum1 := a + b
	fmt.Println(sum1)
	//If else control Statement
	if x > y {
		fmt.Println("x is greater than y", x)
	} else {
		fmt.Println("Y is greater than x", y)
	}
	//Array
	var z [5]int
	z[2] = 7
	fmt.Println(z)

	k := [5]int{5, 4, 3, 2, 1}
	fmt.Println(k)

	// Slice
	l := []int{1, 2, 3, 4, 5}
	l = append(l, 12)
	fmt.Println(l)

}
