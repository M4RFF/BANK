package main

import (
	"fmt"
)

type TransformFn func(int) int

// type AnotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double) // use &numbers to get the address of numbers
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(numbers *[]int, transform TransformFn) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
