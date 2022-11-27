package langserverinterface

import (
	"ReCT-Go-Compiler/lexer"
	"ReCT-Go-Compiler/nodes"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"os"
)

var labelCounter int = 0

func Map(functionSymbol symbols.FunctionSymbol, stmt boundnodes.BoundStatementNode) {
	if TokenMapping == nil {
		TokenMapping = make(map[lexer.Token]TokenMeaning)
	}

	MapStatement(stmt)
}

func MapStatement(stmt boundnodes.BoundStatementNode) {
	switch stmt.NodeType() {
	case boundnodes.BoundBlockStatement:
		MapBlockStatement(stmt.(boundnodes.BoundBlockStatementNode))
	case boundnodes.BoundVariableDeclaration:
		MapVariableDeclaration(stmt.(boundnodes.BoundVariableDeclarationStatementNode))
	case boundnodes.BoundIfStatement:
		MapIfStatement(stmt.(boundnodes.BoundIfStatementNode))
	case boundnodes.BoundWhileStatement:
		MapWhileStatement(stmt.(boundnodes.BoundWhileStatementNode))
	case boundnodes.BoundForStatement:
		MapForStatement(stmt.(boundnodes.BoundForStatementNode))
	case boundnodes.BoundFromToStatement:
		MapFromToStatement(stmt.(boundnodes.BoundFromToStatementNode))
	case boundnodes.BoundLabelStatement:
		MapLabelStatement(stmt.(boundnodes.BoundLabelStatementNode))
	case boundnodes.BoundGotoStatement:
		MapGotoStatement(stmt.(boundnodes.BoundGotoStatementNode))
	case boundnodes.BoundConditionalGotoStatement:
		MapConditionalGotoStatement(stmt.(boundnodes.BoundConditionalGotoStatementNode))
	case boundnodes.BoundReturnStatement:
		MapReturnStatement(stmt.(boundnodes.BoundReturnStatementNode))
	case boundnodes.BoundExpressionStatement:
		MapExpressionStatement(stmt.(boundnodes.BoundExpressionStatementNode))
	default:
		print.PrintC(print.Red, "Statement unaccounted for in mapper! (stuff being in here is important for the language server lol)")
		os.Exit(-1)
	}
}

func MapBlockStatement(stmt boundnodes.BoundBlockStatementNode) {
	for _, statement := range stmt.Statements {
		MapStatement(statement)
	}
}

func MapVariableDeclaration(stmt boundnodes.BoundVariableDeclarationStatementNode) {
	if stmt.Initializer != nil {
		MapExpression(stmt.Initializer)
	}
}

func MapIfStatement(stmt boundnodes.BoundIfStatementNode) {
	MapExpression(stmt.Condition)
	MapStatement(stmt.ThenStatement)

	if stmt.ElseStatement != nil {
		MapStatement(stmt.ElseStatement)
	}
}

func MapWhileStatement(stmt boundnodes.BoundWhileStatementNode) {
	MapExpression(stmt.Condition)
	MapStatement(stmt.Body)
}

func MapForStatement(stmt boundnodes.BoundForStatementNode) {
	MapExpression(stmt.Condition)
	MapStatement(stmt.Variable)
	MapStatement(stmt.Action)
	MapStatement(stmt.Body)
}

func MapFromToStatement(stmt boundnodes.BoundFromToStatementNode) {
	MapExpression(stmt.LowerBound)
	MapExpression(stmt.UpperBound)
	MapStatement(stmt.Body)
}

// -----------------------------------------------------------------------
// no idea why these are even declared (they shouldnt actually exist here)
// -----------------------------------------------------------------------

func MapLabelStatement(stmt boundnodes.BoundLabelStatementNode) {
}

func MapGotoStatement(stmt boundnodes.BoundGotoStatementNode) {
}

func MapConditionalGotoStatement(stmt boundnodes.BoundConditionalGotoStatementNode) {
}

// -----------------------------------------------------------------------

func MapReturnStatement(stmt boundnodes.BoundReturnStatementNode) {
	if stmt.Expression != nil {
		MapExpression(stmt.Expression)
	}
}

func MapExpressionStatement(stmt boundnodes.BoundExpressionStatementNode) {
	MapExpression(stmt.Expression)
}

