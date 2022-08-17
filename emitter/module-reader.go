package emitter

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/irtools"
	"ReCT-Go-Compiler/print"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

func (emt *Emitter) LoadAndReferenceClasses(module *ir.Module) {
	// reference all classes

	// first load all class names
	for _, typ := range module.TypeDefs {
		// if a type name starts with 'struct.class' it's a class
		if strings.HasPrefix(typ.Name(), "struct.class") {
			// get the class name
			className := strings.Split(typ.Name(), "_")[1]
			typeSymbol := builtins.Error

			// find out what type symbol this is for
			for _, sym := range builtins.Types {
				if strings.ToLower(sym.SymbolName()) == strings.ToLower(className) {
					typeSymbol = sym
					break
				}
			}

			// find and link the vtable
			// ------------------------

			// 1. format the name (removing % prefix and * suffix)
			vTableType := typ.(*types.StructType).Fields[0].String()
			vTableType = vTableType[1 : len(vTableType)-1]

			// 2. finding and importing the type
			vTable := FindType(module, vTableType)
			emt.ImportType(vTable)

			// 3. finding and importing the types vtable constant
			vConstantName := strings.Split(vTableType, ".")[1] + "_Const"
			vTableConstant := irtools.FindGlobal(module, vConstantName)

			if !GlobalExists(emt.Module, vTableConstant.GlobalName) {
				emt.Module.NewGlobal(vTableConstant.GlobalName, vTable).Linkage = enum.LinkageExternal
			}

			// create a class object
			emt.Classes[emt.Id(typeSymbol)] = &Class{Type: typ, vTable: vTable, vConstant: vTableConstant, Constructor: nil, Functions: make(map[string]*ir.Func), Name: className}
			emt.ImportType(typ)
		}
	}

	// then load all class functions
	for key, class := range emt.Classes {

		// find the constructor
		constructor := irtools.FindFunction(module, class.Name+"_public_Constructor")
		emt.ImportFunction(constructor)

		// find the destructor
		destructor := irtools.FindFunction(module, class.Name+"_public_Die")
		emt.ImportFunction(destructor)

		// alter class object
		class.Constructor = constructor
		class.Destructor = destructor
		emt.Classes[key] = class

		// find all of its public functions
		classFuncs := irtools.FindFunctionsWithPrefix(module, class.Name+"_public_")
		for _, fnc := range classFuncs {
			// if this isn't the constructor or destructor
			if !strings.HasSuffix(fnc.Name(), "_Constructor") &&
				!strings.HasSuffix(fnc.Name(), "_Die") {
				emt.Classes[key].Functions[strings.Split(fnc.Name(), "_")[2]] = fnc
				emt.ImportFunction(fnc)
			}
		}
	}
}

func GlobalExists(module *ir.Module, name string) bool {
	for _, glb := range module.Globals {
		if glb.Name() == name {
			return true
		}
	}

	return false
}

func FindType(module *ir.Module, name string) types.Type {
	for _, typ := range module.TypeDefs {
		if typ.Name() == name {
			return typ
		}
	}

	print.PrintC(print.Red, "Couldnt find type '"+name+"'")
	return nil
}

func (emt *Emitter) TypeExists(name string) bool {
	for _, typ := range emt.Module.TypeDefs {
		if typ.Name() == name {
			return true
		}
	}

	return false
}

func (emt *Emitter) ImportType(typ types.Type) {
	// check if type already exists
	if !emt.TypeExists(typ.Name()) {
		emt.Module.NewTypeDef(typ.Name(), typ)
	}

}

func (emt *Emitter) ImportFunction(fnc *ir.Func) *ir.Func {
	return emt.Module.NewFunc(fnc.Name(), fnc.Sig.RetType, fnc.Params...)
}
