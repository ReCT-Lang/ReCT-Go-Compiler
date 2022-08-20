package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/symbols"
	"fmt"
	"os"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// this file is just keeping track of how ReCT types map to LLVM types

var arrayTypes = make(map[string]types.Type)
var parrayTypes = make(map[string]types.Type)

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
	case builtins.Long.Fingerprint():
		return types.I64
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
			return emt.ResolveArray(typ, &arrayTypes, builtins.Array)
		} else {
			return emt.ResolveArray(typ, &parrayTypes, builtins.PArray)
		}
	}

	// try looking up a class
	cls, ok := emt.Classes[emt.Id(typ)]
	if ok {
		return types.NewPointer(cls.Type)
	}

	fmt.Println("Unknown Type")
	fmt.Println(typ.Fingerprint())

	for _, v := range emt.Classes {
		fmt.Println("> " + v.Name)
	}

	os.Exit(-1)
	return nil
}

func (emt *Emitter) ResolveArray(typ symbols.TypeSymbol, cache *map[string]types.Type, generic symbols.TypeSymbol) types.Type {
	// see if this type already exists
	arrType, ok := (*cache)[typ.Fingerprint()]
	if ok {
		return arrType
	}

	// if not, copy the array type and rename it
	genericType := emt.Classes[emt.Id(generic)].Type.(*types.StructType)

	// figure out the suffix
	name := typ.SubTypes[0].Name
	suffix := strings.ToUpper(string(name[0])) + name[1:]
	if len(typ.SubTypes[0].SubTypes) > 0 {
		suffix = typ.SubTypes[0].Fingerprint()

		// escape symbols which arent allowed in c
		suffix = strings.Replace(suffix, "[", "$b$", -1)
		suffix = strings.Replace(suffix, "]", "$e$", -1)
		suffix = strings.Replace(suffix, ";", "$s$", -1)
	}

	// figure out the prefix
	prefix := "class_Array_"
	if generic.Fingerprint() == builtins.PArray.Fingerprint() {
		prefix = "class_pArray_"
	}

	var newType types.Type

	// check if this type has already been imported
	if emt.TypeExists("struct." + prefix + suffix) {
		newType = FindType(emt.Module, "struct."+prefix+suffix)

	} else {
		// otherwise, create it
		blueprint := &types.StructType{}
		blueprint.Fields = genericType.Fields

		newType = emt.Module.NewTypeDef("struct."+prefix+suffix, blueprint)
	}

	newType = types.NewPointer(newType)

	// cache this type definition
	(*cache)[typ.Fingerprint()] = newType

	return newType
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
	vTable      types.Type
	vConstant   *ir.Global
	Constructor *ir.Func
	Destructor  *ir.Func
	Functions   map[string]*ir.Func
	Fields      map[string]int
	Name        string
}

type Package struct {
	Functions map[string]*ir.Func
}
