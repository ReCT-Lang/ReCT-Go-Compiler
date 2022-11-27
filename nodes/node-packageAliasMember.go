package nodes

import (
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

// basic global statement member
type PackageAliasMember struct {
	MemberNode

	PackageKeyword lexer.Token
	Package        lexer.Token
	Alias          lexer.Token
}

// implement node type from interface
func (PackageAliasMember) NodeType() NodeType { return PackageAlias }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node PackageAliasMember) Span() print.TextSpan {
	return node.PackageKeyword.Span.SpanBetween(node.Alias.Span)
}

// node print function
func (node PackageAliasMember) Print(indent string) {
	print.PrintC(print.Cyan, indent+"- PackageAliasMember: "+node.Package.Value+" -> "+node.Alias.Value)
}

// "constructor" / ooga booga OOP cave man brain
func CreatePackageAliasMember(kw lexer.Token, pkg lexer.Token, alias lexer.Token) PackageAliasMember {
	return PackageAliasMember{
		PackageKeyword: kw,
		Package:        pkg,
		Alias:          alias,
	}
}
