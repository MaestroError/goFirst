package main

import (
	"fmt"
	"math"
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

}

// muliple return values
func getInitials(n string) (string, string) {

}
