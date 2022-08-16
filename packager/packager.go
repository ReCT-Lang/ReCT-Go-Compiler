package packager

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/irtools"
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"os"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// all packages weve already loaded
var packagesSoFar = make([]symbols.PackageSymbol, 0)

func ResolvePackage(name string) symbols.PackageSymbol {
	// the path where the package *should* be
	path, _ := os.Getwd()
	packagePath := path + "/packages/" + name + ".ll"

	// check if the .ll file exists
	if _, err := os.Stat(packagePath); err != nil {
		print.Error(
			"PACKAGER",
			"lil error",
			0,
			0,
			0,
			"package module file could not be found at path '%s'!",
			packagePath,
		)
		os.Exit(-1)
	}

	// load the LLVM module
	module := irtools.ReadModule(packagePath)

	// load the module's classes
	classes, cls := CreateClassSymbolsFromModule(module)

	// load the module's functions
	funcs := CreateFunctionSymbolsFromModule(name, module, cls)

	return symbols.CreatePackageSymbol(name, funcs, classes)
}

func CreateFunctionSymbolsFromModule(prefix string, module *ir.Module, classes []*symbols.ClassSymbol) []symbols.FunctionSymbol {
	funcs := irtools.FindFunctionsWithPrefix(module, prefix)
	syms := make([]symbols.FunctionSymbol, 0)

	for _, fnc := range funcs {
		print.PrintCF(print.Cyan, "Importing Function '%s'...", fnc.Name())
		syms = append(syms, CreateFunctionSymbolFromModule(fnc, prefix, true, classes))
	}

	return syms
}

func CreateFunctionSymbolFromModule(fnc *ir.Func, prefix string, public bool, classes []*symbols.ClassSymbol) symbols.FunctionSymbol {
	namePure := strings.TrimPrefix(fnc.Name(), prefix)
	returnType := ResolveType(fnc.Sig.RetType, classes)
	params := make([]symbols.ParameterSymbol, 0)

	for i, v := range fnc.Params {
		params = append(params, symbols.CreateParameterSymbol(v.LocalName, i, ResolveType(v.Typ, classes)))
	}

	return symbols.CreateFunctionSymbol(namePure, params, returnType, nodes.FunctionDeclarationMember{}, public)
}

func CreateClassSymbolsFromModule(module *ir.Module) ([]symbols.ClassSymbol, []*symbols.ClassSymbol) {
	classes := make([]*symbols.ClassSymbol, 0)
	clsTypes := make([]types.Type, 0)

	// first load all class names
	for _, typ := range module.TypeDefs {
		// if a type name starts with 'struct.class' it's a class
		if strings.HasPrefix(typ.Name(), "struct.class") {
			// get the class name
			className := strings.Split(typ.Name(), "_")[1]

			// make sure the class is actually defined here and not just referenced
			constructor := irtools.TryFindFunction(module, className+"_public_Constructor")

			// class is declared somewhere else
			if constructor == nil {
				continue
			}

			// this is just a reference, not a defintion
			if len(constructor.Blocks) == 0 {
				continue
			}

			// if this isnt a struct, no idea what it is then
			if !types.IsStruct(typ) {
				continue
			}

			print.PrintCF(print.Cyan, "Importing Class '%s'...", className)

			// create a placeholder symbol
			class := symbols.CreateClassSymbol(
				className,
				nodes.ClassDeclarationMember{},
				make([]symbols.FunctionSymbol, 0),
				make([]symbols.VariableSymbol, 0))

			classes = append(classes, &class)
			clsTypes = append(clsTypes, typ)
		}
	}

	// then load all class functions
	for _, class := range classes {

		// find the constructor
		constructor := irtools.FindFunction(module, class.Name+"_public_Constructor")
		cnst := CreateFunctionSymbolFromModule(constructor, class.Name+"_public_", false, classes)
		cnst.Parameters = cnst.Parameters[1:]

		// alter class object
		class.Functions = append(class.Functions, cnst)

		// find all of its public functions
		classFuncs := irtools.FindFunctionsWithPrefix(module, class.Name+"_public_")
		for _, fnc := range classFuncs {
			// if this isn't the constructor or destructor
			if !strings.HasSuffix(fnc.Name(), "_Constructor") &&
				!strings.HasSuffix(fnc.Name(), "_Die") {
				fncSym := CreateFunctionSymbolFromModule(fnc, class.Name+"_public_", true, classes)

				print.PrintCF(print.Yellow, " Importing Method '%s'...", fncSym.Name)

				class.Functions = append(class.Functions, fncSym)
			}
		}
	}

	// and finally load all the fields
	for i, class := range classes {
		clsType := clsTypes[i].(*types.StructType)
		fields := make([]symbols.VariableSymbol, 0)

		// track down the field list for this class
		fieldConst := irtools.TryFindGlobal(module, class.Name+"_Fields_Const")
		if fieldConst == nil {
			continue
		}

		// read out the constant's value
		fieldNames, ok := irtools.ReadConstStringArray(module, fieldConst)
		if !ok {
			continue
		}

		if len(fieldNames) != len(clsType.Fields)-2 {
			continue
		}

		for i := 2; i < len(clsType.Fields); i++ {
			fieldName := fieldNames[i-2]
			fieldType := ResolveType(clsType.Fields[i], classes)

			print.PrintCF(print.Blue, " Importing Field '%s' (%s)...", fieldName, fieldType.Name)

			fields = append(fields, symbols.CreateGlobalVariableSymbol(
				fieldName,
				false,
				fieldType,
			))
		}

		class.Fields = fields
	}

	// dereference all them pointers
	clsInstances := make([]symbols.ClassSymbol, 0)
	for _, cls := range classes {
		clsInstances = append(clsInstances, *cls)
	}

	return clsInstances, classes
}

