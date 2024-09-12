package file

import (
	"errors"
	"os"
)

func GetFileContent() ([]byte, error) {
	fileName := os.Args[1]

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("There was a problem reading the source file.")
	}

	return data, nil
}
