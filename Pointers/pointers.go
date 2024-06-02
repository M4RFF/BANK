package main

import (
	"fmt"
)

func main() {
	age := 19 // regular variable

	fmt.Println("Age: ", age)

	otherYears := getOtherYears(age)
	fmt.Println(otherYears)
}

func getOtherYears(age int) int {
	return age - 14
}
