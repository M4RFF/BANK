package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)
	// userNames := []string{}

	userNames[0] = "James"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Yeat")

	fmt.Println(userNames)
}
