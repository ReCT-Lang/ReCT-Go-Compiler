package irtools

import (
	"ReCT-Go-Compiler/print"
	"fmt"
	"os"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
)

func ReadModule(path string) *ir.Module {
	// read out contents of the the file
	moduleBytes, _ := os.ReadFile(path)
	module := string(moduleBytes)

	// do a regex replace to replace all function content with 'ret void'
	re := regexp2.MustCompile(`{\n.*?^}\n`, regexp2.Singleline|regexp2.Multiline)
	module, _ = re.Replace(module, " {\n  ret void\n}\n", 0, -1)

	// do a regex replace to remove invalid function declarations
	re2 := regexp2.MustCompile(`(?<=declare.*?)align [0-9]*`, regexp2.Multiline)
	module, _ = re2.Replace(module, " ", 0, -1)

	// do a regex replace to remove invalid function declarations
	re3 := regexp2.MustCompile(`(?<=define.*?)align [0-9]*`, regexp2.Multiline)
	module, _ = re3.Replace(module, " ", 0, -1)

	os.WriteFile("./mod.ll", []byte(module), os.ModePerm)

	// do a regex replace to remove new fangled sret
	//re3 := regexp2.MustCompile(`sret\(.*?\)`, regexp2.Multiline)
	//module, _ = re3.Replace(module, " ", 0, -1)

	// parse the module using llir/llvm
	irModule, err := asm.ParseString(path, module)
	if err != nil {
		print.PrintC(print.Red, "Couldnt load module '"+path+"'")
		fmt.Println(module)
		panic(err)
	}

	return irModule
}

func FindFunctionsWithPrefix(module *ir.Module, prefix string) []*ir.Func {
	funcs := make([]*ir.Func, 0)

	for _, fnc := range module.Funcs {
		if strings.HasPrefix(fnc.Name(), prefix) {
			funcs = append(funcs, fnc)
		}
	}

	return funcs
}

func FindFunction(module *ir.Module, name string) *ir.Func {
	for _, fnc := range module.Funcs {
		if fnc.Name() == name {
			return fnc
		}
	}

	print.PrintC(print.Red, "Couldnt find function '"+name+"'")
	return nil
}

func TryFindFunction(module *ir.Module, name string) *ir.Func {
	for _, fnc := range module.Funcs {
		if fnc.Name() == name {
			return fnc
		}
	}

	return nil
}

func FindGlobal(module *ir.Module, name string) *ir.Global {
	for _, glb := range module.Globals {
		if glb.Name() == name {
			return glb
		}
	}

	print.PrintC(print.Red, "Couldnt find global '"+name+"'")
	return nil
}

func FindGlobalSuffix(module *ir.Module, name string) *ir.Global {
	for _, glb := range module.Globals {
		if strings.HasSuffix(glb.Name(), name) {
			return glb
		}
	}

	print.PrintC(print.Red, "Couldnt find global '"+name+"'")
	return nil
}

func TryFindGlobal(module *ir.Module, name string) *ir.Global {
	for _, glb := range module.Globals {
		if glb.Name() == name {
			return glb
		}
	}

	return nil
}

func ReadConstStringArray(module *ir.Module, glb *ir.Global) ([]string, bool) {
	arr, ok := glb.Init.(*constant.Array)
	if !ok {
		return make([]string, 0), false
	}

	names := make([]string, 0)

	for _, elem := range arr.Elems {
		gep, ok := elem.(*constant.ExprGetElementPtr)
		if !ok {
			return make([]string, 0), false
		}

		cFld := TryFindGlobal(module, strings.TrimPrefix(gep.Src.Ident(), "@"))
		if cFld == nil {
			return make([]string, 0), false
		}

		chr, ok := cFld.Init.(*constant.CharArray)
		if !ok {
			return make([]string, 0), false
		}

		names = append(names, string(chr.X))
	}

	return names, true
}
