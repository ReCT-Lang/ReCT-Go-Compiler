package evaluator

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Evaluator struct {
	Program   binder.BoundProgram
	Globals   map[string]interface{}
	Functions map[string]binder.BoundFunction
	Locals    []map[string]interface{}
}

// locals stack helpers
func (evl *Evaluator) PushLocals() {
	evl.Locals = append(evl.Locals, make(map[string]interface{}))
}

func (evl *Evaluator) PushTheseLocals(locals map[string]interface{}) {
	evl.Locals = append(evl.Locals, locals)
}

func (evl *Evaluator) PopLocals() {
	evl.Locals = evl.Locals[:len(evl.Locals)-1]
}

func (evl *Evaluator) GetLocals() map[string]interface{} {
	return evl.Locals[len(evl.Locals)-1]
}

func (evl *Evaluator) SetLocal(fingerprint string, value interface{}) {
	evl.Locals[len(evl.Locals)-1][fingerprint] = value
}

// variable helper
func (evl *Evaluator) Assign(sym symbols.VariableSymbol, value interface{}) {
	if sym.IsGlobal() {
		evl.Globals[sym.GetFingerprint()] = value
	} else {
		evl.SetLocal(sym.GetFingerprint(), value)
	}
}

var reader *bufio.Reader

// evaluate!
func Evaluate(program binder.BoundProgram) {
	evaluator := Evaluator{
		Program:   program,
		Globals:   make(map[string]interface{}),
		Functions: make(map[string]binder.BoundFunction),
		Locals:    []map[string]interface{}{},
	}

	// setup things
	reader = bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	evaluator.PushLocals()

	for _, fnc := range program.Functions {
		symbol := fnc.Symbol
		evaluator.Functions[symbol.GetFingerprint()] = fnc
	}

	mainFunction := program.MainFunction
	body := evaluator.Functions[mainFunction.GetFingerprint()].Body
	evaluator.EvaluateStatement(body)
}

func (evl *Evaluator) EvaluateStatement(body boundnodes.BoundBlockStatementNode) interface{} {
	labelIndexes := make(map[boundnodes.BoundLabel]int)

	// look through the entire body and look for any labels
	for i, stmt := range body.Statements {
		// if its a label statement...
		if stmt.NodeType() == boundnodes.BoundLabelStatement {
			// register it!
			labelIndexes[stmt.(boundnodes.BoundLabelStatementNode).Label] = i
		}
	}

	index := 0

	for index < len(body.Statements) {
		stmt := body.Statements[index]

		switch stmt.NodeType() {
		case boundnodes.BoundVariableDeclaration:
			evl.EvaluateVariableDeclaration(stmt.(boundnodes.BoundVariableDeclarationStatementNode))
			index++

		case boundnodes.BoundExpressionStatement:
			evl.EvaluateExpressionStatement(stmt.(boundnodes.BoundExpressionStatementNode))
			index++

		case boundnodes.BoundGotoStatement:
			gotoStatement := stmt.(boundnodes.BoundGotoStatementNode)
			index = labelIndexes[gotoStatement.Label]

		case boundnodes.BoundConditionalGotoStatement:
			gotoStatement := stmt.(boundnodes.BoundConditionalGotoStatementNode)
			condition := evl.EvaluateExpression(gotoStatement.Condition)

			if condition == gotoStatement.JumpIfTrue {
				index = labelIndexes[gotoStatement.Label]
			} else {
				index++
			}

		case boundnodes.BoundLabelStatement:
			index++

		case boundnodes.BoundReturnStatement:
			returnStatement := stmt.(boundnodes.BoundReturnStatementNode)

			if returnStatement.Expression != nil {
				return evl.EvaluateExpression(returnStatement.Expression)
			}

			return nil
		}
	}

	return nil
}

func (evl *Evaluator) EvaluateVariableDeclaration(stmt boundnodes.BoundVariableDeclarationStatementNode) {
	value := evl.EvaluateExpression(stmt.Initializer)
	evl.Assign(stmt.Variable, value)
}

func (evl *Evaluator) EvaluateExpressionStatement(stmt boundnodes.BoundExpressionStatementNode) {
	evl.EvaluateExpression(stmt.Expression)
}

