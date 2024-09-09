package main

import (
	"fmt"
	"os"

	"github.com/lamoldy/lamoldy-script/pkg/compiler"
)

func main() {
	err := compiler.Run()
	if err != nil {
		fmt.Printf("Error: %s\nUsage: lsc <source_file>.ls", err)
		os.Exit(1)
	}
}
