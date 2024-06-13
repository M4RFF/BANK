package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath) // allows us to open this file

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

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

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

func New(inputpPath, outputPath string) FileManager { // new file manager value
	return FileManager{
		InputFilePath:  inputpPath,
		OutputFilePath: outputPath,
	}
}
