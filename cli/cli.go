package cli

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/parser"
	"fmt"

	"github.com/fatih/color"
	"github.com/hashicorp/go-version"
	"github.com/jessevdk/go-flags"
)

type CLInterface struct {
	ver *version.Version
}

func New() *CLInterface {
	var cl = new(CLInterface)
	var err error
	cl.ver, err = version.NewVersion("0.1.0")

	if err != nil {
		color.Red("Invalid version" + cl.ver.String())
	}
	return cl
}

func (cl *CLInterface) Execute(args []string) {
	color.Green("ReCT CLI - v" + cl.ver.String())
	if len(args) < 1 {
		color.Yellow("No arguments specified!")
		return
	}

	var opts struct {
		Quiet        bool   `short:"q" long:"quiet" description:"Disable verbose actions"`
		File         string `short:"s" long:"source" alias:"f" description:"The source to compile" required:"true"`
		Unsafe       bool   `long:"unsafe" description:"Enable unsafe code"`
		Experimental bool   `long:"experimental" description:"Enable experimental features"`
	}

	color.Set(color.FgRed)

	var err error
	args, err = flags.ParseArgs(&opts, args)
	if err != nil {
		return
	}

	color.Set(color.Reset)
	color.Cyan("Compiling file " + opts.File)

	if !opts.Quiet {
		if opts.Experimental {
			color.Cyan("Experimental features are enabled!")
		}
		if opts.Unsafe {
			color.Cyan("Unsafe code is enabled!")
		}
	}

	tokens := lexer.Lex(opts.File)
	if !opts.Quiet {
		for _, token := range tokens {
			fmt.Println(token.String(false))
		}
	}

	members := parser.Parse(tokens)

	if !opts.Quiet {
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
}
