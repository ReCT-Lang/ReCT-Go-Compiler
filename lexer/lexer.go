package lexer

import (
	"errors"
	"fmt"
	"os"
)

// lexer : internal struct for Lexical Analysis
type lexer struct {
	Code   []byte
	Line   int
	Column int
	Index  int
}

func Lex(filename string) {
	scanner := lexer{handleFileOpen(filename), 0, 0, 0}

	for scanner.Index < len(scanner.Code) {

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
