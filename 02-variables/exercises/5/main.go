package main

import "fmt"

type myType int

var x myType
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("x type %T\n", x)

	x = 42

	fmt.Println(x)

	y = int(x)

	fmt.Println(y)
	fmt.Printf("y type %T\n", y)
}
