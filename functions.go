package main

import (
	"fmt"
	"math"
	"strings"
)

func greet(n string) {
	fmt.Printf("Good morning %v!\n", n)
}

func bye(n string) {
	fmt.Printf("Good Bye %v!\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	greet("Maestro")
	bye("revaz")
	cycleNames([]string{"Revaz", "Tamar", "Alex"}, greet)
	cycleNames([]string{"Revaz", "Tamar", "Alex"}, bye)
	a1 := circleArea(7.8)
	a2 := circleArea(15)
	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.2f \n", a1, a2)

	// #10
	fn, sn := getInitials("Revaz Gh")
	fmt.Println(fn, sn)
}

// muliple return values
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	words := strings.Split(s, " ")
	var initials []string
	for _, v := range words {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}
