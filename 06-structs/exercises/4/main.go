package main

import "fmt"

func main() {
	v := struct {
		field1 string
		field2 int
	}{
		field1: "Field 1",
		field2: 2,
	}

	fmt.Println(v)
	fmt.Println(v.field1)
	fmt.Println(v.field2)
}
