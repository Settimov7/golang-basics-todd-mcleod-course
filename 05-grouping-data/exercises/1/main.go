package main

import "fmt"

func main() {
	array := [5]int{1, 3, 8, 14, 17}

	for _, value := range array {
		fmt.Println(value)
	}

	fmt.Printf("type %T", array)
}
