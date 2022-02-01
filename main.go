package main

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/parser"
	"ReCT-Go-Compiler/print"
	"fmt"
)

func main() {
	print.PrintC(print.Green, "Testing lexer")
	print.PrintC(print.Green, "-------------\n")
	tokens := lexer.Lex("tests/test0.8.rct")
	for _, token := range tokens {
		fmt.Println(token.String(false))
	}
	fmt.Print("\n")

	print.PrintC(print.Yellow, "Testing parser")
	print.PrintC(print.Yellow, "--------------\n")
	members := parser.Parse(tokens)
	fmt.Printf("Members: %d\n", len(members))

	for _, member := range members {
		member.Print("")
	}
}
