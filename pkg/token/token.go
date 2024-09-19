package token

type Token struct {
	Type  TokenType
	Value string
}

func Create(value string, tokenType TokenType) Token {
	return Token{tokenType, value}
}

func (t *Token) String() string {
	tokenString := "Token Value: " + t.Value + " Token Type: " + t.Type.String()
	return tokenString
}
