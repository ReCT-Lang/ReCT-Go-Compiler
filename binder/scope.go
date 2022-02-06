package binder

import "ReCT-Go-Compiler/symbols"

type Scope struct {
	Parent  *Scope
	Symbols map[string]symbols.Symbol
}

func (s *Scope) TryDeclareSymbol(sym symbols.Symbol) bool {

	_, found := s.Symbols[sym.SymbolName()]

	if found {
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

// constructor
func CreateScope(parent *Scope) Scope {
	return Scope{
		Parent:  parent,
		Symbols: make(map[string]symbols.Symbol),
	}
}
