package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"practice/note"
	"practice/todo"
)

// Creating the new Interface
type Saver interface {
	Save() error
}

func main() {
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

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println("Saving the todo failed")
		return
	}
	fmt.Println("Saving the note succeeded!")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note succeeded!")
}

// func getTodoData() string {
// 	text := getUserImput("Todo text:")
// 	return text
// }

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
