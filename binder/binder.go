package binder

import (
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/symbols"
)

type Binder struct {
	MemberScope    Scope
	FunctionSymbol symbols.FunctionSymbol

	LabelCounter   int
	BreakLabels    []boundnodes.BoundLabel
	ContinueLabels []boundnodes.BoundLabel
}

// helpers for the label stacks
func (bin *Binder) PushLabels(breakLabel boundnodes.BoundLabel, continueLabel boundnodes.BoundLabel) {
	bin.BreakLabels = append(bin.BreakLabels, breakLabel)
	bin.ContinueLabels = append(bin.ContinueLabels, continueLabel)
}

func (bin *Binder) PopLabels() {
	bin.BreakLabels[len(bin.BreakLabels)-1] = ""
	bin.ContinueLabels[len(bin.ContinueLabels)-1] = ""
	bin.BreakLabels = bin.BreakLabels[:len(bin.BreakLabels)-1]
	bin.ContinueLabels = bin.ContinueLabels[:len(bin.ContinueLabels)-1]
}

func (bin *Binder) GetLabels() (boundnodes.BoundLabel, boundnodes.BoundLabel) {
	return bin.BreakLabels[len(bin.BreakLabels)-1],
		bin.ContinueLabels[len(bin.ContinueLabels)-1]
}

// binder action
