package main

import (
	"ReCT-Go-Compiler/lexer"
	"fmt"
)

func main() {
	fmt.Println("Testing lexer")
	tokens := lexer.Lex("tests/test1.rct")
	for _, token := range tokens {
		fmt.Println(token.String())
	}
}
