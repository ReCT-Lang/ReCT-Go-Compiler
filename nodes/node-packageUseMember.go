package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
)

// basic global statement member
type PackageUseMember struct {
	MemberNode

	PackageKeyword lexer.Token
	Package        lexer.Token
}

// implement node type from interface
func (PackageUseMember) NodeType() NodeType { return PackageUse }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node PackageUseMember) Span() print.TextSpan {
	return node.PackageKeyword.Span.SpanBetween(node.Package.Span)
}

// node print function
func (node PackageUseMember) Print(indent string) {
	print.PrintC(print.Cyan, indent+"- PackageUseMember: "+node.Package.Value)
}

// "constructor" / ooga booga OOP cave man brain
func CreatePackageUseMember(kw lexer.Token, pkg lexer.Token) PackageUseMember {
	return PackageUseMember{
		PackageKeyword: kw,
		Package:        pkg,
	}
}
