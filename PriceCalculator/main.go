package main

import (
	"price-calculator/prices"
)

func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.03, 0.1, 0.13}

	// result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxtIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
