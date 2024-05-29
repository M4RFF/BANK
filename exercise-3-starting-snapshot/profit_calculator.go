package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
// => Show error message & exit if valid input is provided
// -No negative numbers
// -Not 0
// Done
// 2) Show calculated results into the file

// ex 2
// create a file and writing in into the file called "results.txt"

const resultsValueFile = "results.txt"

func writeCalculatingsResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ratio)
	os.WriteFile(resultsValueFile, []byte(results), 0644)
}

func main() {
	revenue, err1 := getUserInput("Revenue: ")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	expenses, err2 := getUserInput("Expenses: ")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	taxRate, err3 := getUserInput("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	writeCalculatingsResults(ebt, profit, ratio)

}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio

}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a possitive number")
	}

	return userInput, nil
}
