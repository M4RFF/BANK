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
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_") // replacing
	fileName = strings.ToLower(fileName) + ".json"       // all now contains lowercase charachters

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
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
