package main

import (
	"fmt"

	"example.com/bank/fileops" // i need to find how to solve this error with importing package "fileops" into bank.go
	"github.com/Pallinder/go-randomdata"
)

const accoundBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accoundBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("_-_-_-_-__-_-_-_-_-_-_-_-_-_")
		// return // it stops a program when it cant find an txt file with information about a balance.

		// more drastic way to exit this application
		// panic("Can't continue, sorry!")
	}

	fmt.Println("Welcome to Max Bank")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {

		/*
			fmt.Println("What do you want to do?")
			fmt.Println("1. Check balance")
			fmt.Println("2. Deposit money")
			fmt.Println("3. Withdraw money")
			fmt.Println("4. Exit")
		*/

		presentOptions()

		var choice int
		fmt.Print("Enter: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:

			fmt.Println("Your balance is", accountBalance)

		case 2:

			fmt.Print("How much money do you want to deposite? Plz Enter: ")
			var depositeAmount float64
			fmt.Scan(&depositeAmount)

			if depositeAmount <= 0 {
				fmt.Println("Invalid amout. Must be greater than 0")
				// return // it doesnt give a user the second try, but we wanna give it to him
				continue
			}

			accountBalance += depositeAmount // accoundBalance = accoundBalance + depositeAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accoundBalanceFile)

		case 3:

			fmt.Print("How much money do you want to withdraw? Enter: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			if withdrawAmount <= accountBalance {
				accountBalance -= withdrawAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
				fileops.WriteFloatToFile(accountBalance, accoundBalanceFile)

			} else if withdrawAmount > accountBalance {
				fmt.Println("You don't have enough money on the balance!")
				continue
			}

		default:

			fmt.Println("Goodbye!")
			// return
			// break // not longer breaks out of switch statment
			fmt.Println("Thanks for choosing our bank!")
			return
		}

		// It is before I learned switch case to get rid of the repeating else if statments
		/*
			if choice == 1 {
				fmt.Println("Your balance is", accountBalance)
			} else if choice == 2 {
				fmt.Print("How much money do you want to deposite? Plz Enter: ")
				var depositeAmount float64
				fmt.Scan(&depositeAmount)

				if depositeAmount <= 0 {
					fmt.Println("Invalid amout. Must be greater than 0")
					// return // it doesnt give a user the second try, but we wanna give it to him
					continue
				}

				accountBalance += depositeAmount // accoundBalance = accoundBalance + depositeAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
			} else if choice == 3 {
				fmt.Print("How much money do you want to withdraw? Enter: ")
				var withdrawAmount float64
				fmt.Scan(&withdrawAmount)

				if withdrawAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0")
					return
				}

				if withdrawAmount <= accountBalance {
					accountBalance -= withdrawAmount
					fmt.Println("Balance updated! New amount:", accountBalance)

				} else if withdrawAmount > accountBalance {
					fmt.Println("You don't have enough money on the balance!")
					return
				}
			} else {
				fmt.Println("Goodbye!")
				// return
				break // only in loop it breaks out of the loop
			}
		*/
	}

	// fmt.Println("Thanks for choosing our bank!") // only for old version

}
