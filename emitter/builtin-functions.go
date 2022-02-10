package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func (emt *Emitter) EmitBuiltInFunctions() {

	// emit referecnes to the C standard libs
	emt.EmitCLibReferences()

	// emit our built-ins

	// output functions
	emptyFormat := emt.Module.NewGlobalDef(".str.efmt", constant.NewCharArrayFromString("%s\x00"))
	emptyFormat.Immutable = true
	newlineFormat := emt.Module.NewGlobalDef(".str.nlfmt", constant.NewCharArrayFromString("%s\n\x00"))
	newlineFormat.Immutable = true

	emt.EmitPrint(newlineFormat)
	emt.EmitWrite(emptyFormat)

	// input functions
	emt.EmitInput()

}

func (emt *Emitter) EmitCLibReferences() {
	printf := emt.Module.NewFunc("printf", types.I32, ir.NewParam("format", types.I8Ptr))
	printf.Sig.Variadic = true
	emt.CFunctions["printf"] = printf

	scanf := emt.Module.NewFunc("scanf", types.I32, ir.NewParam("format", types.I8Ptr), ir.NewParam("dest", types.I8Ptr))
	scanf.Sig.Variadic = true
	emt.CFunctions["scanf"] = scanf

	strcpy := emt.Module.NewFunc("strcpy", types.I8Ptr, ir.NewParam("dest", types.I8Ptr), ir.NewParam("src", types.I8Ptr))
	strcpy.Sig.Variadic = true
	emt.CFunctions["strcpy"] = strcpy

	strcat := emt.Module.NewFunc("strcat", types.I8Ptr, ir.NewParam("dest", types.I8Ptr), ir.NewParam("src", types.I8Ptr))
	strcat.Sig.Variadic = true
	emt.CFunctions["strcat"] = strcat

	strlen := emt.Module.NewFunc("strlen", types.I32, ir.NewParam("str", types.I8Ptr))
	strlen.Sig.Variadic = true
	emt.CFunctions["strlen"] = strlen

	malloc := emt.Module.NewFunc("malloc", types.I8Ptr, ir.NewParam("len", types.I32))
	malloc.Sig.Variadic = true
	emt.CFunctions["malloc"] = malloc

	free := emt.Module.NewFunc("free", types.Void, ir.NewParam("dest", types.I8Ptr))
	free.Sig.Variadic = true
	emt.CFunctions["free"] = free
}

func (emt *Emitter) EmitPrint(newlineFormat *ir.Global) {

	// figure out what name to use (name / fingerprint)
	PrintName := tern(emt.UseFingerprints, builtins.Print.Fingerprint(), builtins.Print.Name)
	Print := emt.Module.NewFunc(PrintName, types.Void, ir.NewParam("message", types.I8Ptr))

	emt.Functions[PrintName] = Function{IRFunction: Print, BoundFunction: binder.BoundFunction{Symbol: builtins.Print}}
	body := Print.NewBlock("")

	// our newline format
	newlineFormatPointer := body.NewGetElementPtr(types.NewArray(4, types.I8), newlineFormat, CI32(0), CI32(0))

	body.NewCall(emt.CFunctions["printf"], newlineFormatPointer, Print.Params[0])
	body.NewRet(nil)
}

func (emt *Emitter) EmitWrite(emptyFormat *ir.Global) {

	// figure out what name to use (name / fingerprint)
	WriteName := tern(emt.UseFingerprints, builtins.Write.Fingerprint(), builtins.Write.Name)
	Write := emt.Module.NewFunc(WriteName, types.Void, ir.NewParam("message", types.I8Ptr))

	emt.Functions[WriteName] = Function{IRFunction: Write, BoundFunction: binder.BoundFunction{Symbol: builtins.Write}}
	body := Write.NewBlock("")

	// our empty format
	emptyFormatPointer := body.NewGetElementPtr(types.NewArray(3, types.I8), emptyFormat, CI32(0), CI32(0))

	// print out the string
	body.NewCall(emt.CFunctions["printf"], emptyFormatPointer, Write.Params[0])
	body.NewRet(nil)
}

func (emt *Emitter) EmitInput() {
	// input format constant
	inputFormat := emt.Module.NewGlobalDef(".str.ifmt", constant.NewCharArrayFromString("%1023[^\n]\x00"))
	inputFormat.Immutable = true

	// figure out what name to use (name / fingerprint)
	InputName := tern(emt.UseFingerprints, builtins.Input.Fingerprint(), builtins.Input.Name)
	Input := emt.Module.NewFunc(InputName, types.I8Ptr)

	emt.Functions[InputName] = Function{IRFunction: Input, BoundFunction: binder.BoundFunction{Symbol: builtins.Input}}
	body := Input.NewBlock("")

	// our input format
	inputFormatPointer := body.NewGetElementPtr(types.NewArray(10, types.I8), inputFormat, CI32(0), CI32(0))

	// input buffer (limited to 1024 characters)
	buffer := body.NewCall(emt.CFunctions["malloc"], CI32(1024))

	// use scanf to read in a line (again limited to 1024 characters)
	body.NewCall(emt.CFunctions["scanf"], inputFormatPointer, buffer)

	// copy the input over to another buffer
	// (so we dont use any more space than we need to)
	newStr := body.NewCall(emt.CFunctions["malloc"],
		// with a size of strlen(str) + 1
		body.NewAdd(
			body.NewCall(emt.CFunctions["strlen"], buffer),
			CI32(1),
		),
	)
	body.NewCall(emt.CFunctions["strcpy"], newStr, buffer)

	// free the buffer
	body.NewCall(emt.CFunctions["free"], buffer)

	// print out the string
	body.NewRet(newStr)
}
