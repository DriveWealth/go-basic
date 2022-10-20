package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {

	//Struct Data Type
	p := person{name: "Ram", age: 20}
	fmt.Println(p)
	fmt.Println(p.name)
}