func MapExpression(expr boundnodes.BoundExpressionNode) {
	switch expr.NodeType() {
	case boundnodes.BoundErrorExpression:
		MapErrorExpression(expr.(boundnodes.BoundErrorExpressionNode))
	case boundnodes.BoundLiteralExpression:
		MapLiteralExpression(expr.(boundnodes.BoundLiteralExpressionNode))
	case boundnodes.BoundVariableExpression:
		MapVariableExpression(expr.(boundnodes.BoundVariableExpressionNode))
	case boundnodes.BoundAssignmentExpression:
		MapAssignmentExpression(expr.(boundnodes.BoundAssignmentExpressionNode))
	case boundnodes.BoundUnaryExpression:
		MapUnaryExpression(expr.(boundnodes.BoundUnaryExpressionNode))
	case boundnodes.BoundBinaryExpression:
		MapBinaryExpression(expr.(boundnodes.BoundBinaryExpressionNode))
	case boundnodes.BoundCallExpression:
		MapCallExpression(expr.(boundnodes.BoundCallExpressionNode))
	case boundnodes.BoundPackageCallExpression:
		MapPackageCallExpression(expr.(boundnodes.BoundPackageCallExpressionNode))
	case boundnodes.BoundConversionExpression:
		MapConversionExpression(expr.(boundnodes.BoundConversionExpressionNode))
	case boundnodes.BoundTypeCallExpression:
		MapTypeCallExpression(expr.(boundnodes.BoundTypeCallExpressionNode))
	case boundnodes.BoundClassCallExpression:
		MapClassCallExpression(expr.(boundnodes.BoundClassCallExpressionNode))
	case boundnodes.BoundClassFieldAccessExpression:
		MapClassFieldAccessExpression(expr.(boundnodes.BoundClassFieldAccessExpressionNode))
	case boundnodes.BoundClassFieldAssignmentExpression:
		MapClassFieldAssignmentExpression(expr.(boundnodes.BoundClassFieldAssignmentExpressionNode))
	case boundnodes.BoundArrayAccessExpression:
		MapArrayAccessExpression(expr.(boundnodes.BoundArrayAccessExpressionNode))
	case boundnodes.BoundArrayAssignmentExpression:
		MapArrayAssignmentExpression(expr.(boundnodes.BoundArrayAssignmentExpressionNode))
	case boundnodes.BoundMakeExpression:
		MapMakeExpression(expr.(boundnodes.BoundMakeExpressionNode))
	case boundnodes.BoundMakeArrayExpression:
		MapMakeArrayExpression(expr.(boundnodes.BoundMakeArrayExpressionNode))
	case boundnodes.BoundMakeStructExpression:
		MapMakeStructExpression(expr.(boundnodes.BoundMakeStructExpressionNode))
	case boundnodes.BoundFunctionExpression:
		MapFunctionExpression(expr.(boundnodes.BoundFunctionExpressionNode))
	case boundnodes.BoundTernaryExpression:
		MapTernaryExpression(expr.(boundnodes.BoundTernaryExpressionNode))
	case boundnodes.BoundReferenceExpression:
		MapReferenceExpression(expr.(boundnodes.BoundReferenceExpressionNode))
	case boundnodes.BoundDereferenceExpression:
		MapDereferenceExpression(expr.(boundnodes.BoundDereferenceExpressionNode))
	case boundnodes.BoundLambdaExpression:
		MapLambdaExpression(expr.(boundnodes.BoundLambdaExpressionNode))
	case boundnodes.BoundThisExpression:
		MapThisExpression(expr.(boundnodes.BoundThisExpressionNode))
	case boundnodes.BoundEnumExpression:
		MapEnumExpression(expr.(boundnodes.BoundEnumExpressionNode))
	default:
		print.PrintC(print.Red, "Expression unaccounted for in mappr! (stuff being in here is important for the language server lol)")
		os.Exit(-1)
	}
}

func MapErrorExpression(expr boundnodes.BoundErrorExpressionNode) {
}

func MapLiteralExpression(expr boundnodes.BoundLiteralExpressionNode) {
}

func MapVariableExpression(expr boundnodes.BoundVariableExpressionNode) {
	TokenMapping[expr.Source().(nodes.NameExpressionNode).Identifier] = VariableTokenMeaning{Variable: expr.Variable}
}

func MapAssignmentExpression(expr boundnodes.BoundAssignmentExpressionNode) {
	if expr.Source().NodeType() == nodes.AssignmentExpression {
		TokenMapping[expr.Source().(nodes.AssignmentExpressionNode).Identifier] = VariableTokenMeaning{Variable: expr.Variable}
		MapExpression(expr.Expression)
	}
}

func MapUnaryExpression(expr boundnodes.BoundUnaryExpressionNode) {
	MapExpression(expr.Expression)
}