func ResolveType(typ types.Type, classes []*symbols.ClassSymbol) symbols.TypeSymbol {

	// =========================================================================
	// PRIMITIVES
	// =========================================================================

	// void type
	if typ.Equal(types.Void) {
		return builtins.Void
	}

	// bool primitive type
	if typ.Equal(types.I1) {
		return builtins.Bool
	}

	// byte primitive type
	if typ.Equal(types.I8) {
		return builtins.Byte
	}

	// int primitive type
	if typ.Equal(types.I32) {
		return builtins.Int
	}

	// float primitive type
	if typ.Equal(types.Float) {
		return builtins.Float
	}

	// =========================================================================
	// OBJECTS
	// =========================================================================
	typeName := ProcessTypeName(typ.LLString())

	// disallow boxed types
	if typeName == "Int" || typeName == "Byte" || typeName == "Float" || typeName == "Bool" {
		print.Error(
			"PACKAGER",
			"lil error",
			0,
			0,
			0,
			"the use of boxed types (object versions of int, byte, float, bool) is not allowed. If you wish to give back an object of a primitive please cast it to 'any'. (Caused by: %s)",
			typ.LLString(),
		)
		os.Exit(-1)
	}

	// string type
	if typeName == "String" {
		return builtins.String
	}

	// thread type
	if typeName == "Thread" {
		return builtins.Thread
	}

	// array types
	// these are a bit wonky as we dont actually know what type of elements they contain
	if typeName == "Array" {
		return builtins.Array
	}
	if typeName == "pArray" {
		return builtins.PArray
	}

	// if its something else, we need to look through our other classes and packages
	for _, cls := range classes {
		if typeName == cls.Name {
			return cls.Type
		}
	}

	for _, pkg := range packagesSoFar {
		for _, cls := range pkg.Classes {
			if typeName == cls.Name {
				return cls.Type
			}
		}
	}

	// aaaaand if we found nothing -> cry
	print.Error(
		"PACKAGER",
		"lil error",
		0,
		0,
		0,
		"could not resolve referenced type '%s' while loading package!",
		typ.LLString(),
	)
	os.Exit(-1)

	return symbols.TypeSymbol{}
}

func ProcessTypeName(name string) string {
	// if this type name doesnt match the rect class pattern and also isnt a primitive
	// => no idea what the fuck this is
	if !strings.HasPrefix(name, "%struct.class_") {
		print.Error(
			"PACKAGER",
			"lil error",
			0,
			0,
			0,
			"referenced type '%s' does not match ReCT class pattern! Absolutely no clue what todo with it lol",
			name,
		)
		os.Exit(-1)
	}

	// if this is a valid class but not a pointer -> hm?????
	// (objects are referential types so they NEED to be pointers)
	if !strings.HasSuffix(name, "*") {
		print.Error(
			"PACKAGER",
			"lil error",
			0,
			0,
			0,
			"referenced object type '%s' needs to be a pointer but isnt!",
			name,
		)
		os.Exit(-1)
	}

	// if all those things are alright we can cut away all the un-needed stuff
	return strings.TrimSuffix(strings.TrimPrefix(name, "%struct.class_"), "*")
}
