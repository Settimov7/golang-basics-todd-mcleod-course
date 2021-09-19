package main

import "fmt"

func main() {
	sliceOfSliceOfString := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James"},
	}

	for _, sliceOfString := range sliceOfSliceOfString {
		for _, value := range sliceOfString {
			fmt.Println(value)
		}
	}
}
