package main

import (
	"fmt"
)

func main() {

	//loop: Only one type loop is in Go i.e. for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//Also used as while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//Use of range
	fmt.Println("String Array")
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
	//Use Range with map
	fmt.Println("String Map")
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"
	for key, value := range m {
		fmt.Println("index:", key, "value:", value)
	}
}
