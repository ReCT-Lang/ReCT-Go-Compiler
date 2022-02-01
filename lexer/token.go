package lexer

import "fmt"

// TokenKind basically an enum containing all token types
type TokenKind int // Change these to strings for better debugging
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

	Semicolon = 24 // Used to separate statements (for now... )
)

// Token stores information about lexical structures in the text
type Token struct {
	Value  string
	Kind   TokenKind
	Line   int 
	Column int 
}

// CreateToken returns a Token created from the arguments provided
func CreateToken(value string, kind TokenKind, line int, column int) Token {
	return Token{
		value, kind, line, column,
	}
}

// String easy representation of a Token
func (t Token) String(pretty bool) string {
	if !pretty) {
		return fmt.Sprintf("Token { value: %s, kind: %d, position: (%d, %d)}", t.Value, t.Kind, t.Line, t.Column)
	} else {
		return fmt.Sprintf("Token { \n\tvalue: %s, \n\tkind: %d, \n\tposition: (%d, %d)\n}", t.Value, t.Kind, t.Line, t.Column)
	}
}
