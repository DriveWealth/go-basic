package main

import (
	"fmt"
)

func main() {

	//Pointer Data Type
	i := 3
	fmt.Println(i)
	fmt.Println(&i)
	inc(&i)
	fmt.Println(i)
}

func inc(x *int) {
	*x++
}
