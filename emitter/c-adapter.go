package emitter

import (
	"bytes"
	"fmt"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/builtins"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/print"
	"github.com/ReCT-Lang/ReCT-Go-Compiler/symbols"
	"os"
	"os/exec"
)

var cCode string
var AdapterModule string

type ConvertedStruct struct {
	key    string
	code   string
	fields []symbols.VariableSymbol
}

func (emt *Emitter) Adapt() {
	cCode = "" // we no cod

	// we cod
	cCode += "#include <stdbool.h>\n"

	cCode += "\n//classes\n"
	for _, cls := range emt.Classes {
		cCode += "typedef struct " + cls.Name + " " + cls.Name + ";\n"
	}

	cCode += "\n//structs\n"

	structMap := make([]ConvertedStruct, 0)

	for _, stc := range emt.Structs {

		code := "typedef struct " + stc.Name + " { "

		for i, fld := range stc.Symbol.Fields {
			code += fmt.Sprintf("   %s Fld%d;", emt.ConvertType(fld.VarType(), false), i)
		}

		code += " } " + stc.Name + "; \n"

		structMap = append(structMap, ConvertedStruct{
			key:    stc.Symbol.Type.Fingerprint(),
			code:   code,
			fields: stc.Symbol.Fields,
		})
	}

	maxRounds := 1_000_000 // one million loops at max
	rounds := 0

	// while tru (very good idea)
	for {
		// make sure we dont loop forever
		if rounds > maxRounds {
			print.Error(
				"[INTERNAL C-ADAPTER]",
				print.CAdapterCompilationError,
				print.TextSpan{},
				"Error while compiling C-Adapter module! Struct sorting seems to be in an infinite loop.")
			return
		}

		madeChange := false

		for i, convertedStruct := range structMap {
			insertIndex := i

			for _, fld := range convertedStruct.fields {
				typ := fld.VarType()

				// if the type is a pointer -> find its base
				for typ.Name == builtins.Pointer.Name {
					typ = typ.SubTypes[0]
				}

				// make sure this goes in the right order
				if typ.IsUserDefined && !typ.IsObject {
					// look up this struct in the map
					for i2, cstc := range structMap {
						if cstc.key == typ.Fingerprint() {
							if i2 > insertIndex {
								insertIndex = i2 + 1
								break
							}
						}
					}
				}
			}

			if insertIndex > i {
				structMap = insert(structMap, insertIndex, convertedStruct)
				structMap[i] = ConvertedStruct{key: "", code: ""}
				madeChange = true
			}
		}

		if !madeChange {
			break
		}

		rounds++
	}

	for _, convertedStruct := range structMap {
		cCode += convertedStruct.code
	}

	// we adapt
	externalCode := ""
	adapterCode := ""
	actuallyDidSomething := false

	for _, function := range emt.Program.ExternalFunctions {
		if function.Adapted {
			// adapted function
			adapterCode += emt.ConvertType(function.Type, true) + " " + function.Name + "$ADAPTED ( "

			for _, parameter := range function.Parameters {
				adapterCode += emt.ConvertType(parameter.Type, true) + " " + parameter.Name + ","
			}

			adapterCode = adapterCode[:len(adapterCode)-1] // remove last comma
			adapterCode += ") {\n"

			// I S   T H I S   A   V O I D ?
			if function.Type.Fingerprint() == builtins.Void.Fingerprint() {
				adapterCode += "   " + emt.ConvertExternalCall(function) + "\n}\n"

				// I S   T H I S   A   S T R U C T?
			} else if function.Type.IsUserDefined && !function.Type.IsObject {
				adapterCode += "   " + emt.ConvertType(function.Type, false) + " ret;\n" // temp var
				adapterCode += "   ret = " + emt.ConvertExternalCall(function) + "\n"
				adapterCode += "   return &ret;\n}\n"

				// primimimv (yes)
			} else {
				adapterCode += "return " + emt.ConvertExternalCall(function) + ";\n}\n"
			}

			// external function
			externalCode += emt.ConvertType(function.Type, false) + " " + function.Name + " ( "

			for _, parameter := range function.Parameters {
				externalCode += emt.ConvertType(parameter.Type, false) + " " + parameter.Name + ","
			}

			externalCode = externalCode[:len(externalCode)-1] // remove last comma
			externalCode += ");\n"

			actuallyDidSomething = true
		}
	}

	if !actuallyDidSomething {
		return
	}

	cCode += "\n//external funcs\n"
	cCode += externalCode
	cCode += "\n//adapter funcs\n"
	cCode += adapterCode

	// we compile
	cmd := exec.Command("clang", "-x", "c", "-O0", "-S", "-emit-llvm", "-", "-o", "-")

	buffer := bytes.Buffer{}
	buffer.Write([]byte(cCode))

	cmd.Stdin = &buffer

	out, err := cmd.Output()
	if err != nil { //Use start, not run
		fmt.Println(cCode)
		print.Error(
			"[INTERNAL C-ADAPTER]",
			print.CAdapterCompilationError,
			print.TextSpan{},
			"Error while compiling C-Adapter module! (THIS SHOULDNT HAPPEN! Please file a bug report!)")
		return
	}

	//fmt.Println(cCode)
	//fmt.Println(string(out))
	AdapterModule = string(out)
}

