package compiler

import (
	"errors"
	"fmt"
	"os"

	"github.com/lamoldy/lamoldy-script/pkg/file"
	"github.com/lamoldy/lamoldy-script/pkg/lexer"
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

	fileData, err := file.GetFileContent()
	if err != nil {
		return err
	}

	fmt.Println(string(fileData))

	lexer := lexer.Create(fileData)
	err = lexer.LexProgram()
	if err != nil {
		return err
	}

	for _, t := range lexer.Tokens {
		fmt.Println(t.String())
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
