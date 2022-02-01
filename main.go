package main

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/parser"
	"ReCT-Go-Compiler/print"
	"fmt"
)

func main() {
	print.PrintC(print.Yellow, "Testing lexer")
	tokens := lexer.Lex("tests/test1.rct")
	for _, token := range tokens {
		fmt.Println(token.String(false))
	}

	print.PrintC(print.Yellow, "Testing parser")
	members := parser.Parse(tokens)
	fmt.Println(len(members))

	for _, member := range members {
		// if the statement is a global one -> get the statement inside
		if member.NodeType() == 0 {
			fmt.Println(member.(*nodes.GlobalStatementMember).Statement.NodeType())
		} else {
			fmt.Println(member.NodeType())
		}
	}
}
