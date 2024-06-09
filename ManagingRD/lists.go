package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)
	prices[1] = 13.13

	// adding a new element to the slice
	prices = append(prices, 13.14)
	fmt.Println(prices)

	// removing an element from the slice
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{102.99, 8.99, 13.13}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.9
// 	highlightedPrices := featuredPrices[:1]
// 	// fmt.Println(featuredPrices)
// 	// fmt.Println(prices)
// 	// fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }
