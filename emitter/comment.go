package emitter

// implementation for a comment
// stolen from github
// https://github.com/llir/llvm/issues/99

import (
	"strings"

	"github.com/llir/llvm/ir"
)

func NewComment(s string) *Comment {
	return &Comment{
		Text: s,
	}
}

type Comment struct {
	Text string
	// embed a dummy ir.Instruction to have Comment implement the ir.Instruction
	// interface.
	ir.Instruction
}

func (c *Comment) LLString() string {
	return "; " + strings.Replace(c.Text, "\n", "; ", -1)
}
