package binder

import "ReCT-Go-Compiler/symbols"

type Scope struct {
	Parent  *Scope
	Symbols map[string]symbols.Symbol
}

func (s *Scope) TryDeclareSymbol(sym symbols.Symbol) bool {

	lookup := s.TryLookupSymbol(sym.SymbolName())

	if lookup != nil {
		return false // symbol already exists
	} else {
		s.Symbols[sym.SymbolName()] = sym
		return true
	}
}

func (s Scope) TryLookupSymbol(name string) symbols.Symbol {
	sym, found := s.Symbols[name]

	if found {
		return sym
	}

	if s.Parent != nil {
		return s.Parent.TryLookupSymbol(name)
	}

	return nil
}

func (s Scope) InsertFunctionSymbols(symbols []symbols.FunctionSymbol) {
	for _, sym := range symbols {
		s.TryDeclareSymbol(sym)
	}
}

func (s Scope) InsertVariableSymbols(symbols []symbols.VariableSymbol) {
	for _, sym := range symbols {
		s.TryDeclareSymbol(sym)
	}
}

func (s Scope) GetAllFunctions() []symbols.FunctionSymbol {
	functions := make([]symbols.FunctionSymbol, 0)

	for _, sym := range s.Symbols {
		if sym.SymbolType() == symbols.Function {
			functions = append(functions, sym.(symbols.FunctionSymbol))
		}
	}

	moreFunctions := make([]symbols.FunctionSymbol, 0)
	if s.Parent != nil {
		moreFunctions = s.Parent.GetAllFunctions()
	}

	functions = append(functions, moreFunctions...)

	return functions
}

func (s Scope) GetAllVariables() []symbols.VariableSymbol {
	variables := make([]symbols.VariableSymbol, 0)

	for _, sym := range s.Symbols {
		if sym.SymbolType() == symbols.LocalVariable ||
			sym.SymbolType() == symbols.GlobalVariable ||
			sym.SymbolType() == symbols.Parameter {
			variables = append(variables, sym.(symbols.VariableSymbol))
		}
	}

	moreVariables := make([]symbols.VariableSymbol, 0)
	if s.Parent != nil {
		moreVariables = s.Parent.GetAllVariables()
	}

	variables = append(variables, moreVariables...)

	return variables
}

func (s Scope) GetAllClasses() []symbols.ClassSymbol {
	classes := make([]symbols.ClassSymbol, 0)

	for _, sym := range s.Symbols {
		if sym.SymbolType() == symbols.Class {
			classes = append(classes, sym.(symbols.ClassSymbol))
		}
	}

	moreClasses := make([]symbols.ClassSymbol, 0)
	if s.Parent != nil {
		moreClasses = s.Parent.GetAllClasses()
	}

	classes = append(classes, moreClasses...)

	return classes
}

func (s Scope) GetAllStructs() []symbols.StructSymbol {
	structs := make([]symbols.StructSymbol, 0)

	for _, sym := range s.Symbols {
		if sym.SymbolType() == symbols.Struct {
			structs = append(structs, sym.(symbols.StructSymbol))
		}
	}

	moreStructs := make([]symbols.StructSymbol, 0)
	if s.Parent != nil {
		moreStructs = s.Parent.GetAllStructs()
	}

	structs = append(structs, moreStructs...)

	return structs
}

func (s Scope) GetAllPackages() []symbols.PackageSymbol {
	packages := make([]symbols.PackageSymbol, 0)

	for _, sym := range s.Symbols {
		if sym.SymbolType() == symbols.Package {
			packages = append(packages, sym.(symbols.PackageSymbol))
		}
	}

	morePackages := make([]symbols.PackageSymbol, 0)
	if s.Parent != nil {
		morePackages = s.Parent.GetAllPackages()
	}

	packages = append(packages, morePackages...)

	return packages
}

// constructor
func CreateScope(parent *Scope) Scope {
	return Scope{
		Parent:  parent,
		Symbols: make(map[string]symbols.Symbol),
	}
}
