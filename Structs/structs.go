package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// I atteched the func bellow to the struct User above
func (u *User) outPutUserDetails() { // (u User) receiver argument
	// ....

	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) clearUserName() { // I have to add "*" to the User if i wanna edit an original method not a copy
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser User

	appUser = User{ // Instance of a struct type "User"
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	appUser.outPutUserDetails() // called the func that I've atteched to the struct
	appUser.clearUserName()
	appUser.outPutUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