func (emt *Emitter) ConvertExternalCall(function symbols.FunctionSymbol) string {
	code := function.Name + " ( "

	// pass all params, make sure to dereference any structs
	for _, parameter := range function.Parameters {
		if parameter.Type.IsUserDefined && !parameter.Type.IsObject {
			code += "*" // deref
		}
		code += parameter.Name + ","
	}

	code = code[:len(code)-1] // remove last comma
	code += ");"

	return code
}

func (emt *Emitter) ConvertType(fld symbols.TypeSymbol, structsAsPointers bool) string {
	switch fld.Fingerprint() {
	case builtins.Void.Fingerprint():
		return "void"
	case builtins.Bool.Fingerprint():
		return "bool"
	case builtins.Byte.Fingerprint():
		return "char"
	case builtins.Int.Fingerprint():
		return "int"
	case builtins.Long.Fingerprint():
		return "long"
	case builtins.UInt.Fingerprint():
		return "int"
	case builtins.ULong.Fingerprint():
		return "long"
	case builtins.Float.Fingerprint():
		return "float"
	case builtins.Double.Fingerprint():
		return "double"
	case builtins.String.Fingerprint():
		return emt.Classes[emt.Id(builtins.String)].Name + "*"
	case builtins.Any.Fingerprint():
		return emt.Classes[emt.Id(builtins.Any)].Name + "*"
	case builtins.Thread.Fingerprint():
		return emt.Classes[emt.Id(builtins.Thread)].Name + "*"
	}

	if fld.Name == builtins.Array.Name {
		if fld.SubTypes[0].IsObject {
			return emt.Classes[emt.Id(builtins.Array)].Name + "*"
		} else {
			return emt.Classes[emt.Id(builtins.PArray)].Name + "*"
		}
	}

	if fld.Name == builtins.Pointer.Name {
		return emt.ConvertType(fld.SubTypes[0], structsAsPointers) + "*"
	}

	// try looking up a class
	cls, ok := emt.Classes[emt.Id(fld)]
	if ok {
		return cls.Name + "*"
	}

	// try looking up a struct
	stc, ok := emt.Structs[emt.Id(fld)]
	if ok {
		if structsAsPointers {
			return stc.Name + "*"
		} else {
			return stc.Name
		}
	}

	fmt.Println("Unknown Type")
	fmt.Println(fld.Fingerprint())

	for _, v := range emt.Classes {
		fmt.Println("> " + v.Name)
	}

	os.Exit(-1) // we crashin
	return ""
}

// theft (https://stackoverflow.com/questions/46128016/insert-a-value-in-a-slice-at-a-given-index)
func insert(a []ConvertedStruct, index int, value ConvertedStruct) []ConvertedStruct {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
