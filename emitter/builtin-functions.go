package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/irtools"
	"ReCT-Go-Compiler/symbols"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func (emt *Emitter) EmitBuiltInFunctions() {

	// link references to the C standard libs
	emt.EmitCLibReferences()

	// read the system lib module
	module := irtools.ReadModule("./systemlib/systemlib_lin.ll")

	// link our classes and arc
	emt.EmitClassAndArcReferences(module)

	// link our built-ins
	emt.EmitSystemFuncReferences(module)

	// Version()
	//Version := emt.Module.NewFunc("rct_Version", types.I8Ptr)
	//body := Version.NewBlock("")
	//body.NewRet(emt.CopyStringNoFree(body, emt.GetStringConstant(body, "1.1")))

	//emt.Functions[tern(emt.UseFingerprints, builtins.Version.Fingerprint(), builtins.Version.Name)] = Function{IRFunction: Version, BoundFunction: binder.BoundFunction{Symbol: builtins.Version}}

}

func (emt *Emitter) EmitCLibReferences() {
	printf := emt.Module.NewFunc("printf", types.I32, ir.NewParam("format", types.I8Ptr))
	printf.Sig.Variadic = true
	emt.CFuncs["printf"] = printf

	scanf := emt.Module.NewFunc("scanf", types.I32, ir.NewParam("format", types.I8Ptr), ir.NewParam("dest", types.I8Ptr))
	scanf.Sig.Variadic = true
	emt.CFuncs["scanf"] = scanf

	strcpy := emt.Module.NewFunc("strcpy", types.Void, ir.NewParam("dest", types.I8Ptr), ir.NewParam("src", types.I8Ptr))
	emt.CFuncs["strcpy"] = strcpy

	strcat := emt.Module.NewFunc("strcat", types.Void, ir.NewParam("dest", types.I8Ptr), ir.NewParam("src", types.I8Ptr))
	emt.CFuncs["strcat"] = strcat

	strlen := emt.Module.NewFunc("strlen", types.I32, ir.NewParam("str", types.I8Ptr))
	emt.CFuncs["strlen"] = strlen

	strcmp := emt.Module.NewFunc("strcmp", types.I32, ir.NewParam("left", types.I8Ptr), ir.NewParam("right", types.I8Ptr))
	emt.CFuncs["strcmp"] = strcmp

	malloc := emt.Module.NewFunc("malloc", types.I8Ptr, ir.NewParam("len", types.I32))
	emt.CFuncs["malloc"] = malloc

	free := emt.Module.NewFunc("free", types.Void, ir.NewParam("dest", types.I8Ptr))
	emt.CFuncs["free"] = free

	snprintf := emt.Module.NewFunc("snprintf", types.I32, ir.NewParam("dest", types.I8Ptr), ir.NewParam("len", types.I32), ir.NewParam("format", types.I8Ptr))
	snprintf.Sig.Variadic = true
	emt.CFuncs["snprintf"] = snprintf

	atoi := emt.Module.NewFunc("atoi", types.I32, ir.NewParam("str", types.I8Ptr))
	emt.CFuncs["atoi"] = atoi

	atof := emt.Module.NewFunc("atof", types.Double, ir.NewParam("str", types.I8Ptr))
	emt.CFuncs["atof"] = atof
}

func (emt *Emitter) EmitSystemFuncReferences(module *ir.Module) {

	rctFuncs := irtools.FindFunctionsWithPrefix(module, "rct_")

	for _, fnc := range rctFuncs {
		functionName := strings.Split(fnc.Name(), "_")[1]
		functionSymbol := symbols.FunctionSymbol{}

		// find what system function symbol this links to
		for _, sysFnc := range builtins.Functions {
			if functionName == sysFnc.Name {
				functionSymbol = sysFnc
				break
			}
		}

		importedFunction := emt.ImportFunction(fnc)
		emt.Functions[emt.Id(functionSymbol)] = Function{IRFunction: importedFunction, BoundFunction: binder.BoundFunction{Symbol: functionSymbol}}
	}
}

func (emt *Emitter) EmitClassAndArcReferences(module *ir.Module) {
	// load module
	emt.LoadAndReferenceClasses(module)

	// reference arc functions
	arcFuncs := irtools.FindFunctionsWithPrefix(module, "arc_")

	for _, fnc := range arcFuncs {
		emt.ArcFuncs[strings.Split(fnc.Name(), "_")[1]] = fnc
		emt.ImportFunction(fnc)
	}

	// reference exc functions
	excFuncs := irtools.FindFunctionsWithPrefix(module, "exc_")

	for _, fnc := range excFuncs {
		emt.ExcFuncs[strings.Split(fnc.Name(), "_")[1]] = fnc
		emt.ImportFunction(fnc)
	}
}
