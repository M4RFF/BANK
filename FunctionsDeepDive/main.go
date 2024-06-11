package main

import "fmt"

func main() {
	// numbers := []int{3, 13, 23}
	sum := sumup(3, 13, 23, 40, -5)

	fmt.Println(sum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}
	return sum
}
