package emitter

import (
	"ReCT-Go-Compiler/builtins"

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