// all them expressionz
func (evl *Evaluator) EvaluateExpression(expr boundnodes.BoundExpressionNode) interface{} {
	switch expr.NodeType() {
	case boundnodes.BoundLiteralExpression:
		return evl.EvaluateLiteralExpression(expr.(boundnodes.BoundLiteralExpressionNode))

	case boundnodes.BoundVariableExpression:
		return evl.EvaluateVariableExpression(expr.(boundnodes.BoundVariableExpressionNode))

	case boundnodes.BoundAssignmentExpression:
		return evl.EvaluateAssignmentExpression(expr.(boundnodes.BoundAssignmentExpressionNode))

	case boundnodes.BoundUnaryExpression:
		return evl.EvaluateUnaryExpression(expr.(boundnodes.BoundUnaryExpressionNode))

	case boundnodes.BoundBinaryExpression:
		return evl.EvaluateBinaryExpression(expr.(boundnodes.BoundBinaryExpressionNode))

	case boundnodes.BoundCallExpression:
		return evl.EvaluateCallExpression(expr.(boundnodes.BoundCallExpressionNode))

	case boundnodes.BoundConversionExpression:
		return evl.EvaluateConversionExpression(expr.(boundnodes.BoundConversionExpressionNode))
	}

	return nil
}

func (evl *Evaluator) EvaluateLiteralExpression(expr boundnodes.BoundLiteralExpressionNode) interface{} {
	return expr.Value
}

func (evl *Evaluator) EvaluateVariableExpression(expr boundnodes.BoundVariableExpressionNode) interface{} {
	if expr.Variable.IsGlobal() {
		return evl.Globals[expr.Variable.GetFingerprint()]
	} else {
		locals := evl.GetLocals()
		return locals[expr.Variable.GetFingerprint()]
	}
}

func (evl *Evaluator) EvaluateAssignmentExpression(expr boundnodes.BoundAssignmentExpressionNode) interface{} {
	value := evl.EvaluateExpression(expr.Expression)
	evl.Assign(expr.Variable, value)
	return value
}

func (evl *Evaluator) EvaluateUnaryExpression(expr boundnodes.BoundUnaryExpressionNode) interface{} {
	value := evl.EvaluateExpression(expr.Expression)
	switch expr.Op.OperatorKind {

	case boundnodes.Identity:
		if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return value.(int)
		} else if expr.Expression.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return value.(float32)
		}

	case boundnodes.Negation:
		if expr.Expression.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return -(value.(int))
		} else if expr.Expression.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return -(value.(float32))
		}

	case boundnodes.LogicalNegation:
		return !(value.(bool))

	default:
		print.PrintC(print.Red, "Unknown unary operation!")
		os.Exit(-1)
	}

	return nil
}

func (evl *Evaluator) EvaluateBinaryExpression(expr boundnodes.BoundBinaryExpressionNode) interface{} {
	left := evl.EvaluateExpression(expr.Left)
	right := evl.EvaluateExpression(expr.Right)

	switch expr.Op.OperatorKind {
	case boundnodes.Addition:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return left.(int) + right.(int)
		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return left.(float32) + right.(float32)
		} else if expr.Left.Type().Fingerprint() == builtins.String.Fingerprint() {
			return left.(string) + right.(string)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.Subtraction:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return left.(int) - right.(int)
		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return left.(float32) - right.(float32)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.Multiplication:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return left.(int) * right.(int)
		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return left.(float32) * right.(float32)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.Division:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			if right.(int) == 0 {
				print.PrintC(print.Red, "Division by 0 is illegal!")
				os.Exit(-1)
			}
			return left.(int) / right.(int)
		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			if right.(float32) == 0 {
				print.PrintC(print.Red, "Division by 0 is illegal!")
				os.Exit(-1)
			}
			return left.(float32) / right.(float32)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.Modulus:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			if right.(int) == 0 {
				print.PrintC(print.Red, "Division by 0 is illegal!")
				os.Exit(-1)
			}
			return left.(int) % right.(int)
		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			if right.(float32) == 0 {
				print.PrintC(print.Red, "Division by 0 is illegal!")
				os.Exit(-1)
			}
			return math.Mod(left.(float64), right.(float64))
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.BitwiseAnd:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return left.(int) & right.(int)
		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return left.(bool) && right.(bool)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.BitwiseOr:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return left.(int) | right.(int)
		} else if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return left.(bool) || right.(bool)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.BitwiseXor:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return left.(int) ^ right.(int)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.Equals:
		return left == right

	case boundnodes.NotEquals:
		return left != right

	case boundnodes.Greater:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return left.(int) > right.(int)
		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return left.(float32) > right.(float32)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.GreaterOrEquals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return left.(int) >= right.(int)
		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return left.(float32) >= right.(float32)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.Less:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return left.(int) < right.(int)
		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return left.(float32) < right.(float32)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.LessOrEquals:
		if expr.Left.Type().Fingerprint() == builtins.Int.Fingerprint() {
			return left.(int) <= right.(int)
		} else if expr.Left.Type().Fingerprint() == builtins.Float.Fingerprint() {
			return left.(float32) <= right.(float32)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.LogicalAnd:
		if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return left.(bool) && right.(bool)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	case boundnodes.LogicalOr:
		if expr.Left.Type().Fingerprint() == builtins.Bool.Fingerprint() {
			return left.(bool) || right.(bool)
		}
		print.PrintC(print.Red, "How did this even happen..? (invalid type on operator evaluation)")
		os.Exit(-1)

	}

	return left
}

