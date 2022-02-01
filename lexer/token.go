package lexer

import "fmt"

// TokenKind basically an enum containing all token types.
// TokenKind has been changed from int to string for better debugging. 
type TokenKind string 
const (
	// Keywords
	VarKeyword   TokenKind = "var (Keyword)"
	SetKeyword             = "set (Keyword)"
	ToKeyword              = "to (Keyword)"
	IfKeyword              = "if (Keyword)"
	ElseKeyword            = "else (Keyword)"
	TrueKeyword            = "true (Keyword)"
	FalseKeyword           = "false (Keyword)"
	PrintKeyword           = "Print (Keyword)"
	FunctionKeyword		   = "function (Keyword)"
	FromKeyword 		   = "from (Keyword)"
	ForKeyword 			   = "for (Keyword)"
	ReturnKeyword = "Return (Keyword)"

	// Tokens
	EOF                   = "EndOfFile"
	IdToken               = "Identifier"
	StringToken           = "String"
	NumberToken           = "Number"
	PlusToken             = "Plus(+)"
	MinusToken            = "Minus(-)"
	StarToken             = "Star(*)"
	SlashToken            = "Slash(/)"
	EqualsToken           = "Equals(=)"
	OpenBraceToken        = "OpenBrace"
	CloseBraceToken       = "Closebrace"
	OpenParenthesisToken  = "OpenParenthesis"
	CloseParenthesisToken = "CloseParenthesis"
	AssignToken           = "<- (AssignToken)"
	GreaterThanToken      = "> (GreaterThanToken)"
	LessThanToken         = "< (LessthanToken" 

	BadToken = "Token Error (BadToken)" // Naughty ;)

	Semicolon = "Semicolon ;" // Used to separate statements (for now... )
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
	if !pretty {
		return fmt.Sprintf("Token { value: %s, kind: %d, position: (%d, %d)}", t.Value, t.Kind, t.Line, t.Column)
	} else {
		return fmt.Sprintf("Token { \n\tvalue: %s, \n\tkind: %d, \n\tposition: (%d, %d)\n}", t.Value, t.Kind, t.Line, t.Column)
	}
}
