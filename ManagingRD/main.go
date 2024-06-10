package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	// userNames := []string{}

	userNames[0] = "James"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Yeat")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["py"] = 5.3
	courseRatings["c++"] = 6.8

	courseRatings.output()

	// fmt.Println(courseRatings)
}
