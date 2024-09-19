package lexer

import (
	"reflect"
	"testing"

	"github.com/lamoldy/lamoldy-script/pkg/lexer"
	"github.com/lamoldy/lamoldy-script/pkg/token"
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

var expectedTokens = []token.Token{
	{Value: "function", Type: token.Identifier},
	{Value: ":", Type: token.Colon},
	{Value: ":", Type: token.Colon},
	{Value: "is_equal", Type: token.Identifier},
	{Value: "(", Type: token.LeftParen},
	{Value: "x", Type: token.Identifier},
	{Value: ":", Type: token.Colon},
	{Value: "int", Type: token.Identifier},
	{Value: ",", Type: token.Comma},
	{Value: "y", Type: token.Identifier},
	{Value: ":", Type: token.Colon},
	{Value: "int", Type: token.Identifier},
	{Value: ")", Type: token.RightParen},
	{Value: ":", Type: token.Colon},
	{Value: ":", Type: token.Colon},
	{Value: "bool", Type: token.Identifier},
	{Value: "{", Type: token.LeftCurly},
	{Value: "return", Type: token.Identifier},
	{Value: "x", Type: token.Identifier},
	{Value: "=", Type: token.Assign},
	{Value: "=", Type: token.Assign},
	{Value: "y", Type: token.Identifier},
	{Value: ";", Type: token.Semicolon},
	{Value: "}", Type: token.RightCurly},
	{Value: "function", Type: token.Identifier},
	{Value: ":", Type: token.Colon},
	{Value: ":", Type: token.Colon},
	{Value: "main", Type: token.Identifier},
	{Value: "(", Type: token.LeftParen},
	{Value: ")", Type: token.RightParen},
	{Value: ":", Type: token.Colon},
	{Value: ":", Type: token.Colon},
	{Value: "int", Type: token.Identifier},
	{Value: "{", Type: token.LeftCurly},
	{Value: "let", Type: token.Identifier},
	{Value: "a", Type: token.Identifier},
	{Value: ":", Type: token.Colon},
	{Value: "int", Type: token.Identifier},
	{Value: "=", Type: token.Assign},
	{Value: "10", Type: token.Number},
	{Value: ";", Type: token.Semicolon},
	{Value: "let", Type: token.Identifier},
	{Value: "b", Type: token.Identifier},
	{Value: ":", Type: token.Colon},
	{Value: "int", Type: token.Identifier},
	{Value: "=", Type: token.Assign},
	{Value: "11", Type: token.Number},
	{Value: ";", Type: token.Semicolon},
	{Value: "let", Type: token.Identifier},
	{Value: "result", Type: token.Identifier},
	{Value: "=", Type: token.Assign},
	{Value: "is_equal", Type: token.Identifier},
	{Value: "(", Type: token.LeftParen},
	{Value: "a", Type: token.Identifier},
	{Value: ",", Type: token.Comma},
	{Value: "b", Type: token.Identifier},
	{Value: ")", Type: token.RightParen},
	{Value: ";", Type: token.Semicolon},
	{Value: "println", Type: token.Identifier},
	{Value: "(", Type: token.LeftParen},
	{Value: "\"", Type: token.DoubleQuotes},
	{Value: "Hello", Type: token.Identifier},
	{Value: ",", Type: token.Comma},
	{Value: "World", Type: token.Identifier},
	{Value: "\"", Type: token.DoubleQuotes},
	{Value: ")", Type: token.RightParen},
	{Value: ";", Type: token.Semicolon},
	{Value: "return", Type: token.Identifier},
	{Value: "0", Type: token.Number},
	{Value: ";", Type: token.Semicolon},
	{Value: "}", Type: token.RightCurly},
}

func TestCreateLexer(t *testing.T) {
	l := lexer.Create(nil)

	expected := lexer.Lexer{
		LineNumber:     1,
		CharacterCount: 0,
		Position:       0,
		Tokens:         nil,
		Source:         nil,
	}

	if !reflect.DeepEqual(l, expected) {
		t.Error("Lexer create failed to intialize with proper values")
	}
}

func TestLexer(t *testing.T) {
	l := lexer.Create([]byte(exampleText))

	err := l.LexProgram()
	if err != nil {
		t.Errorf("There was an error when lexing the program, expected non\nError: %v", err)
	}

	if !reflect.DeepEqual(l.Tokens, expectedTokens) {
		t.Errorf("The lexer tokens don't equal the expected tokens\n%v\n%v", l.Tokens, expectedTokens)
	}
}
