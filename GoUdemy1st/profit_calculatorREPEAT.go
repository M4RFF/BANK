/*
package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Revenue: ")
	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	expenses = getUserInput("Expenses: ")
	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	taxRate = getUserInput("Tax Rate: ")
	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	ebt, profit, ration := solvings(revenue, expenses, taxRate)
	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ration := ebt / profit

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ration)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput) // user input a number
	return userInput
}

func solvings(revenue, expenses, taxRate float64) (ebt float64, profit float64, ration float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ration = ebt / profit
	return ebt, profit, ration
}
*/