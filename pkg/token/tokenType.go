package token

type TokenType int

const (
	Illegal TokenType = iota + 1
	Identifier
	Number
	Assign
	Plus
	Comma
	Semicolon
	Colon
	DoubleQuotes
	SingleQuotes
	LeftParen
	RightParen
	LeftCurly
	RightCurly
)

func (t TokenType) String() string {
	switch t {
	case Illegal:
		return "Illegal"
	case Identifier:
		return "Identifier"
	case Number:
		return "Number"
	case Assign:
		return "Assign"
	case Plus:
		return "Plus"
	case Comma:
		return "Comma"
	case Semicolon:
		return "Semi Colon"
	case Colon:
		return "Colon"
	case DoubleQuotes:
		return "Double Quotes"
	case SingleQuotes:
		return "Single Quotes"
	case LeftParen:
		return "Left Parenthesis"
	case RightParen:
		return "Right Parenthesis"
	case LeftCurly:
		return "Left Curly Bracket"
	case RightCurly:
		return "Right Curly Bracket"
	default:
		return "Illegal"
	}
}
