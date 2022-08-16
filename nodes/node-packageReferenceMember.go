package nodes

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/print"
)

// basic global statement member
type PackageReferenceMember struct {
	MemberNode
	Package lexer.Token
}

// implement node type from interface
func (PackageReferenceMember) NodeType() NodeType { return PackageReference }

// Position returns the starting line and column, and the total length of the statement
// The starting line and column aren't always the absolute beginning of the statement just what's most
// convenient.
func (node PackageReferenceMember) Position() (int, int, int) {
	return 0, 0, 0
}

// node print function
func (node PackageReferenceMember) Print(indent string) {
	print.PrintC(print.Cyan, indent+"- PackageReferenceMember: "+node.Package.Value)
}

// "constructor" / ooga booga OOP cave man brain
func CreatePackageReferenceMember(pkg lexer.Token) PackageReferenceMember {
	return PackageReferenceMember{
		Package: pkg,
	}
}
