package main

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/parser"
	"ReCT-Go-Compiler/print"
	"fmt"
)

func main() {
	//var i = cli.New()
	//i.Execute([]string{"-s", "tests/test1.rct", "--experimental", "-q"})
	//return

	print.PrintC(print.Green, "Testing lexer")
	print.PrintC(print.Green, "-------------\n")
	tokens := lexer.Lex("tests/test1.rct")
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
