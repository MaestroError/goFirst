package main

import (
	"fmt"
)

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateName2(x *string) {
	*x = "wedge"
}

func main() {
	name := "Tifa"

	updateName(name)

	fmt.Println("memory address of name is: ", &name)

	// & - gets pointer
	m := &name
	fmt.Println("memory address:", m)

	// * - get value
	fmt.Println("Value at memory adress:", *m)

	fmt.Println(name)
	updateName2(m)
	fmt.Println(name)
}
