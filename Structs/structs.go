package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// creation or constructor func

func newUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, Last name, Birthdata are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
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

	var appUser *User

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}
	// ... do something awesome with that gathered data!

	appUser.outPutUserDetails() // called the func that I've atteched to the struct
	appUser.clearUserName()
	appUser.outPutUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value) // Scanln() stops waiting for an input when a user skips inputs
	return value
}
