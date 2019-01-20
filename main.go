package main

import (
	"./proto1"
	"./proto2"
	"fmt"
)

func main() {
	fmt.Printf("start\n")
	
	e := proto1.Enum1{}
	fmt.Printf("enum 1: %v\n", e)
	
	e2 := proto2.Enum2{}
	fmt.Printf("enum 2: %v\n", e2)
}
