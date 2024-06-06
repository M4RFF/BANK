package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.title, note.content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.title, " ", "_") // replacing
	fileName = strings.ToLower(fileName)                 // all now contains lowercase charachters

	json, err := json.Marshal(note) // converts to text json format

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// constructor function
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
