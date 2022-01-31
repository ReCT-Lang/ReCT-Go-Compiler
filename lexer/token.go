package lexer

import "fmt"

// TokenKind : Basically an enum
type TokenKind int

const (
	// Keywords
	VarKeyword   TokenKind = 0
	SetKeyword             = 1
	ToKeyword              = 2
	IfKeyword              = 3
	ElseKeyword            = 4
	TrueKeyword            = 5
	FalseKeyword           = 6
	PrintKeyword           = 7

	// Tokens
	EOF                   = 8
	IdToken               = 9
	StringToken           = 10
	NumberToken           = 11
	PlusToken             = 12
	MinusToken            = 13
	StarToken             = 14
	SlashToken            = 15
	EqualsToken           = 16
	OpenBraceToken        = 17
	CloseBraceToken       = 18
	OpenParenthesisToken  = 19
	CloseParenthesisToken = 20
	AssignToken           = 21
	GreaterThanToken      = 22
	LessThanToken         = 23

	BadToken = -1 // Naughty ;)

	Semicolon = 24 // Used to separate statements
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
