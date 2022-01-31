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
		} else if unicode.IsDigit(rune(c)) {
			scanner.getNumber()
		} else if c == '"' {
			scanner.getString()
		} else {

			if !scanner.GetOperator() {
				// It's whitespace
				scanner.Index++
			}
		}
	}
	return scanner.Tokens
}

func (lxr lexer) getNumber() {

}

func (lxr lexer) getString() {

}

func (lxr lexer) getId() {

}

func (lxr lexer) GetOperator() bool { // Returns true if it is an operator

	return false
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
