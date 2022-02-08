package main

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/evaluator"
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/parser"
	"ReCT-Go-Compiler/print"
)

func main() {
	//var i = cli.New()
	//i.Execute([]string{"-s", "tests/test1.rct", "--experimental", "-q"})
	//return

	print.WriteC(print.Green, "-> Lexing...  ")
	tokens := lexer.Lex("tests/BuiltInFunctions.rct")
	print.PrintC(print.Green, "Done!")

	print.WriteC(print.Yellow, "-> Parsing... ")
	members := parser.Parse(tokens)
	print.PrintC(print.Green, "Done!")

	print.WriteC(print.Red, "-> Binding... ")
	boundProgram := binder.BindProgram(members)
	print.PrintC(print.Green, "Done!")
	//boundProgram.Print()

	print.PrintC(print.Cyan, "-> Evaluating!")
	evaluator.Evaluate(boundProgram)

}
