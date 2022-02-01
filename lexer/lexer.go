package lexer

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

// TODO(tokorv): hey gamer, could you add a field to the token struct for "typed value" (you can prob come up with a better name)
//               its just gonna be the value but already converted into the correct datatype


// lexer : internal struct for Lexical Analysis
type Lexer struct {
	Code   []byte
	Line   int 
	Column int 
	Index  int
	Tokens []Token
}

// Lex 
func Lex(filename string) []Token {
	scanner := &Lexer{ handleFileOpen(filename), 1, 1, 0, make([]Token, 0) }

	for c := scanner.Code[scanner.Index]; scanner.Index < len(scanner.Code); {
		if unicode.IsLetter(rune(c)) {
			scanner.getId()
		} else if unicode.IsNumber(rune(c)) {
			scanner.getNumber()
		} else if c == '"' {
			scanner.getString()
		} else if c != ' ' || c != '\n' || c != '\t' || c != '\v' {
			scanner.getOperator()
		} else {
			scanner.Increment()
		}
	}
	scanner.Tokens = append(scanner.Tokens, CreateToken("\000", EOF, lxr.Line, lxr.Column))
	return scanner.Tokens
}

// getNumber 
func (lxr *Lexer) getNumber() {
	var buffer string
	buffer = string(lxr.Code[lxr.Index])
	lxr.Increment()

	for char := lxr.current(); lxr.Index < len(lxr.Code) && unicode.IsDigit(rune(char)) {
		buffer += string(char)
		lxr.Increment()
	}

	lxr.Tokens = append(lxr.Tokens, CreateToken(buffer, NumberToken, lxr.Line, lxr.Column))
}

// getString
func (lxr *Lexer) getString() {
	var buffer string
	lxr.Increment()

	for char := lxr.current(); lxr.Index < len(lxr.Code) && char != '"'; {
		buffer += string(char)
		lxr.Increment()
	}
	lxr.Increment()
	lxr.Tokens = append(lxr.Tokens, CreateToken(buffer, StringToken, lxr.Line, lxr.Column))
}

// Increment increases the scanner's Index, Column, and Lin (if needed).
func (lxr *Lexer) Increment() {
	lxr.Index++
	lxr.Column++
	if lxr.Index < len(lxr.Code) {
		return
	} else if lxr.Code[lxr.Index] == '\n' {
		lxr.Line++
		lxr.Column = 1
	}
}

// getId
func (lxr *Lexer) getId() {
	buffer := string(lxr.Code[lxr.Index])
	lxr.Increment()

	IsLetterOrDigitOrWhatever := func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsDigit(c) || string(c) == "_" || string(c) == "."
	}
	for char := lxr.current(); lxr.Index < len(lxr.Code) && IsLetterOrDigitOrWhatever(rune(char)); {
		buffer += string(char)
		lxr.Increment()
	}

	lxr.Tokens = append(lxr.Tokens, CreateToken(buffer, IdToken, lxr.Line, lxr.Column))
}

// getOperator 
func (lxr *Lexer) getOperator() {
	var _token TokenKind

	peek := func(offset int) byte {
		if lxr.Index+offset < len(lxr.Code) {
			return lxr.Code[lxr.Index]
		}
		return '\000'
	}

	switch lxr.current() {
	case '+':
		_token = PlusToken
	case '-':
		_token = MinusToken
	case '/':
		_token = SlashToken
	case '*':
		_token = StarToken
	case '=':
		_token = EqualsToken
	case '(':
		_token = OpenParenthesisToken
	case ')':
		_token = CloseParenthesisToken
	case '{':
		_token = OpenBraceToken
	case '}':
		_token = CloseBraceToken
	case ';':
		_token = Semicolon
	case '<':
		if peek(1) == '-' {
			lxr.Increment()
			_token = AssignToken
		} else {
			_token = LessThanToken
		}
	case '>':
		_token = GreaterThanToken
	default:
		fmt.Printf("ERROR(%d, %d): Unexpected character \"%s\"!\n", lxr.Line, lxr.Column, string(lxr.Code[lxr.Index])))
		_token = BadToken
	}
	// AssignToken is 2 characters long while every other operator is 1 character.
	// (that is why they are separated).
	if _token == AssignToken {
		lxr.Tokens = append(lxr.Tokens, 
			CreateToken(string(peek(-1))+string(lxr.Code[lxr.Index]), 
			_token, 
			lxr.Line, 
			lxr.Column,
			))
	} else {
		lxr.Tokens = append(lxr.Tokens, CreateToken(
			string(lxr.Code[lxr.Index]), 
			_token, 
			lxr.Line, 
			lxr.Column,
			))
	}
	lxr.Increment()
}

// handleFileOpen 
func handleFileOpen(filename string) []byte {
	contents, err := os.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("ERROR: file \"%s\" does not exist!\n", filename)
		os.Exit(1)
	} else if errors.Is(err, os.ErrPermission) {
		fmt.Printf("ERROR: do not have permission to open file \"%s\"!\n", filename)
		os.Exit(1)
	} else if err != nil {
		fmt.Printf("ERROR: unable to open file \"%s\" for unknown reasons!", filename)
		os.Exit(1)
	}
	return contents
}
