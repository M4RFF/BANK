package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error { // here save() doesn't have any inputs but return error
	fileName := "todo.json"

	json, err := json.Marshal(todo) // converts to text json format

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// constructor function
func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}
