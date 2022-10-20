package main

import (
	"fmt"
)

func main() {

	//Map Data Type
	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 4
	fmt.Println(vertices)
	fmt.Println(vertices["square"])
	delete(vertices, "square")
	fmt.Println(vertices)

}
