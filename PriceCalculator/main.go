package main

import (
	"fmt"
	"price-calculator/cmdmanager"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.03, 0.1, 0.13}

	// result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxtIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
