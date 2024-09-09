package compiler

import (
	"errors"
	"os"
)

func Run() error {
	err := checkArguments()
	if err != nil {
		return err
	}

	err = checkFileName()
	if err != nil {
		return err
	}

	return nil
}

func checkArguments() error {
	if !areArgumentsValid(os.Args) {
		return errors.New("Invalid amount of arguments provided to the compiler.")
	}
	return nil
}

func areArgumentsValid(arguments []string) bool {
	return len(arguments) == 2
}

func checkFileName() error {
	var fileName string = os.Args[1]
	if len(fileName) <= 3 {
		return errors.New("Invalid file name provided to the compiler.")
	} else if fileName[len(fileName)-3:] != ".ls" {
		return errors.New("Invalid file type provided to the compiler.")
	}

	return nil
}
