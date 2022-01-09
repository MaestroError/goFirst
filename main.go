package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

}

func vars() {
	// will be an error if you don't use declared variable
	var name string = "Maestro"
	// auto asign`s type
	var name2 = "Revaz"
	// Default value - empty string
	var name3 string
	fmt.Println(name, name2, name3)

	// shorthand - can't use outside of function
	name4 := "Ghambara"
	fmt.Println(name4)

	// Bits & Memory
	var age int8 = 24

	// can't be lower than 0
	var age2 uint = 24

	var score1 float32 = 24.7
	// default type:
	var score2 float64 = 223423418343228324.3

	fmt.Println(age, age2, score1, score2)
}

func formatedStrings() {

	name := "Revaz"
	age := 24
	// Print
	fmt.Print("Hello,")
	fmt.Print(" World! \n")
	// Println
	fmt.Println("My name is ", name, ", I am ", age, " years old")
	// Printf (Formated strings) %_ = format specifier
	fmt.Printf("My name is %v, I am %v years old \n", name, age)
	fmt.Printf("My name is %q, I am %q years old \n", name, age)
	fmt.Printf("age is a type: %T \n", age)
	fmt.Printf("you scored %f \n", 225.25)
	fmt.Printf("you scored %0.1f \n", 225.25)

	// Sprintf (saved formated string)
	var str = fmt.Sprintf("My name is %v, I am %v years old \n", name, age)
	fmt.Println("saved string:", str)

}

func arraySlice() {

	// fixed lenght & type
	var ages [3]int = [3]int{24, 17, 23}
	var ages2 = [3]int{24, 17, 23}

	fmt.Println(ages, len(ages), ages2)

	names := [4]string{"MaestroError", "R.Gh.", "gh", "rg"}
	names[1] = "Revaz"
	fmt.Println(names, len(names))

	// Slices (Use arrays under the hood)
	var scores = []int{10, 20, 30}
	scores[1] = 25
	scores = append(scores, 40)

	fmt.Println(scores, len(scores))

	// slice ranges
	range1 := names[1:3]
	rangr2 := names[2:]
	range3 := names[:3]
	fmt.Println(range1, rangr2, range3)

	range1 = append(range1, "Tamar")
	fmt.Println(range1)
}

func standardLibrary() {

	greeting := "Hello My Bro!"
	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "l", "L"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "My"))
	fmt.Println(strings.Split(greeting, "My"))

	// original string
	fmt.Println("original string value:", greeting)

	ages := []int{1, 4, 2, 8, 7}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 7)
	// if index not founded it returns length of slice
	NotFoundIndex := sort.SearchInts(ages, 10)
	fmt.Println(index, NotFoundIndex)

	names := []string{"maestroError", "bGh.", "gh", "ag"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "gh"))
}
