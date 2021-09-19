package main

import "fmt"

func main() {
	array := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(array[:5])
	fmt.Println(array[5:])
	fmt.Println(array[2:7])
	fmt.Println(array[1:6])
}
