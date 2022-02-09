package main

/* ReCT-Go-Compiler
 * Heyyy new cli has been added to main! There are some new flags and stuff you can add (checkout cli.go for docs)...
 * Anyway, I suppose you're wondering what to do when compiling a project for development/testing purposes?
 * I would recommend running: go build -v -a -o "rgoc" . && ./rgoc -t
 * The command above will build the project, create an executable called "rgoc" and run the executable by testing all the test files in /tests/
 * furthermore, you can run: ./rgoc path/to/file.rct, to interpret the file.
 */

// main it looks super empty because all the handling has moved to cli lol
func main() {
	// all these functions can be found in cli.go

	// Init defines all the flags and initializes them
	Init()
	// ProcessFlags does as it says, takes the flags from Init and uses them to run parts of the compiler
	ProcessFlags()
}
