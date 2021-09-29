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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, person := range m {
		fmt.Println(person.firstName)
		fmt.Println(person.lastName)

		for index, flavour := range person.favoriteIceCreamFlavors {
			fmt.Println(index, flavour)
		}
	}
}
