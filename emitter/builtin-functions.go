package emitter

import (
	"github.com/ReCT-Lang/ReCT-Go-Compiler/irtools"
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

	// Version()
	//Version := emt.Module.NewFunc("rct_Version", types.I8Ptr)
	//body := Version.NewBlock("")
	//body.NewRet(emt.CopyStringNoFree(body, emt.GetStringConstant(body, "1.1")))

	//emt.Functions[tern(emt.UseFingerprints, builtins.Version.Fingerprint(), builtins.Version.Name)] = Function{IRFunction: Version, BoundFunction: binder.BoundFunction{Symbol: builtins.Version}}

}

func (emt *Emitter) EmitCLibReferences() {

	malloc := emt.Module.NewFunc("malloc", types.I8Ptr, ir.NewParam("len", types.I32))
	emt.CFuncs["malloc"] = malloc

	free := emt.Module.NewFunc("free", types.Void, ir.NewParam("dest", types.I8Ptr))
	emt.CFuncs["free"] = free

	printf := emt.Module.NewFunc("printf", types.I32, ir.NewParam("format", types.I8Ptr))
	printf.Sig.Variadic = true
	emt.CFuncs["printf"] = printf

	snprintf := emt.Module.NewFunc("snprintf", types.I32, ir.NewParam("dest", types.I8Ptr), ir.NewParam("len", types.I32), ir.NewParam("format", types.I8Ptr))
	snprintf.Sig.Variadic = true
	emt.CFuncs["snprintf"] = snprintf

	atoi := emt.Module.NewFunc("atoi", types.I32, ir.NewParam("str", types.I8Ptr))
	emt.CFuncs["atoi"] = atoi

	atol := emt.Module.NewFunc("atol", types.I64, ir.NewParam("str", types.I8Ptr))
	emt.CFuncs["atol"] = atol

	atof := emt.Module.NewFunc("atof", types.Double, ir.NewParam("str", types.I8Ptr))
	emt.CFuncs["atof"] = atof

	gc_init := emt.Module.NewFunc("GC_init", types.I8Ptr)
	emt.CFuncs["gc_init"] = gc_init

	gc_malloc := emt.Module.NewFunc("GC_malloc", types.I8Ptr, ir.NewParam("len", types.I32))
	emt.CFuncs["gc_malloc"] = gc_malloc

	gc_realloc := emt.Module.NewFunc("GC_realloc", types.I8Ptr, ir.NewParam("ptr", types.I8Ptr), ir.NewParam("len", types.I32))
	emt.CFuncs["gc_realloc"] = gc_realloc
}

func (emt *Emitter) EmitClassAndArcReferences(module *ir.Module) {
	// load module
	emt.LoadAndReferenceClasses(module)

	// reference exc functions
	excFuncs := irtools.FindFunctionsWithPrefix(module, "exc_")

	for _, fnc := range excFuncs {
		emt.ExcFuncs[strings.Split(fnc.Name(), "_")[1]] = fnc
		emt.ImportFunction(fnc)
	}
}
