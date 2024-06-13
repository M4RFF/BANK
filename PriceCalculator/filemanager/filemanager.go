package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path) // allows us to open this file

	if err != nil {
		// fmt.Println("could not open a file")
		// fmt.Println(err)
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file) // allows us to go line by line in this "file"

	// we move farword one line of the line where we were
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// fmt.Println("reading the file content failed")
		// fmt.Println(err)
		file.Close()
		return nil, errors.New("failed to read in file")
	}

	file.Close()
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to creare a file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	file.Close()
	return nil
}