func MapBinaryExpression(expr boundnodes.BoundBinaryExpressionNode) {
	MapExpression(expr.Left)
	MapExpression(expr.Right)
}

func MapCallExpression(expr boundnodes.BoundCallExpressionNode) {
	TokenMapping[expr.Source().(nodes.CallExpressionNode).Identifier] = FunctionTokenMeaning{Function: expr.Function}
	for _, arg := range expr.Arguments {
		MapExpression(arg)
	}
}

func MapPackageCallExpression(expr boundnodes.BoundPackageCallExpressionNode) {
	TokenMapping[expr.Source().(nodes.PackageCallExpressionNode).Package] = PackageTokenMeaning{Package: expr.Package}
	TokenMapping[expr.Source().(nodes.PackageCallExpressionNode).Identifier] = FunctionTokenMeaning{Function: expr.Function}
	for _, arg := range expr.Arguments {
		MapExpression(arg)
	}
}

func MapConversionExpression(expr boundnodes.BoundConversionExpressionNode) {

	if expr.Source().NodeType() == nodes.CallExpression {
		// classes
		if expr.ToType.IsObject && expr.ToType.IsUserDefined {
			TokenMapping[expr.Source().(nodes.CallExpressionNode).Identifier] = ClassTokenMeaning{Class: expr.ToType.SourceSymbol.(symbols.ClassSymbol)}

			// structs
		} else if !expr.ToType.IsObject && expr.ToType.IsUserDefined {
			TokenMapping[expr.Source().(nodes.CallExpressionNode).Identifier] = StructTokenMeaning{Struct: expr.ToType.SourceSymbol.(symbols.StructSymbol)}

			// enums
		} else if expr.ToType.IsEnum {
			TokenMapping[expr.Source().(nodes.CallExpressionNode).Identifier] = EnumTokenMeaning{Enum: expr.ToType.SourceSymbol.(symbols.EnumSymbol)}

			// complex type (type with subtypes)
		} else if expr.Source().(nodes.CallExpressionNode).CastingType.ClauseIsSet {
			MapComplexType(expr.Source().(nodes.CallExpressionNode).CastingType, expr.ToType)

			// simple type
		} else {
			TokenMapping[expr.Source().(nodes.CallExpressionNode).Identifier] = TypeTokenMeaning{TypeSym: expr.ToType}
		}
	} else if expr.Source().NodeType() == nodes.PackageCallExpression {
		// packages can only contain classes
		TokenMapping[expr.Source().(nodes.PackageCallExpressionNode).Package] = PackageTokenMeaning{Package: expr.ToType.Package}
		TokenMapping[expr.Source().(nodes.PackageCallExpressionNode).Identifier] = ClassTokenMeaning{Class: expr.ToType.SourceSymbol.(symbols.ClassSymbol)}
	}

	MapExpression(expr.Expression)
}

func MapTypeCallExpression(expr boundnodes.BoundTypeCallExpressionNode) {
	MapExpression(expr.Base)

	for _, arg := range expr.Arguments {
		MapExpression(arg)
	}

	TokenMapping[expr.Source().(nodes.TypeCallExpressionNode).CallIdentifier] = TypeFunctionTokenMeaning{TypeFunction: expr.Function}
}

func MapClassCallExpression(expr boundnodes.BoundClassCallExpressionNode) {
	MapExpression(expr.Base)

	for _, arg := range expr.Arguments {
		MapExpression(arg)
	}

	TokenMapping[expr.Source().(nodes.TypeCallExpressionNode).CallIdentifier] = FunctionTokenMeaning{Function: expr.Function}
}

func MapClassFieldAccessExpression(expr boundnodes.BoundClassFieldAccessExpressionNode) {
	MapExpression(expr.Base)
	TokenMapping[expr.Source().(nodes.ClassFieldAccessExpressionNode).FieldIdentifier] = VariableTokenMeaning{Variable: expr.Field}
}

func MapClassFieldAssignmentExpression(expr boundnodes.BoundClassFieldAssignmentExpressionNode) {
	MapExpression(expr.Base)
	MapExpression(expr.Value)

	TokenMapping[expr.Source().(nodes.ClassFieldAccessExpressionNode).FieldIdentifier] = VariableTokenMeaning{Variable: expr.Field}
}

func MapArrayAccessExpression(expr boundnodes.BoundArrayAccessExpressionNode) {
	MapExpression(expr.Base)
	MapExpression(expr.Index)
}

func MapArrayAssignmentExpression(expr boundnodes.BoundArrayAssignmentExpressionNode) {
	MapExpression(expr.Base)
	MapExpression(expr.Index)
	MapExpression(expr.Value)
}

