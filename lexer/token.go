package lexer

import "fmt"

// TokenKind : Basically an enum
type TokenKind int

const (
	// Keywords
	VarKeyword TokenKind = 0
	SetKeyword
	ToKeyword
	IfKeyword
	ElseKeyword
	TrueKeyword
	FalseKeyword
	PrintKeyword

	// Tokens
	EOF
	IdToken
	StringToken
	NumberToken
	PlusToken
	MinusToken
	StarToken
	SlashToken
	EqualsToken
	OpenBraceToken
	CloseBraceToken
	OpenParenthesisToken
	CloseParenthesisToken
	AssignToken
	GreaterThanToken
	LessThanToken

	BadToken = -1 // Naughty ;)

	Semicolon // Used to separate statements
)

type Token struct {
	Value  string
	Kind   TokenKind
	Line   int // Not implemented yet (see lexer)
	Column int // Not implemented yet (see lexer)
}

func CreateToken(value string, kind TokenKind, line int, column int) Token {
	return Token{
		value, kind, line, column,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("Token { value: %s, kind: %d }", t.Value, t.Kind)
}
