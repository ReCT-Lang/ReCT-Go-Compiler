package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func (emt *Emitter) EmitBuiltInFunctions() {

	// link our classes and arc
	emt.EmitClassAndArcReferences()

	// link referecnes to the C standard libs
	emt.EmitCLibReferences()

	// link our built-ins
	emt.EmitSystemFuncReferences()

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

func (emt *Emitter) EmitSystemFuncReferences() {
	Print := emt.Module.NewFunc("rct_Print", types.Void, ir.NewParam("text", emt.IRTypes(builtins.String.Fingerprint())))
	emt.Functions[emt.Id(builtins.Print)] = Function{IRFunction: Print, BoundFunction: binder.BoundFunction{Symbol: builtins.Print}}

	Write := emt.Module.NewFunc("rct_Write", types.Void, ir.NewParam("text", emt.IRTypes(builtins.String.Fingerprint())))
	emt.Functions[emt.Id(builtins.Write)] = Function{IRFunction: Write, BoundFunction: binder.BoundFunction{Symbol: builtins.Write}}

	Input := emt.Module.NewFunc("rct_Input", emt.IRTypes(builtins.String.Fingerprint()))
	emt.Functions[emt.Id(builtins.Input)] = Function{IRFunction: Input, BoundFunction: binder.BoundFunction{Symbol: builtins.Input}}

	Clear := emt.Module.NewFunc("rct_Clear", types.Void)
	emt.Functions[emt.Id(builtins.Clear)] = Function{IRFunction: Clear, BoundFunction: binder.BoundFunction{Symbol: builtins.Clear}}

	SetCursor := emt.Module.NewFunc("rct_SetCursor", types.Void, ir.NewParam("x", types.I32), ir.NewParam("y", types.I32))
	emt.Functions[emt.Id(builtins.SetCursor)] = Function{IRFunction: SetCursor, BoundFunction: binder.BoundFunction{Symbol: builtins.SetCursor}}

	SetCursorVisible := emt.Module.NewFunc("rct_SetCursorVisible", types.Void, ir.NewParam("x", types.I1))
	emt.Functions[emt.Id(builtins.SetCursorVisible)] = Function{IRFunction: SetCursorVisible, BoundFunction: binder.BoundFunction{Symbol: builtins.SetCursorVisible}}

	GetCursorVisible := emt.Module.NewFunc("rct_GetCursorVisible", types.I1)
	emt.Functions[emt.Id(builtins.GetCursorVisible)] = Function{IRFunction: GetCursorVisible, BoundFunction: binder.BoundFunction{Symbol: builtins.GetCursorVisible}}

	Random := emt.Module.NewFunc("rct_Random", types.I32, ir.NewParam("maxValue", types.I32))
	emt.Functions[emt.Id(builtins.Random)] = Function{IRFunction: Random, BoundFunction: binder.BoundFunction{Symbol: builtins.Random}}

	Sleep := emt.Module.NewFunc("rct_Sleep", types.Void, ir.NewParam("ms", types.I32))
	emt.Functions[emt.Id(builtins.Sleep)] = Function{IRFunction: Sleep, BoundFunction: binder.BoundFunction{Symbol: builtins.Sleep}}
}

func (emt *Emitter) EmitClassAndArcReferences() {

	// function pointer type for destructors
	dieFunction := types.NewPointer(types.NewFunc(types.Void, types.I8Ptr))

	// vTables
	Any_vTable := emt.Module.NewTypeDef("%struct.Any_vTable", types.NewStruct(types.I8Ptr, types.I8Ptr, dieFunction))
	String_vTable := emt.Module.NewTypeDef("%struct.String_vTable", types.NewStruct(types.NewPointer(Any_vTable), types.I8Ptr, dieFunction))
	Int_vTable := emt.Module.NewTypeDef("%struct.Int_vTable", types.NewStruct(types.NewPointer(Any_vTable), types.I8Ptr, dieFunction))
	Float_vTable := emt.Module.NewTypeDef("%struct.Float_vTable", types.NewStruct(types.NewPointer(Any_vTable), types.I8Ptr, dieFunction))
	Bool_vTable := emt.Module.NewTypeDef("%struct.Bool_vTable", types.NewStruct(types.NewPointer(Any_vTable), types.I8Ptr, dieFunction))

	// load all classes
	class_Any := emt.Module.NewTypeDef("%struct.class_Any", types.NewStruct(Any_vTable, types.I32))
	class_String := emt.Module.NewTypeDef("%struct.class_String", types.NewStruct(String_vTable, types.I32, types.I8Ptr, types.I32, types.I32, types.I32))
	class_Int := emt.Module.NewTypeDef("%struct.class_Int", types.NewStruct(Int_vTable, types.I32, types.I32))
	class_Float := emt.Module.NewTypeDef("%struct.class_Float", types.NewStruct(Float_vTable, types.I32, types.Float))
	class_Bool := emt.Module.NewTypeDef("%struct.class_Bool", types.NewStruct(Bool_vTable, types.I32, types.I8))

	// load all of their constructors
	Any_public_Constructor := emt.Module.NewFunc("Any_public_Constructor", types.Void, ir.NewParam("this", types.NewPointer(class_Any)))
	String_public_Constructor := emt.Module.NewFunc("String_public_Constructor", types.Void, ir.NewParam("this", types.NewPointer(class_String)))
	Int_public_Constructor := emt.Module.NewFunc("Int_public_Constructor", types.Void, ir.NewParam("this", types.NewPointer(class_Int)), ir.NewParam("value", types.I32))
	Float_public_Constructor := emt.Module.NewFunc("Float_public_Constructor", types.Void, ir.NewParam("this", types.NewPointer(class_Float)), ir.NewParam("value", types.Float))
	Bool_public_Constructor := emt.Module.NewFunc("Bool_public_Constructor", types.Void, ir.NewParam("this", types.NewPointer(class_Bool)), ir.NewParam("value", types.I8))

	// load the string.Load() function
	String_public_Load := emt.Module.NewFunc("String_public_Load", types.Void, ir.NewParam("this", types.NewPointer(class_String)), ir.NewParam("source", types.I8Ptr))

	// find out what names to use for the classes
	anyName := emt.Id(builtins.Any)
	stringName := emt.Id(builtins.String)
	intName := emt.Id(builtins.Int)
	floatName := emt.Id(builtins.Float)
	boolName := emt.Id(builtins.Bool)

	// store all of them gamers
	emt.Classes[anyName] = Class{Type: class_Any, Constructor: Any_public_Constructor, Functions: make(map[string]*ir.Func)}
	emt.Classes[stringName] = Class{Type: class_String, Constructor: String_public_Constructor, Functions: make(map[string]*ir.Func)}
	emt.Classes[intName] = Class{Type: class_Int, Constructor: Int_public_Constructor, Functions: make(map[string]*ir.Func)}
	emt.Classes[floatName] = Class{Type: class_Float, Constructor: Float_public_Constructor, Functions: make(map[string]*ir.Func)}
	emt.Classes[boolName] = Class{Type: class_Bool, Constructor: Bool_public_Constructor, Functions: make(map[string]*ir.Func)}

	// store string functions
	emt.Classes[stringName].Functions["load"] = String_public_Load

	registerReference := emt.Module.NewFunc("arc_RegisterReference", types.Void, ir.NewParam("obj", types.NewPointer(class_Any)))
	emt.ArcFuncs["registerReference"] = registerReference

	unregisterReference := emt.Module.NewFunc("arc_UnregisterReference", types.Void, ir.NewParam("obj", types.NewPointer(class_Any)))
	emt.ArcFuncs["dieReference"] = unregisterReference

	registerVerboseReference := emt.Module.NewFunc("arc_RegisterReferenceVerbose", types.Void, ir.NewParam("obj", types.NewPointer(class_Any)), ir.NewParam("comment", types.I8Ptr))
	emt.ArcFuncs["registerVerboseReference"] = registerVerboseReference

	unregisterVerboseReference := emt.Module.NewFunc("arc_UnregisterReferenceVerbose", types.Void, ir.NewParam("obj", types.NewPointer(class_Any)), ir.NewParam("comment", types.I8Ptr))
	emt.ArcFuncs["dieVerboseReference"] = unregisterVerboseReference
}
