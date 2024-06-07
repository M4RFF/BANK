package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"practice/note"
	"practice/todo"
)

// Creating the first Interface for Save()
type Saver interface {
	Save() error
}

// Creating the 2nd interface for Display()
type Displayer interface {
	Display()
}

// // create an interface that includes two methods
// type Outputtable interface {
// 	Save() error
// 	Display()
// }

// Creating an Embedded Interface
type Outputtable interface {
	Saver
	Displayer
}

func main() {
	printSomething(1)      // int
	printSomething(1.5)    // float
	printSomething("Maks") // string

	title, content := getNoteData()

	todoText := getUserImput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
}

// func getTodoData() string {
// 	text := getUserImput("Todo text:")
// 	return text
// }

// Special Kinder of Interface
func printSomething(value interface{}) { // any values are allowed
	fmt.Println(value)
}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserImput("Note title:")

	content := getUserImput("Note content:")

	return title, content
}

func getUserImput(prompt string) string {
	fmt.Printf("%v ", prompt)
	// var value string
	// fmt.Scanln(&value)
	// return value

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
