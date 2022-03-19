package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/symbols"
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"os"
)

// this file is just keeping track of how ReCT types map to LLVM types

func (emt *Emitter) IRTypes(typ symbols.TypeSymbol) types.Type {
	switch typ.Fingerprint() {
	case builtins.Void.Fingerprint():
		return types.Void
	case builtins.Bool.Fingerprint():
		return types.I1
	case builtins.Byte.Fingerprint():
		return types.I8
	case builtins.Int.Fingerprint():
		return types.I32
	case builtins.Float.Fingerprint():
		return types.Float
	case builtins.String.Fingerprint():
		return types.NewPointer(emt.Classes[emt.Id(builtins.String)].Type)
	case builtins.Any.Fingerprint():
		return types.NewPointer(emt.Classes[emt.Id(builtins.Any)].Type)
	case builtins.Thread.Fingerprint():
		return types.NewPointer(emt.Classes[emt.Id(builtins.Thread)].Type)
	}
	if typ.Name == builtins.Array.Name {
		if typ.SubTypes[0].IsObject {
			return types.NewPointer(emt.Classes[emt.Id(builtins.Array)].Type)
		} else {
			return types.NewPointer(emt.Classes[emt.Id(builtins.PArray)].Type)
		}
	}

	fmt.Println(typ.Fingerprint())
	os.Exit(-1)
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
	Functions   map[string]*ir.Func
	Name        string
}
