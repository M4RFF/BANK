package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	// var investmentAmount, years float64 = 1000, 10
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futuureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futuureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value:", futureValue)

	//fmt.Printf("Future Value: %v\n", futureValue)
	//fmt.Println("Future Value (adjusted for Inflation):", futuureRealValue)

	// fmt.Printf(`Future Value: %.1f
	// Future Value (adjusted for Inflation): %.1f`, futureValue, futuureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (afjasred for Inflation): %.1f\n", futuureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

// Functions

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv
	return
}
