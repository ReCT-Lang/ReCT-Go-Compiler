package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/symbols"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// this file is just keeping track of how ReCT types map to LLVM types

func (emt *Emitter) IRTypes(fingerprint string) types.Type {
	switch fingerprint {
	case builtins.Void.Fingerprint():
		return types.Void
	case builtins.Bool.Fingerprint():
		return types.I1
	case builtins.Int.Fingerprint():
		return types.I32
	case builtins.Float.Fingerprint():
		return types.Float
	case builtins.String.Fingerprint():
		return types.NewPointer(emt.Classes[emt.Id(builtins.String)].Type)
	case builtins.Any.Fingerprint():
		return types.NewPointer(emt.Classes[emt.Id(builtins.Any)].Type)
	}

	return nil
}

type Global struct {
	IRGlobal *ir.Global
	Type     symbols.TypeSymbol
}

type Local struct {
	IRLocal value.Value
	IRBlock *ir.Block
	Type    symbols.TypeSymbol
	IsSet   bool
}

type Function struct {
	IRFunction    *ir.Func
	BoundFunction binder.BoundFunction
}

type Class struct {
	Type        types.Type
	Constructor *ir.Func
}
