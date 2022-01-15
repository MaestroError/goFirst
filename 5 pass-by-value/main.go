package main

import (
	"fmt"
)

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["cofffee"] = 2.99
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "Tifa"

	updateName(name)

	// prints Tifa, coz variable copy is passing
	// in function and not the original one, see comp_memory.png
	fmt.Println(name)

	name = updateName(name)
	fmt.Println(name)

	// group B types -> slices, maps, functions
	menu := map[string]float64{
		"pie":       7.99,
		"ice cream": 6.99,
	}

	updateMenu(menu)

	// it updates original value, coz variable menu stores pointer
	// for another data block and copied pointer is still points to original data block, see pointers.png
	fmt.Println(menu)

	// A - Non-Pointer Values
	// B - Pointer Wrapper Values
}