func MapMakeExpression(expr boundnodes.BoundMakeExpressionNode) {
	for _, arg := range expr.Arguments {
		MapExpression(arg)
	}

	if expr.Source().(nodes.MakeExpressionNode).Package != nil {
		TokenMapping[*expr.Source().(nodes.MakeExpressionNode).Package] = PackageTokenMeaning{Package: expr.BaseType.Type.Package}
	}

	TokenMapping[expr.Source().(nodes.MakeExpressionNode).BaseType] = ClassTokenMeaning{Class: expr.BaseType}
}

func MapMakeArrayExpression(expr boundnodes.BoundMakeArrayExpressionNode) {
	MapGenericType(expr.Source().(nodes.MakeArrayExpressionNode).Type, expr.BaseType)

	if expr.IsLiteral {
		for _, literal := range expr.Literals {
			MapExpression(literal)
		}
		return
	}

	MapExpression(expr.Length)
}

func MapMakeStructExpression(expr boundnodes.BoundMakeStructExpressionNode) {
	for _, literal := range expr.Literals {
		MapExpression(literal)
	}

	TokenMapping[expr.Source().(nodes.MakeStructExpressionNode).Type] = StructTokenMeaning{Struct: expr.StructType.SourceSymbol.(symbols.StructSymbol)}
}

func MapFunctionExpression(expr boundnodes.BoundFunctionExpressionNode) {
	TokenMapping[expr.Source().(nodes.NameExpressionNode).Identifier] = FunctionTokenMeaning{Function: expr.Function}
}

func MapTernaryExpression(expr boundnodes.BoundTernaryExpressionNode) {
	MapExpression(expr.Condition)
	MapExpression(expr.If)
	MapExpression(expr.Else)
}

func MapReferenceExpression(expr boundnodes.BoundReferenceExpressionNode) {
	MapExpression(expr.Expression)
}

func MapDereferenceExpression(expr boundnodes.BoundDereferenceExpressionNode) {
	MapExpression(expr.Expression)
}

func MapLambdaExpression(expr boundnodes.BoundLambdaExpressionNode) {
	MapStatement(expr.Body)
}

func MapThisExpression(expr boundnodes.BoundThisExpressionNode) {
	// nothing to do here right now ('this' is a keyword)
}

func MapEnumExpression(expr boundnodes.BoundEnumExpressionNode) {
	TokenMapping[expr.Source().(nodes.ClassFieldAccessExpressionNode).Base.(nodes.NameExpressionNode).Identifier] = EnumTokenMeaning{Enum: expr.Enum}
	TokenMapping[expr.Source().(nodes.ClassFieldAccessExpressionNode).FieldIdentifier] = EnumFieldTokenMeaning{Value: expr.Value}
}

func MapComplexType(clause nodes.TypeClauseNode, typ symbols.TypeSymbol) {
	// map my type
	TokenMapping[clause.TypeIdentifier] = TypeTokenMeaning{TypeSym: typ}

	// map any subtypes
	for i, subType := range typ.SubTypes {
		MapGenericType(clause.SubClauses[i], subType)
	}
}

func MapGenericType(clause nodes.TypeClauseNode, typ symbols.TypeSymbol) {
	// this is a class from a package
	if clause.Package != nil {
		TokenMapping[*clause.Package] = PackageTokenMeaning{Package: typ.Package}
		TokenMapping[clause.TypeIdentifier] = ClassTokenMeaning{Class: typ.SourceSymbol.(symbols.ClassSymbol)}
		return
	}

	// this is from  s o m e w h e r e
	if typ.IsObject && typ.IsUserDefined {
		TokenMapping[clause.TypeIdentifier] = ClassTokenMeaning{Class: typ.SourceSymbol.(symbols.ClassSymbol)}

		// structs
	} else if !typ.IsObject && typ.IsUserDefined {
		TokenMapping[clause.TypeIdentifier] = StructTokenMeaning{Struct: typ.SourceSymbol.(symbols.StructSymbol)}

		// enums
	} else if typ.IsEnum {
		TokenMapping[clause.TypeIdentifier] = EnumTokenMeaning{Enum: typ.SourceSymbol.(symbols.EnumSymbol)}

		// complex type (type with subtypes)
	} else if len(clause.SubClauses) > 0 {
		MapComplexType(clause, typ)

		// simple type
	} else {
		TokenMapping[clause.TypeIdentifier] = TypeTokenMeaning{TypeSym: typ}
	}
}
