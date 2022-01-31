package lexer

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

// lexer : internal struct for Lexical Analysis
type lexer struct {
	Code   []byte
	Line   int // Not implemented yet
	Column int // Not implemented yet
	Index  int
	Tokens []Token
}

func Lex(filename string) []Token {
	scanner := lexer{
		handleFileOpen(filename),
		0,
		0,
		0,
		make([]Token, 0),
	}

	for scanner.Index < len(scanner.Code) {
		c := scanner.Code[scanner.Index]
		if unicode.IsLetter(rune(c)) {
			scanner.getId()
		} else if unicode.IsNumber(rune(c)) {
			scanner.getNumber()
		} else if c == '"' {
			scanner.getString()
		} else {
			scanner.getOperator()
			scanner.Index++
		}
	}
	return scanner.Tokens
}

func (lxr lexer) getNumber() {
	var buffer string
	buffer = string(lxr.Code[lxr.Index])
	lxr.Index++

	for lxr.Index < len(lxr.Code) && unicode.IsDigit(rune(lxr.Code[lxr.Index])) {
		buffer += string(lxr.Code[lxr.Index])
		lxr.Index++
	}

	lxr.Tokens = append(lxr.Tokens, Token{buffer, NumberToken, lxr.Line, lxr.Column})
}

func (lxr lexer) getString() {
	var buffer string
	lxr.Index++

	for lxr.Index < len(lxr.Code) && lxr.Code[lxr.Index] != '"' {
		buffer += string(lxr.Code[lxr.Index])
		lxr.Index++
	}

	lxr.Tokens = append(lxr.Tokens, Token{buffer, StringToken, lxr.Line, lxr.Column})
}

func (lxr lexer) getId() {
	var buffer string
	buffer = string(lxr.Code[lxr.Index])
	lxr.Index++

	IsLetterOrDigitOrWhatever := func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsDigit(c) || string(c) == "_"
	}
	for lxr.Index < len(lxr.Code) && IsLetterOrDigitOrWhatever(rune(lxr.Code[lxr.Index])) {
		buffer += string(lxr.Code[lxr.Index])
		lxr.Index++
	}

	lxr.Tokens = append(lxr.Tokens, Token{buffer, IdToken, lxr.Line, lxr.Column})
}

func (lxr lexer) getOperator() {
	var _token TokenKind = -1
	peek := func(offset int) byte {
		fmt.Println(lxr.Tokens[0].String())
		fmt.Printf("%d", lxr.Index)
		if lxr.Index+offset < len(lxr.Code) {
			return lxr.Code[lxr.Index+offset]
		}
		return '\000'
	}
	switch lxr.Code[lxr.Index] {
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
			lxr.Index++
			_token = AssignToken
		} else {
			_token = LessThanToken
		}
	case '>':
		_token = GreaterThanToken
	default:
		fmt.Println("ERROR: Unexpected character somewhere idk where though")
		_token = BadToken
	}
	if _token == AssignToken {
		lxr.Tokens = append(lxr.Tokens, CreateToken(
			string(lxr.Code[lxr.Index]+peek(-1)),
			_token,
			0,
			0,
		),
		)
	} else {
		lxr.Tokens = append(lxr.Tokens, CreateToken(string(lxr.Code[lxr.Index]), _token, 0, 0))
	}
}

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
