package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/symbols"

	"github.com/llir/llvm/ir"
)

type Emitter struct {
	Program         binder.BoundProgram
	Module          *ir.Module
	UseFingerprints bool
}

func Emit(program binder.BoundProgram, useFingerprints bool) *ir.Module {
	emitter := Emitter{
		Program:         program,
		Module:          ir.NewModule(),
		UseFingerprints: useFingerprints,
	}

	for _, fnc := range emitter.Program.Functions {
		if !fnc.Symbol.BuiltIn {
			emitter.EmitFunction(fnc.Symbol, fnc.Body)
		}
	}

	return emitter.Module
}

func (emt *Emitter) EmitFunction(sym symbols.FunctionSymbol, body boundnodes.BoundBlockStatementNode) {
	// figure out all parameters and their types
	params := make([]*ir.Param, 0)
	for _, param := range sym.Parameters {
		// figure out how to call this parameter
		paramName := tern(emt.UseFingerprints, param.Fingerprint(), param.Name)

		// create it
		params = append(params, ir.NewParam(paramName, IRTypes[param.Type.Fingerprint()]))
	}

	// figure out the return type
	returnType := IRTypes[sym.Type.Fingerprint()]

	// the function name
	functionName := tern(emt.UseFingerprints, sym.Fingerprint(), sym.Name)

	// create an IR function definition
	function := emt.Module.NewFunc(functionName, returnType, params...)
	block := function.NewBlock("")

	block.NewRet(nil)
}

// yes
func tern(cond bool, str1 string, str2 string) string {
	if cond {
		return str1
	} else {
		return str2
	}
}
