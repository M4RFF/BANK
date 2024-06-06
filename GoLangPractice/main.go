package main

import (
	"bufio"
	"fmt"
	"os"
	"practice/note"
	"strings"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the nore failed")
		return
	}

	fmt.Println("Saving the note succeeded")
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