func (evl *Evaluator) EvaluateCallExpression(expr boundnodes.BoundCallExpressionNode) interface{} {
	// Built in functions
	if expr.Function.GetFingerprint() == builtins.Print.GetFingerprint() {
		text := evl.EvaluateExpression(expr.Arguments[0])
		fmt.Println(text)
		return nil

	} else if expr.Function.GetFingerprint() == builtins.Write.GetFingerprint() {
		text := evl.EvaluateExpression(expr.Arguments[0])
		fmt.Print(text)
		return nil

	} else if expr.Function.GetFingerprint() == builtins.Input.GetFingerprint() {
		text, _ := reader.ReadString('\n')
		return text

	} else if expr.Function.GetFingerprint() == builtins.InputKey.GetFingerprint() {
		char, _ := reader.ReadByte()
		return string(char)

	} else if expr.Function.GetFingerprint() == builtins.Clear.GetFingerprint() {
		fmt.Print("\033[2J")            // clear screen
		fmt.Printf("\033[%d;%dH", 0, 0) // set cursor
		return nil

	} else if expr.Function.GetFingerprint() == builtins.SetCursor.GetFingerprint() {
		x := evl.EvaluateExpression(expr.Arguments[0])
		y := evl.EvaluateExpression(expr.Arguments[1])
		fmt.Printf("\033[%d;%dH", x, y) // set cursor
		return nil

	} else if expr.Function.GetFingerprint() == builtins.Random.GetFingerprint() {
		max := evl.EvaluateExpression(expr.Arguments[0])
		return rand.Intn(max.(int))

	} else if expr.Function.GetFingerprint() == builtins.Sleep.GetFingerprint() {
		mills := evl.EvaluateExpression(expr.Arguments[0])
		time.Sleep(time.Duration(mills.(int)) * time.Millisecond)
		return nil
	}

	locals := make(map[string]interface{})
	for i, arg := range expr.Arguments {
		parameter := expr.Function.Parameters[i]
		argument := evl.EvaluateExpression(arg)
		locals[parameter.GetFingerprint()] = argument
	}

	evl.PushTheseLocals(locals)

	body := evl.Functions[expr.Function.GetFingerprint()].Body
	result := evl.EvaluateStatement(body)

	evl.PopLocals()

	return result
}

func (evl *Evaluator) EvaluateConversionExpression(expr boundnodes.BoundConversionExpressionNode) interface{} {
	value := evl.EvaluateExpression(expr.Expression)

	if expr.ToType.Fingerprint() == builtins.Any.Fingerprint() {
		return value

		// to string conversion
	} else if expr.ToType.Fingerprint() == builtins.String.Fingerprint() {
		switch value.(type) {
		case string:
			return value.(string)
		case bool:
			return fmt.Sprintf("%t", value.(bool))
		case int:
			return fmt.Sprintf("%d", value.(int))
		case float32:
			return fmt.Sprintf("%g", value.(float32))
		default:
			print.PrintC(print.Red, "No Conversion! (cringe)")
			os.Exit(-1)
		}

		// string -> bool
	} else if expr.ToType.Fingerprint() == builtins.Bool.Fingerprint() {
		switch value.(type) {
		case string:
			val, _ := strconv.ParseBool(value.(string))
			return val
		default:
			print.PrintC(print.Red, "No Conversion! (cringe)")
			os.Exit(-1)
		}
	} else if expr.ToType.Fingerprint() == builtins.Int.Fingerprint() {
		switch value.(type) {
		case string:
			val, _ := strconv.Atoi(value.(string))
			return val
		default:
			print.PrintC(print.Red, "No Conversion! (cringe)")
			os.Exit(-1)
		}
	} else if expr.ToType.Fingerprint() == builtins.Float.Fingerprint() {
		switch value.(type) {
		case string:
			val, _ := strconv.ParseFloat(value.(string), 32)
			return val
		default:
			print.PrintC(print.Red, "No Conversion! (cringe)")
			os.Exit(-1)
		}
	}

	print.PrintC(print.Red, "idek how this happened (conversion went completely wrong)")
	os.Exit(-1)
	return value
}
