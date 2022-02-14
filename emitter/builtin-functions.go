package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func (emt *Emitter) EmitBuiltInFunctions() {

	// emit referecnes to the C standard libs
	emt.EmitCLibReferences()

	// emit our built-ins
	emt.EmitSystemFuncReferences()
}

func (emt *Emitter) EmitCLibReferences() {
	printf := emt.Module.NewFunc("printf", types.I32, ir.NewParam("format", types.I8Ptr))
	printf.Sig.Variadic = true
	emt.CFunctions["printf"] = printf

	scanf := emt.Module.NewFunc("scanf", types.I32, ir.NewParam("format", types.I8Ptr), ir.NewParam("dest", types.I8Ptr))
	scanf.Sig.Variadic = true
	emt.CFunctions["scanf"] = scanf

	strcpy := emt.Module.NewFunc("strcpy", types.Void, ir.NewParam("dest", types.I8Ptr), ir.NewParam("src", types.I8Ptr))
	strcpy.Sig.Variadic = true
	emt.CFunctions["strcpy"] = strcpy

	strcat := emt.Module.NewFunc("strcat", types.Void, ir.NewParam("dest", types.I8Ptr), ir.NewParam("src", types.I8Ptr))
	strcat.Sig.Variadic = true
	emt.CFunctions["strcat"] = strcat

	strlen := emt.Module.NewFunc("strlen", types.I32, ir.NewParam("str", types.I8Ptr))
	strlen.Sig.Variadic = true
	emt.CFunctions["strlen"] = strlen

	strcmp := emt.Module.NewFunc("strcmp", types.I32, ir.NewParam("left", types.I8Ptr), ir.NewParam("right", types.I8Ptr))
	strcmp.Sig.Variadic = true
	emt.CFunctions["strcmp"] = strcmp

	malloc := emt.Module.NewFunc("malloc", types.I8Ptr, ir.NewParam("len", types.I32))
	malloc.Sig.Variadic = true
	emt.CFunctions["malloc"] = malloc

	free := emt.Module.NewFunc("free", types.Void, ir.NewParam("dest", types.I8Ptr))
	free.Sig.Variadic = true
	emt.CFunctions["free"] = free
}

func (emt *Emitter) EmitSystemFuncReferences() {
	Print := emt.Module.NewFunc("rct_Print", types.Void, ir.NewParam("text", types.I8Ptr))
	Print.Sig.Variadic = true
	emt.Functions[tern(emt.UseFingerprints, builtins.Print.Fingerprint(), builtins.Print.Name)] = Function{IRFunction: Print, BoundFunction: binder.BoundFunction{Symbol: builtins.Print}}

	Write := emt.Module.NewFunc("rct_Write", types.Void, ir.NewParam("text", types.I8Ptr))
	Write.Sig.Variadic = true
	emt.Functions[tern(emt.UseFingerprints, builtins.Write.Fingerprint(), builtins.Write.Name)] = Function{IRFunction: Write, BoundFunction: binder.BoundFunction{Symbol: builtins.Write}}

	Input := emt.Module.NewFunc("rct_Input", types.I8Ptr)
	Input.Sig.Variadic = true
	emt.Functions[tern(emt.UseFingerprints, builtins.Input.Fingerprint(), builtins.Input.Name)] = Function{IRFunction: Input, BoundFunction: binder.BoundFunction{Symbol: builtins.Input}}

	Clear := emt.Module.NewFunc("rct_Clear", types.Void)
	Clear.Sig.Variadic = true
	emt.Functions[tern(emt.UseFingerprints, builtins.Clear.Fingerprint(), builtins.Clear.Name)] = Function{IRFunction: Clear, BoundFunction: binder.BoundFunction{Symbol: builtins.Clear}}

	SetCursor := emt.Module.NewFunc("rct_SetCursor", types.Void, ir.NewParam("x", types.I32), ir.NewParam("y", types.I32))
	SetCursor.Sig.Variadic = true
	emt.Functions[tern(emt.UseFingerprints, builtins.SetCursor.Fingerprint(), builtins.SetCursor.Name)] = Function{IRFunction: SetCursor, BoundFunction: binder.BoundFunction{Symbol: builtins.SetCursor}}

	SetCursorVisible := emt.Module.NewFunc("rct_SetCursorVisible", types.Void, ir.NewParam("x", types.I1))
	SetCursorVisible.Sig.Variadic = true
	emt.Functions[tern(emt.UseFingerprints, builtins.SetCursorVisible.Fingerprint(), builtins.SetCursorVisible.Name)] = Function{IRFunction: SetCursorVisible, BoundFunction: binder.BoundFunction{Symbol: builtins.SetCursorVisible}}

	GetCursorVisible := emt.Module.NewFunc("rct_GetCursorVisible", types.I1)
	GetCursorVisible.Sig.Variadic = true
	emt.Functions[tern(emt.UseFingerprints, builtins.GetCursorVisible.Fingerprint(), builtins.GetCursorVisible.Name)] = Function{IRFunction: GetCursorVisible, BoundFunction: binder.BoundFunction{Symbol: builtins.GetCursorVisible}}

	Random := emt.Module.NewFunc("rct_Random", types.I32, ir.NewParam("maxValue", types.I32))
	Random.Sig.Variadic = true
	emt.Functions[tern(emt.UseFingerprints, builtins.Random.Fingerprint(), builtins.Random.Name)] = Function{IRFunction: Random, BoundFunction: binder.BoundFunction{Symbol: builtins.Random}}

	Sleep := emt.Module.NewFunc("rct_Sleep", types.Void, ir.NewParam("ms", types.I32))
	Sleep.Sig.Variadic = true
	emt.Functions[tern(emt.UseFingerprints, builtins.Sleep.Fingerprint(), builtins.Sleep.Name)] = Function{IRFunction: Sleep, BoundFunction: binder.BoundFunction{Symbol: builtins.Sleep}}

	stringCopy := emt.Module.NewFunc("util_copy_string", types.I8Ptr, ir.NewParam("source", types.I8Ptr))
	stringCopy.Sig.Variadic = true
	emt.CFunctions["uStringCopy"] = stringCopy

	freeIfNotNull := emt.Module.NewFunc("util_free_string_if_not_null", types.I8Ptr, ir.NewParam("source", types.I8Ptr))
	freeIfNotNull.Sig.Variadic = true
	emt.CFunctions["uFreeIfNotNull"] = freeIfNotNull
}
