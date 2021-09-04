package main

import "fmt"

type myType int

func main() {
	var x myType

	fmt.Println(x)
	fmt.Printf("type %T\n", x)

	x = 42

	fmt.Println(x)
}
