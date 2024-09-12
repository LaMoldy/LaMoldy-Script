package file

import (
	"os"
	"strings"
	"testing"

	"github.com/lamoldy/lamoldy-script/pkg/file"
)

var exampleText string = `function :: is_equal(x: int, y: int) :: bool {
    return x == y;
}

function :: main() :: int {
    let a: int = 10;
    let b: int = 11;

    let result = is_equal(a,b);

    println("Hello, World");

    return 0;
}
`

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

func TestGetFileContent(t *testing.T) {
	os.Args = []string{"", "../../example.ls"}

	data, _ := file.GetFileContent()
	normalizeData := strings.ReplaceAll(string(data), "\r\n", "\n")

	if strings.Compare(normalizeData, exampleText) != 0 {
		t.Errorf("got:\n%s\nExpected:\n%s\n", normalizeData, exampleText)
	}
}

func TestGetFileContentErr(t *testing.T) {
	os.Args = []string{"", "../exampl.ls"}

	_, err := file.GetFileContent()
	if err == nil {
		t.Error("got nil, expected err")
	}
}
