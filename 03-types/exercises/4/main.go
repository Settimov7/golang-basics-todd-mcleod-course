package main

import "fmt"

func main() {
	a := 228

	fmt.Printf("binary %b\t\tdecimal %d\thex %X\n", a, a, a)

	a = a << 1

	fmt.Printf("binary %b\tdecimal %d\thex %X\n", a, a, a)
}
