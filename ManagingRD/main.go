package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)
	// userNames := []string{}

	userNames[0] = "James"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Yeat")

	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.7
	courseRatings["py"] = 5.3
	courseRatings["c++"] = 6.8

	fmt.Println(courseRatings)
}
