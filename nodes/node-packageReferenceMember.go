package nodes

import (
	"github.com/ReCT-Lang/ReCT-Go-Compiler/lexer"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
)

// basic global statement member
type PackageReferenceMember struct {
	MemberNode

	PackageKeyword lexer.Token
	Package        lexer.Token
}

// implement node type from interface
func (PackageReferenceMember) NodeType() NodeType { return PackageReference }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node PackageReferenceMember) Span() print.TextSpan {
	return node.PackageKeyword.Span.SpanBetween(node.Package.Span)
}

// node print function
func (node PackageReferenceMember) Print(indent string) {
	print.PrintC(print.Cyan, indent+"- PackageReferenceMember: "+node.Package.Value)
}

// "constructor" / ooga booga OOP cave man brain
func CreatePackageReferenceMember(kw lexer.Token, pkg lexer.Token) PackageReferenceMember {
	return PackageReferenceMember{
		PackageKeyword: kw,
		Package:        pkg,
	}
}
