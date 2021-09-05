package main

import "fmt"

func main() {
	const (
		a = iota + 200
		b
		c
		d
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
