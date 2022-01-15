package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Print(menu["pie"])

	// looping maps
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// ints as key type
	phonebook := map[int]string{
		21331231: "Mario",
		48567452: "Kote",
		78732214: "Maddona",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[78732214])

	phonebook[48567452] = "butkhuza"
	fmt.Println(phonebook)

	phonebook[78732214] = "Mad"
	fmt.Println(phonebook)

}
