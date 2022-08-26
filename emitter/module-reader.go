package emitter

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/irtools"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"os"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

func (emt *Emitter) ImportPackage(pack symbols.PackageSymbol) {
	// create a package object
	pck := Package{
		Functions: make(map[string]*ir.Func),
	}

	// load the package's classes
	emt.LoadAndReferenceClassesFromPackage(pack.Module, pack)

	// load the  package's functions
	for _, fnc := range pack.Functions {
		pck.Functions[emt.Id(fnc)] = fnc.IRFunction
		emt.ImportFunction(fnc.IRFunction)
	}

	// store this package
	emt.Packages[emt.Id(pack)] = &pck
}

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

func (emt *Emitter) LoadAndReferenceClassesFromPackage(module *ir.Module, pack symbols.PackageSymbol) {
	packClasses := make(map[*symbols.ClassSymbol]*Class)

	// reference all classes
	for _, cls := range pack.Classes {
		typ := cls.IRType

		// find and link the vtable
		// ------------------------

		// 1. format the name (removing % prefix and * suffix)
		vTableType := typ.(*types.StructType).Fields[0].String()
		vTableType = vTableType[1 : len(vTableType)-1]

		// 2. finding and importing the type
		vTable := FindType(module, vTableType)
		if vTable == nil {
			print.Error(
				"EMITTER",
				"ow", 0, 0, 0,
				"Couldn't find vTable for class \"%s\" from package \"%s\"! Is the package set up correctly?",
				cls.Name,
				pack.Name,
			)
			os.Exit(-1)
		}

		emt.ImportType(vTable)

		// 3. finding and importing the types vtable constant
		vConstantName := cls.Name + "_vTable_Const"
		vTableConstant := irtools.FindGlobalSuffix(module, vConstantName)
		if vTableConstant == nil {
			print.Error(
				"EMITTER",
				"ow", 0, 0, 0,
				"Couldn't find vTable-Constant for class \"%s\" from package \"%s\"! Is the package set up correctly?",
				cls.Name,
				pack.Name,
			)
			os.Exit(-1)
		}

		if !GlobalExists(emt.Module, vTableConstant.GlobalName) {
			emt.Module.NewGlobal(vTableConstant.GlobalName, vTable).Linkage = enum.LinkageExternal
		}

		// create a class object
		clsCpy := cls
		packClasses[&clsCpy] = &Class{Type: typ, vTable: vTable, vConstant: vTableConstant, Constructor: nil, Functions: make(map[string]*ir.Func), Name: cls.Name}
		emt.ImportType(typ)
	}

	// then load all class functions
	for key, class := range packClasses {

		// find the constructor
		constructor := irtools.FindFunction(module, class.Name+"_public_Constructor")
		if constructor == nil {
			print.Error(
				"EMITTER",
				"ow", 0, 0, 0,
				"Couldn't find Constructor for class \"%s\" from package \"%s\"! Is the package set up correctly?",
				class.Name,
				pack.Name,
			)
			os.Exit(-1)
		}
		emt.ImportFunction(constructor)

		// find the destructor
		destructor := irtools.FindFunction(module, class.Name+"_public_Die")
		if destructor == nil {
			print.Error(
				"EMITTER",
				"ow", 0, 0, 0,
				"Couldn't find Destructor for class \"%s\" from package \"%s\"! Is the package set up correctly?",
				class.Name,
				pack.Name,
			)
			os.Exit(-1)
		}
		emt.ImportFunction(destructor)

		// alter class object
		class.Constructor = constructor
		class.Destructor = destructor
		packClasses[key] = class

		// import all of its public functions
		for _, fnc := range key.Functions {
			// if this isn't the constructor or destructor
			if fnc.Name != "Constructor" && fnc.Name != "Die" {
				packClasses[key].Functions[emt.Id(fnc)] = fnc.IRFunction
				emt.ImportFunction(fnc.IRFunction)
			}
		}
	}

	// then load all class fields
	for key, class := range packClasses {
		fieldMap := make(map[string]int)
		for i, fld := range key.Fields {
			fieldMap[fld.Fingerprint()] = i + 2
		}

		class.Fields = fieldMap
	}

	for k, v := range packClasses {
		emt.Classes[emt.Id(k.Type)] = v
	}

	// then load literally all types to clear up any loose ends
	for _, typ := range module.TypeDefs {
		emt.ImportType(typ)
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
