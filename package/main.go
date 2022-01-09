package main

import (
	"fmt"
)

// package scope
var score = 99.5

func main() {

	// function scope
	var score = 99.5

	sayHello("Revaz")

	for _, v := range points {
		fmt.Println(v)
	}
	showScore()
}

// go run main.go greetings.go
