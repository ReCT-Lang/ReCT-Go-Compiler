package packager

import (
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/irtools"
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"os"
	"strings"
	"unicode"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

var PackagePaths []string

// all packages weve already loaded
var PackagesSoFar = make([]symbols.PackageSymbol, 0)

func ResolvePackage(name string, errorLocation print.TextSpan) symbols.PackageSymbol {
	// the path where the package *should* be
	packagePath := "/" + name + ".ll"
	exists := false

	for _, pth := range PackagePaths {
		checkPath := pth + packagePath

		// check if the .ll file exists
		if _, err := os.Stat(checkPath); err == nil {
			// we good
			packagePath = checkPath
			exists = true
			break
		}
	}

	// if we didnt find anything
	if !exists {
		if len(PackagePaths) == 1 {
			print.Error(
				"PACKAGER",
				print.UnknownPackageModuleFileError,
				errorLocation,
				"Package module file could not be found at path '%s'!",
				PackagePaths[0]+packagePath,
			)
		} else {
			print.Error(
				"PACKAGER",
				print.UnknownPackageModuleFileError,
				errorLocation,
				"Package module file could not be found in any of the given package directories!",
			)
		}
		os.Exit(-1)
	}

	// load the LLVM module
	module := irtools.ReadModule(packagePath)

	// load the module's classes
	classes, cls := CreateClassSymbolsFromModule(module, errorLocation)

	// load the module's functions
	funcs := CreateFunctionSymbolsFromModule(name+"_", module, cls, errorLocation)

	// create a package symbol
	pack := symbols.CreatePackageSymbol(name, funcs, classes, module, errorLocation)
	PackagesSoFar = append(PackagesSoFar, pack)

	return pack
}

func CreateFunctionSymbolsFromModule(prefix string, module *ir.Module, classes []*symbols.ClassSymbol, errorLocation print.TextSpan) []symbols.FunctionSymbol {
	funcs := irtools.FindFunctionsWithPrefix(module, prefix)
	syms := make([]symbols.FunctionSymbol, 0)

	for _, fnc := range funcs {
		namePure := strings.TrimPrefix(fnc.Name(), prefix)

		// check if this is reeeeeally a function and not a method in disguise
		if strings.HasPrefix(namePure, "public_") || strings.HasPrefix(namePure, "private_") {
			continue
		}

		fncSym := CreateFunctionSymbolFromModule(fnc, prefix, true, classes, errorLocation)
		//print.PrintCF(print.Cyan, "Importing Function '%s'...", fncSym.Name)
		syms = append(syms, fncSym)
	}

	return syms
}

func CreateFunctionSymbolFromModule(fnc *ir.Func, prefix string, public bool, classes []*symbols.ClassSymbol, errorLocation print.TextSpan) symbols.FunctionSymbol {
	namePure := strings.TrimPrefix(fnc.Name(), prefix)
	returnType, ok := ResolveType(fnc.Sig.RetType, classes, errorLocation)
	if !ok {
		print.Error(
			"PACKAGER",
			print.ImpossibleFunctionProcessingError,
			errorLocation,
			"Unable to process function '%s'!",
			fnc.Name(),
		)
		os.Exit(-1)
	}

	params := make([]symbols.ParameterSymbol, 0)

	for i, v := range fnc.Params {
		typ, ok := ResolveType(v.Typ, classes, errorLocation)
		if !ok {
			print.Error(
				"PACKAGER",
				print.ImpossibleFunctionProcessingError,
				errorLocation,
				"Unable to process function '%s'!",
				fnc.Name(),
			)
			os.Exit(-1)
		}

		params = append(params, symbols.CreateParameterSymbol(v.LocalName, i, typ))
	}

	symbol := symbols.CreateFunctionSymbol(namePure, params, returnType, nodes.FunctionDeclarationMember{}, public)
	symbol.IRFunction = fnc
	return symbol
}

func CreateClassSymbolsFromModule(module *ir.Module, errorLocation print.TextSpan) ([]symbols.ClassSymbol, []*symbols.ClassSymbol) {
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

			//print.PrintCF(print.Cyan, "Importing Class '%s'...", className)

			// create a placeholder symbol
			class := symbols.CreateClassSymbol(
				className,
				nodes.ClassDeclarationMember{},
				make([]symbols.FunctionSymbol, 0),
				make([]symbols.VariableSymbol, 0))

			class.IRType = typ
			classes = append(classes, &class)
			clsTypes = append(clsTypes, typ)
		}
	}

	// then load all class functions
	for _, class := range classes {

		// find the constructor
		constructor := irtools.FindFunction(module, class.Name+"_public_Constructor")
		cnst := CreateFunctionSymbolFromModule(constructor, class.Name+"_public_", false, classes, errorLocation)
		cnst.Parameters = cnst.Parameters[1:]

		// alter class object
		class.Functions = append(class.Functions, cnst)

		// find all of its public functions
		classFuncs := irtools.FindFunctionsWithPrefix(module, class.Name+"_public_")
		for _, fnc := range classFuncs {
			// if this isn't the constructor
			if !strings.HasSuffix(fnc.Name(), "_Constructor") {
				fncSym := CreateFunctionSymbolFromModule(fnc, class.Name+"_public_", true, classes, errorLocation)
				fncSym.Parameters = fncSym.Parameters[1:]

				//print.PrintCF(print.Yellow, " Importing Method '%s'...", fncSym.Name)

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
			fieldName := strings.TrimSuffix(fieldNames[i-2], "\x00")
			fieldType, ok := ResolveType(clsType.Fields[i], classes, errorLocation)
			if !ok {
				print.Error(
					"PACKAGER",
					print.ImpossibleFieldProcessingError,
					errorLocation,
					"Unable to process field '%s' of class '%s'!",
					fieldName, class.Name,
				)
				os.Exit(-1)
			}

			//print.PrintCF(print.Blue, " Importing Field '%s' (%s)...", fieldName, fieldType.Name)

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

func ResolveType(typ types.Type, classes []*symbols.ClassSymbol, errorLocation print.TextSpan) (symbols.TypeSymbol, bool) {

	// =========================================================================
	// PRIMITIVES
	// =========================================================================

	// void type
	if typ.Equal(types.Void) {
		return builtins.Void, true
	}

	// bool primitive type
	if typ.Equal(types.I1) {
		return builtins.Bool, true
	}

	// byte primitive type
	if typ.Equal(types.I8) {
		return builtins.Byte, true
	}

	// int primitive type
	if typ.Equal(types.I32) {
		return builtins.Int, true
	}

	// long primitive type
	if typ.Equal(types.I64) {
		return builtins.Long, true
	}

	// float primitive type
	if typ.Equal(types.Float) {
		return builtins.Float, true
	}

	// double primitive type
	if typ.Equal(types.Double) {
		return builtins.Double, true
	}

	// =========================================================================
	// OBJECTS
	// =========================================================================
	typeName, ok := ProcessTypeName(typ.LLString(), errorLocation)
	if !ok {
		return symbols.TypeSymbol{}, false
	}

	typeSymbol := ResolveObjectType(typeName, classes, false, errorLocation)
	if typeSymbol != nil {
		return *typeSymbol, true
	}

	// aaaaand if we found nothing -> cry
	print.Error(
		"PACKAGER",
		print.UnknownDataTypeError,
		errorLocation,
		"Could not resolve referenced type '%s' while loading package!",
		typ.LLString(),
	)
	os.Exit(-1)

	return symbols.TypeSymbol{}, false
}

func ResolveObjectType(typeName string, classes []*symbols.ClassSymbol, allowLower bool, errorLocation print.TextSpan) *symbols.TypeSymbol {
	// disallow boxed types
	if typeName == "Byte" || typeName == "Int" || typeName == "Long" || typeName == "Float" || typeName == "Double" || typeName == "Bool" {
		print.Error(
			"PACKAGER",
			print.IllegalBoxedTypeError,
			errorLocation,
			"the use of boxed types (object versions of int, byte, float, bool) is not allowed. If you wish to give back an object of a primitive please cast it to 'any'. (Caused by: %s)",
			typeName,
		)
		os.Exit(-1)
	}

	// string type
	if typeName == "String" || (allowLower && typeName == "string") {
		return &builtins.String
	}

	// thread type
	if typeName == "Thread" || (allowLower && typeName == "thread") {
		return &builtins.Thread
	}

	// thread type
	if typeName == "Any" || (allowLower && typeName == "any") {
		return &builtins.Any
	}

	// array types
	// these arent final types and need to be resolved by fingerprint
	if typeName == "Array" {
		print.Error(
			"PACKAGER",
			print.IllegalUnspecificArrayTypeError,
			errorLocation,
			"Use of unspecific array type is not allowed!",
		)
		os.Exit(-1)
		return &builtins.Array
	}
	if strings.HasPrefix(typeName, "Array_") {
		typ := ResolveArrayType(typeName, false, classes, errorLocation)
		return &typ
	}

	if typeName == "pArray" {
		print.Error(
			"PACKAGER",
			print.IllegalUnspecificArrayTypeError,
			errorLocation,
			"Use of unspecific p-array type is not allowed!",
		)
		os.Exit(-1)
		return &builtins.PArray
	}

	// if its something else, we need to look through our other classes and packages
	for _, cls := range classes {
		if typeName == cls.Name {
			return &cls.Type
		}
	}

	for _, pkg := range PackagesSoFar {
		for _, cls := range pkg.Classes {
			if typeName == cls.Name {
				return &cls.Type
			}
		}
	}

	return nil
}

func ResolveTypeFromName(typeName string, classes []*symbols.ClassSymbol, errorLocation print.TextSpan) symbols.TypeSymbol {
	if typeName == "Bool" || typeName == "bool" {
		return builtins.Bool
	}

	if typeName == "Byte" || typeName == "byte" {
		return builtins.Byte
	}

	if typeName == "Int" || typeName == "int" {
		return builtins.Int
	}

	if typeName == "Long" || typeName == "long" {
		return builtins.Long
	}

	if typeName == "UInt" || typeName == "uint" {
		return builtins.UInt
	}

	if typeName == "ULong" || typeName == "ulong" {
		return builtins.ULong
	}

	if typeName == "Float" || typeName == "float" {
		return builtins.Float
	}

	if typeName == "Double" || typeName == "double" {
		return builtins.Double
	}

	typeSymbol := ResolveObjectType(typeName, classes, true, errorLocation)
	if typeSymbol != nil {
		return *typeSymbol
	}

	// aaaaand if we found nothing -> cry
	print.Error(
		"PACKAGER",
		print.UnknownDataTypeError,
		errorLocation,
		"Could not resolve type '%s' while loading package!",
		typeName,
	)
	os.Exit(-1)
	return symbols.TypeSymbol{}
}

func ProcessTypeName(name string, errorLocation print.TextSpan) (string, bool) {
	// if this type name doesnt match the rect class pattern and also isnt a primitive
	// => no idea what the fuck this is
	if !strings.HasPrefix(name, "%struct.class_") {
		print.Error(
			"PACKAGER",
			print.MonkeError,
			errorLocation,
			"Referenced type '%s' does not match ReCT class pattern! Absolutely no clue what todo with it lol",
			name,
		)
		return "", false
	}

	// if this is a valid class but not a pointer -> hm?????
	// (objects are referential types so they NEED to be pointers)
	if !strings.HasSuffix(name, "*") {
		print.Error(
			"PACKAGER",
			print.InvalidNonPointerReferenceError,
			errorLocation,
			"Referenced object type '%s' needs to be a pointer but isnt!",
			name,
		)
		os.Exit(-1)
	}

	// if all those things are alright we can cut away all the un-needed stuff
	return strings.TrimSuffix(strings.TrimPrefix(name, "%struct.class_"), "*"), true
}

func ResolveArrayType(typeName string, isPrimitive bool, classes []*symbols.ClassSymbol, errorLocation print.TextSpan) symbols.TypeSymbol {

	// choose the correct prefix for this type
	prefix := "Array_"
	symName := "array"
	if isPrimitive {
		prefix = "pArray_"
		symName = "parray"
	}

	// remove the prefix
	baseType := strings.TrimPrefix(typeName, prefix)

	// find out it this is a fingerprint or a type name
	if strings.HasPrefix(baseType, "T_") {
		// fingerprint

		// replace identifier escape sequences with the characters they represent
		baseType = strings.Replace(baseType, "$b$", "[", -1)
		baseType = strings.Replace(baseType, "$e$", "]", -1)
		baseType = strings.Replace(baseType, "$s$", ";", -1)

		// parse the fingerprint
		return symbols.CreateTypeSymbol(symName,
			[]symbols.TypeSymbol{ParseFingerprint(baseType, baseType, classes, errorLocation)},
			true, false)
	}

	base := ResolveTypeFromName(baseType, classes, errorLocation)
	return symbols.CreateTypeSymbol(symName, []symbols.TypeSymbol{base}, true, false)
}

func ParseFingerprint(o, fingerprint string, classes []*symbols.ClassSymbol, errorLocation print.TextSpan) symbols.TypeSymbol {
	fingerprint = strConsume(o, fingerprint, "T_", errorLocation)

	// type name
	fingerprint, name := strReadWord(fingerprint)

	// sub types
	subTypes := make([]symbols.TypeSymbol, 0)

	fingerprint = strConsume(o, fingerprint, "_[", errorLocation)
	for !strCurrent(fingerprint, "]") {
		if strCurrent(fingerprint, "T_") {
			subTypes = append(subTypes, ParseFingerprint(o, fingerprint, classes, errorLocation))
		} else {
			f, typ := strReadWord(fingerprint)
			fingerprint = f
			subTypes = append(subTypes, ResolveTypeFromName(typ, classes, errorLocation))
		}

		fingerprint = strConsume(o, fingerprint, ";", errorLocation)
	}

	return symbols.CreateTypeSymbol(name, subTypes, true, false)
}

func strCurrent(fingerprint string, match string) bool {
	cutout := fingerprint[:len(match)]
	return cutout == match
}

func strConsume(o, fingerprint string, match string, errorLocation print.TextSpan) string {
	cutout := fingerprint[:len(match)]
	if cutout != match {
		print.Error(
			"PACKAGER",
			print.UnparsableFingerprintError,
			errorLocation,
			"error parsing fingerprint for type '%s'!",
			o,
		)
		os.Exit(-1)
	}

	return fingerprint[len(match):]
}

func strReadWord(fingerprint string) (string, string) {
	word := ""
	pointer := 0

	for {
		if unicode.IsLetter(rune(fingerprint[pointer])) || unicode.IsDigit(rune(fingerprint[pointer])) {
			word += string(fingerprint[pointer])
			pointer++
		} else {
			break
		}
	}

	return fingerprint[pointer:], word
}
