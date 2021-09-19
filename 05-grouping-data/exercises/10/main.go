package main

import "fmt"

func main() {
	nameToFavoriteThings := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "sunsets"},
	}

	nameToFavoriteThings["name"] = []string{"thing 1", "thing 2"}

	delete(nameToFavoriteThings, "no_dr")

	for name, favoriteThings := range nameToFavoriteThings {
		fmt.Printf("%v likes:\n", name)

		for index, thing := range favoriteThings {
			fmt.Println("\t", index, thing)
		}
	}
}
