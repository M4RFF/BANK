package users

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

type Admin struct {
	email    string
	password string
	User
}

// I atteched the func bellow to the struct User above
func (u *User) OutPutUserDetails() { // (u User) receiver argument
	// ....

	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() { // I have to add "*" to the User if i wanna edit an original method not a copy
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "___",
			createdAt: time.Now(),
		},
	}
}

// creation or constructor func
func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" { // checks an input from the user if he skips smth then
		return nil, errors.New("First name, Last name, Birthdata are required") // it will output an error
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}
