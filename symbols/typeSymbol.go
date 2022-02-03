package symbols

type TypeSymbol struct {
	Symbol

	Name     string
	SubTypes []TypeSymbol
}

// implement the interface
func (TypeSymbol) SymbolType() SymbolType { return Type }
func (s TypeSymbol) SymbolName() string   { return s.Name }

// a unique identifier for each type
func (sym TypeSymbol) Fingerprint() string {
	id := "T_" + sym.Name + "_"

	for _, subtype := range sym.SubTypes {
		id += subtype.Name + ";"
	}

	return id
}

// constructor
func CreateTypeSymbol(name string, subTypes []TypeSymbol) TypeSymbol {
	return TypeSymbol{
		Name:     name,
		SubTypes: subTypes,
	}
}
