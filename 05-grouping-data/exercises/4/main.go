package main

import "fmt"

func main() {
	firstArray := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	secondArray := []int{56, 57, 58, 59, 60}

	firstArray = append(firstArray, 52)

	fmt.Println(firstArray)

	firstArray = append(firstArray, 53, 54, 55)

	fmt.Println(firstArray)

	firstArray = append(firstArray, secondArray...)

	fmt.Println(firstArray)
}
