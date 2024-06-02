package main

import (
	"fmt"
)

func main() {
	age := 19 // regular variable

	var agePointer *int
	agePointer = &age

	// agePointer := &age // the 2nd method how to initialize

	// fmt.Println("Age:", agePointer)  // it will output the address 0xc00000a0b8
	fmt.Println("Age:", *agePointer) // it will output the value

	// otherYears := getOtherYears(age)
	// fmt.Println(otherYears)
}

func getOtherYears(age int) int {
	return age - 14
}
