package main

import "fmt"

type person struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	p1 := person{
		"Name 1",
		"Last Name 1",
		[]string{
			"Chocolate",
			"Mint",
			"Strawberry",
		},
	}

	p2 := person{
		"Name 2",
		"Last Name 2",
		[]string{
			"Chocolate",
			"Strawberry",
		},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)

	for index, flavour := range p1.favoriteIceCreamFlavors {
		fmt.Println(index, flavour)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)

	for index, flavour := range p2.favoriteIceCreamFlavors {
		fmt.Println(index, flavour)
	}
}
