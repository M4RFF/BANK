package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}
	// ... do something awesome with that gathered data!

	appUser.OutPutUserDetails() // called the func that I've atteched to the struct
	appUser.ClearUserName()
	appUser.OutPutUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value) // Scanln() stops waiting for an input when a user skips inputs
	return value
}
