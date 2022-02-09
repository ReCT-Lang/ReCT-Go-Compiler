package emitter

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/symbols"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Emitter struct {
	Program         binder.BoundProgram
	Module          *ir.Module
	UseFingerprints bool

	// global variables
	Globals map[string]*ir.Global

	// local things for this current function
	Function *ir.Func
	Locals   map[string]*ir.InstAlloca
	Labels   map[string]*ir.Block
}

func Emit(program binder.BoundProgram, useFingerprints bool) *ir.Module {
	emitter := Emitter{
		Program:         program,
		Module:          ir.NewModule(),
		UseFingerprints: useFingerprints,
		Globals:         make(map[string]*ir.Global),
	}

	for _, fnc := range emitter.Program.Functions {
		if !fnc.Symbol.BuiltIn {
			emitter.EmitFunction(fnc.Symbol, fnc.Body)
		}
	}

	return emitter.Module
}

// <FUNCTIONS>-----------------------------------------------------------------
func (emt *Emitter) EmitFunction(sym symbols.FunctionSymbol, body boundnodes.BoundBlockStatementNode) {
	// figure out all parameters and their types
	params := make([]*ir.Param, 0)
	for _, param := range sym.Parameters {
		// figure out how to call this parameter
		paramName := tern(emt.UseFingerprints, param.Fingerprint(), param.Name)

		// create it
		params = append(params, ir.NewParam(paramName, IRTypes[param.Type.Fingerprint()]))
	}

	// figure out the return type
	returnType := IRTypes[sym.Type.Fingerprint()]

	// the function name
	functionName := tern(emt.UseFingerprints, sym.Fingerprint(), sym.Name)

	// create an IR function definition
	function := emt.Module.NewFunc(functionName, returnType, params...)

	// set up our environment
	emt.Function = function
	emt.Locals = make(map[string]*ir.InstAlloca)
	emt.Labels = make(map[string]*ir.Block)

	// yes
	emt.EmitBlockStatement(function, body)
}

func (emt *Emitter) EmitBlockStatement(fnc *ir.Func, body boundnodes.BoundBlockStatementNode) {
	// create a root block
	// each label statement will create a new block and store it in this variable
	currentBlock := fnc.NewBlock("")

	// go through the body and register all label blocks
	for _, stmt := range body.Statements {
		if stmt.NodeType() == boundnodes.BoundLabelStatement {
			// create a new block when we encounter a label
			labelStatement := stmt.(boundnodes.BoundLabelStatementNode)
			emt.Labels[string(labelStatement.Label)] = fnc.NewBlock(string(labelStatement.Label))
		}
	}

	for _, stmt := range body.Statements {
		if stmt.NodeType() == boundnodes.BoundLabelStatement {
			// if we encounter a laber statement -> switch to the label's block
			labelStatement := stmt.(boundnodes.BoundLabelStatementNode)
			currentBlock = emt.Labels[string(labelStatement.Label)]
		} else {
			emt.EmitStatement(currentBlock, stmt)
		}
	}
}

func (emt *Emitter) EmitStatement(blk *ir.Block, stmt boundnodes.BoundStatementNode) {
	// emit a statement to the current block
	switch stmt.NodeType() {
	case boundnodes.BoundVariableDeclaration:
		emt.EmitVariableDeclarationStatement(blk, stmt.(boundnodes.BoundVariableDeclarationStatementNode))

	case boundnodes.BoundGotoStatement:
		emt.EmitGotoStatement(blk, stmt.(boundnodes.BoundGotoStatementNode))

	case boundnodes.BoundConditionalGotoStatement:
		emt.EmitConditionalGotoStatement(blk, stmt.(boundnodes.BoundConditionalGotoStatementNode))

	case boundnodes.BoundExpressionStatement:
		emt.EmitBoundExpressionStatement(blk, stmt.(boundnodes.BoundExpressionStatementNode))

	case boundnodes.BoundReturnStatement:
		emt.EmitReturnStatement(blk, stmt.(boundnodes.BoundReturnStatementNode))
	}
}

// </FUNCTIONS>----------------------------------------------------------------
// <STATEMENTS>----------------------------------------------------------------

func (emt *Emitter) EmitVariableDeclarationStatement(blk *ir.Block, stmt boundnodes.BoundVariableDeclarationStatementNode) {
	varName := tern(emt.UseFingerprints, stmt.Variable.Fingerprint(), stmt.Variable.SymbolName())

	if stmt.Variable.IsGlobal() {
		// create a new global
		global := emt.Module.NewGlobal(varName, IRTypes[stmt.Variable.VarType().Fingerprint()])

		// save it for referencing later
		emt.Globals[varName] = global

		// emit its assignment
		blk.NewStore(emt.EmitExpression(blk, stmt.Initializer), global)
	} else {
		// create local variable
		local := ir.NewAlloca(IRTypes[stmt.Variable.VarType().Fingerprint()])
		local.SetName(varName)

		// save it for referencing later
		emt.Locals[varName] = local

		// emit its assignemnt
		blk.NewStore(emt.EmitExpression(blk, stmt.Initializer), local)
	}
}

func (emt *Emitter) EmitGotoStatement(blk *ir.Block, stmt boundnodes.BoundGotoStatementNode) {
	blk.NewBr(emt.Labels[string(stmt.Label)])
}

func (emt *Emitter) EmitConditionalGotoStatement(blk *ir.Block, stmt boundnodes.BoundConditionalGotoStatementNode) {
	// figure out where to jump
	ifLabel := emt.Labels[string(stmt.IfLabel)]
	elseLabel := emt.Labels[string(stmt.ElseLabel)]

	blk.NewCondBr(emt.EmitExpression(blk, stmt.Condition), ifLabel, elseLabel)
}

func (emt *Emitter) EmitBoundExpressionStatement(blk *ir.Block, stmt boundnodes.BoundExpressionStatementNode) {
	// dummy local because i cant have expression be on their own
	dummyLocal := ir.NewAlloca(IRTypes[stmt.Expression.Type().Fingerprint()])
	expression := emt.EmitExpression(blk, stmt.Expression)
	blk.NewStore(expression, dummyLocal)
}

func (emt *Emitter) EmitReturnStatement(blk *ir.Block, stmt boundnodes.BoundReturnStatementNode) {
	blk.NewRet(nil)
}

// </STATEMENTS>---------------------------------------------------------------
// <EXPRESSIONS>---------------------------------------------------------------

func (emt *Emitter) EmitExpression(blk *ir.Block, expr boundnodes.BoundExpressionNode) value.Value {
	return constant.NewInt(IRTypes[builtins.Int.Fingerprint()].(*types.IntType), 0)
}

// </EXPRESSIONS>--------------------------------------------------------------

// yes
func tern(cond bool, str1 string, str2 string) string {
	if cond {
		return str1
	} else {
		return str2
	}
}

func btern(cond bool, v1 *ir.Block, v2 *ir.Block) *ir.Block {
	if cond {
		return v1
	} else {
		return v2
	}
}

func itern(cond bool, v1 interface{}, v2 interface{}) interface{} {
	if cond {
		return v1
	} else {
		return v2
	}
}
