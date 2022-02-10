package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/symbols"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// this file is just keeping track of how ReCT types map to LLVM types

var IRTypes map[string]types.Type = map[string]types.Type{
	builtins.Void.Fingerprint():   types.Void,
	builtins.Bool.Fingerprint():   types.I1,
	builtins.Int.Fingerprint():    types.I32,
	builtins.Float.Fingerprint():  types.Float,
	builtins.String.Fingerprint(): types.I8Ptr,
	builtins.Any.Fingerprint():    types.I8Ptr,
}

type Global struct {
	IRGlobal *ir.Global
	Type     symbols.TypeSymbol
}

type Local struct {
	IRLocal *ir.InstAlloca
	IRBlock *ir.Block
	Type    symbols.TypeSymbol
}

type Function struct {
	IRFunction    *ir.Func
	BoundFunction binder.BoundFunction
}