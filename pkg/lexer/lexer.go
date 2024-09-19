package lexer

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/lamoldy/lamoldy-script/pkg/token"
)

type Lexer struct {
	LineNumber     int
	CharacterCount int
	Position       int
	Tokens         []token.Token
	Source         []byte
}

func Create(source []byte) Lexer {
	return Lexer{
		LineNumber:     1,
		CharacterCount: 0,
		Position:       0,
		Tokens:         nil,
		Source:         source,
	}
}

func isWhitespace(character rune) bool {
	return unicode.IsSpace(character)
}

func isLetter(character rune) bool {
	return unicode.IsLetter(character)
}

func isDigit(character rune) bool {
	return unicode.IsDigit(character)
}

func isSymbol(character rune) bool {
	symbols := []rune{'{', '}', '(', ')', '=', '+', ';', ':', '"', '\'', ','}

	for symbol := range symbols {
		if character == symbols[symbol] {
			return true
		}
	}
	return false
}

func skipWhitespace(l *Lexer) {
	for true {
		if l.Position == len(l.Source)-1 {
			break
		}
		if !isWhitespace(rune(l.Source[l.Position])) {
			break
		}
		if rune(l.Source[l.CharacterCount]) == '\n' {
			l.LineNumber++
			l.CharacterCount = 0
		}
		l.CharacterCount++
		l.Position++
	}
}

func getDigits(l *Lexer) {
	value := ""
	for true {
		if !isDigit(rune(l.Source[l.Position])) && l.Position < len(l.Source) {
			t := token.Create(value, token.Number)
			l.Tokens = append(l.Tokens, t)
			break
		}
		value += string(l.Source[l.Position])
		l.CharacterCount++
		l.Position++
	}
}

func getIdentifier(l *Lexer) {
	value := ""
	for true {
		if !isLetter(rune(l.Source[l.Position])) && rune(l.Source[l.Position]) != '_' && l.Position < len(l.Source) {
			t := token.Create(value, token.Identifier)
			l.Tokens = append(l.Tokens, t)
			break
		}
		value += string(l.Source[l.Position])
		l.CharacterCount++
		l.Position++
	}
}

func getSpecialCharacters(l *Lexer) {
	character := l.Source[l.Position]
	value := string(character)
	tokenType := token.Illegal

	switch character {
	case '{':
		tokenType = token.LeftCurly
	case '}':
		tokenType = token.RightCurly
	case '(':
		tokenType = token.LeftParen
	case ')':
		tokenType = token.RightParen
	case '=':
		tokenType = token.Assign
	case '+':
		tokenType = token.Plus
	case ';':
		tokenType = token.Semicolon
	case ':':
		tokenType = token.Colon
	case ',':
		tokenType = token.Comma
	case '"':
		tokenType = token.DoubleQuotes
	case '\'':
		tokenType = token.SingleQuotes
	}

	l.Position++
	l.CharacterCount++

	t := token.Create(value, tokenType)
	l.Tokens = append(l.Tokens, t)
}

func (l *Lexer) LexProgram() error {
	for true {
		if l.Position == len(l.Source)-1 {
			break
		}
		character := rune(l.Source[l.Position])
		if isWhitespace(character) {
			skipWhitespace(l)
		} else if isDigit(character) {
			getDigits(l)
		} else if isLetter(character) || character == '_' {
			getIdentifier(l)
		} else if isSymbol(character) {
			getSpecialCharacters(l)
		} else {
			message := fmt.Sprintf(
				"Invalid character on line: %d character: %d, VALUE: %v",
				l.LineNumber,
				l.CharacterCount,
				string(l.Source[l.Position]),
			)
			return errors.New(message)
		}
	}
	return nil
}
