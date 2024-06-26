package main

import (
	"fmt"
)

func main() {
	age := 19 // regular variable

	var agePointer *int

	agePointer = &age

	// agePointer := &age // the 2nd method how to initialize

	// fmt.Println("Age:", agePointer)  // it will output the address "0xc00000a0b8"
	fmt.Println("Age:", *agePointer) // it will output the value "19"

	// Now there's no copies of 19 its only one single "19" value
	// otherYears := getOtherYears(agePointer)
	// fmt.Println(otherYears)

	edutAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func edutAgeToAdultYears(age *int) {
	// return *age - 14
	*age = *age - 14 // we stored it in our pointer
}
