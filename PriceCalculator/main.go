package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.3, 0.1, 0.13}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		taxIncludePrice := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludePrice[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludePrice
	}

	fmt.Println(result)
}
