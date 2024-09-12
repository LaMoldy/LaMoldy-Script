package compiler

import (
	"os"
	"testing"

	"github.com/lamoldy/lamoldy-script/pkg/compiler"
)

func TestMain(m *testing.M) {
	// Setup code
	originalArgs := os.Args
	// Run the tests
	exitCode := m.Run()
	// Teardown code
	os.Args = originalArgs
	// Exit
	os.Exit(exitCode)
}

func TestInvalidArgumentCount(t *testing.T) {
	os.Args = []string{""}

	err := compiler.Run()
	if err == nil {
		t.Error("Got nil expected err")
	}
}

func TestFileNameWithSingleCharacter(t *testing.T) {
	os.Args = []string{"", "a"}

	err := compiler.Run()
	if err == nil {
		t.Error("Got nil expected err")
	}
}

func TestFileNameWithTwoCharacters(t *testing.T) {
	os.Args = []string{"", "ab"}

	err := compiler.Run()
	if err == nil {
		t.Error("Got nil expected err")
	}
}

func TestFileNameWithThreeCharacters(t *testing.T) {
	os.Args = []string{"", "abc"}

	err := compiler.Run()
	if err == nil {
		t.Error("Got nil expected err")
	}
}

func TestValidFileName(t *testing.T) {
	os.Args = []string{"", "../../example.ls"}

	err := compiler.Run()
	if err != nil {
		t.Error("Got err expected nil")
	}
}

func TestInvalidFileExtension(t *testing.T) {
	os.Args = []string{"", "example.lc"}

	err := compiler.Run()
	if err == nil {
		t.Error("Got nil expected err")
	}
}
