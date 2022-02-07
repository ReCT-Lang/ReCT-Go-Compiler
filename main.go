package main

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/evaluator"
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
	print.PrintC(print.Green, "-----c3a120fd-d6cb-45d2-ae8a-8da4be73bbaf--------\n")
	tokens := lexer.Lex("tests/modulus.rct")
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

	print.PrintC(print.Red, "Testing binder")
	print.PrintC(print.Red, "--------------\n")
	boundProgram := binder.BindProgram(members)
	boundProgram.Print()

	print.PrintC(print.Cyan, "\nTesting evaluator")
	print.PrintC(print.Cyan, "-----------------\n")
	evaluator.Evaluate(boundProgram)

}
