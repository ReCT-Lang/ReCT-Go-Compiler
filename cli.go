package main

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/evaluator"
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/parser"
	"ReCT-Go-Compiler/print"
	"flag"
	"fmt"
	"os"
)

// Defaults

var helpFlag bool      // false -h
var interpretFlag bool // true  -i
var showVersion bool   // false -v
var fileLog bool       // false -l
var debug bool         // -xx
var files []string

const executableName string = "rgoc" // in case we change it later
const discordInvite string = "https://discord.gg/kk9MsnABdF"
const currentVersion string = "1.1"

// init initializes and processes (parses) compiler flags
func init() {
	flag.BoolVar(&helpFlag, "h", false, "Shows this help message")
	flag.BoolVar(&interpretFlag, "i", true, "Enables interpreter mode, source code will be interpreted instead of compiled.")
	flag.BoolVar(&showVersion, "v", false, "Shows current ReCT version the compiler supports")
	flag.BoolVar(&fileLog, "l", false, "Logs process information in a log file")
	flag.BoolVar(&debug, "xx", false, "Shows brief process information in the command line")
	files = flag.Args() // Other arguments like executable name or files
	flag.Parse()
}

// processFlags goes through each flag and decides how they have an effect on the output of the compiler
func processFlags() {
	// Show version has higher priority than help menu
	if showVersion {
		version()
		os.Exit(0)
	}
	// If they use "-h" or only enter the executable name "rgoc"
	// Show the help menu because they're obviously insane.
	if helpFlag || len(files) <= 1 {
		help()
		os.Exit(0)
	}
	if interpretFlag {
		interpretFiles()
	}
}

// interpretFiles runs everything to interpret the files, currently only supports up to one file
func interpretFiles() {
	print.WriteC(print.Green, "-> Lexing...  ")
	tokens := lexer.Lex(files[1])
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

// help shows help message (pretty standard nothing special)
func help() {
	fmt.Println("--------------\nReCT Go Compiler\n--------------")
	fmt.Print("Usage: ")
	print.PrintC(print.Green, "rgoc <file> [options]\n")
	fmt.Println("<file> can be the path to any ReCT file (.rct)\n")
	fmt.Println("[Options]")
	fmt.Printf("Help : %s -h : disabled (default) : Shows this help message!\n", executableName)
	fmt.Printf("Interpret : %s -i : enabled (default) : Enables interpreter mode, source code will be interpreted instead of compiled.\n", executableName)
	fmt.Printf("File logging : %s -l : disabled (default) : Logs process information in a log file\n", executableName)
	fmt.Printf("Debug : %s -xx : disabled (default) : Shows brief process information in the command line\n\n", executableName)
	fmt.Printf("Still having troubles? Get help on the offical Discord server: %s!\n", discordInvite)
}

// version Shows the current compiler version
func version() {
	fmt.Println("ReCT Go Compiler")
	fmt.Print("ReCT version: ")
	print.PrintC(print.Blue, currentVersion)
	fmt.Printf("\nFor more informatin, why not join the discord? %s\n\n", discordInvite)
}
