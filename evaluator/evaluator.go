package evaluator

import (
	"ReCT-Go-Compiler/binder"
	"ReCT-Go-Compiler/builtins"
	"ReCT-Go-Compiler/info"
	"ReCT-Go-Compiler/nodes/boundnodes"
	"ReCT-Go-Compiler/print"
	"ReCT-Go-Compiler/symbols"
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh/terminal"
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

func (evl *Evaluator) GetLocal(name string) interface{} {
	return evl.Locals[len(evl.Locals)-1][name]
}

func (evl *Evaluator) SetLocal(fingerprint string, value interface{}) {
	evl.Locals[len(evl.Locals)-1][fingerprint] = value
}

// variable helpers
func (evl *Evaluator) Assign(sym symbols.VariableSymbol, value interface{}) {
	if sym.IsGlobal() {
		evl.Globals[sym.Fingerprint()] = value
	} else {
		evl.SetLocal(sym.Fingerprint(), value)
	}
}

func (evl *Evaluator) Read(sym symbols.VariableSymbol) interface{} {
	if sym.IsGlobal() {
		return evl.Globals[sym.Fingerprint()]
	} else {
		return evl.GetLocal(sym.Fingerprint())
	}
}

var reader *bufio.Reader
var cursorVisible bool = true

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
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run() // disable input buffering
	rand.Seed(time.Now().UnixNano())

	evaluator.PushLocals()

	for _, fnc := range program.Functions {
		symbol := fnc.Symbol
		evaluator.Functions[symbol.Fingerprint()] = fnc
	}

	mainFunction := program.MainFunction
	body := evaluator.Functions[mainFunction.Fingerprint()].Body
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

	case boundnodes.BoundTypeCallExpression:
		return evl.EvaluateTypeCallExpression(expr.(boundnodes.BoundTypeCallExpressionNode))

	case boundnodes.BoundConversionExpression:
		return evl.EvaluateConversionExpression(expr.(boundnodes.BoundConversionExpressionNode))
	}

	return nil
}

func (evl *Evaluator) EvaluateLiteralExpression(expr boundnodes.BoundLiteralExpressionNode) interface{} {
	return expr.Value
}

func (evl *Evaluator) EvaluateVariableExpression(expr boundnodes.BoundVariableExpressionNode) interface{} {
	return evl.Read(expr.Variable)
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
	if expr.Function.Fingerprint() == builtins.Print.Fingerprint() {
		text := evl.EvaluateExpression(expr.Arguments[0])
		fmt.Println(text)
		return nil

	} else if expr.Function.Fingerprint() == builtins.Write.Fingerprint() {
		text := evl.EvaluateExpression(expr.Arguments[0])
		fmt.Print(text)
		return nil

	} else if expr.Function.Fingerprint() == builtins.Input.Fingerprint() {
		text, _ := reader.ReadString('\n')
		return text

	} else if expr.Function.Fingerprint() == builtins.InputKey.Fingerprint() {
		b := make([]byte, 1)
		os.Stdin.Read(b)
		return string(b)

	} else if expr.Function.Fingerprint() == builtins.Clear.Fingerprint() {
		fmt.Print("\033[2J")            // clear screen
		fmt.Printf("\033[%d;%dH", 0, 0) // set cursor
		return nil

	} else if expr.Function.Fingerprint() == builtins.SetCursor.Fingerprint() {
		x := evl.EvaluateExpression(expr.Arguments[0])
		y := evl.EvaluateExpression(expr.Arguments[1])
		fmt.Printf("\033[%d;%dH", y, x) // set cursor
		return nil

	} else if expr.Function.Fingerprint() == builtins.SetCursorVisible.Fingerprint() {
		visible := evl.EvaluateExpression(expr.Arguments[0])
		cursorVisible = visible.(bool)

		if cursorVisible {
			fmt.Print("\033[?25h")
		} else {
			fmt.Print("\033[?25l")
		}

		return nil

	} else if expr.Function.Fingerprint() == builtins.GetCursorVisible.Fingerprint() {
		return cursorVisible

	} else if expr.Function.Fingerprint() == builtins.GetSizeX.Fingerprint() {
		width, _, _ := terminal.GetSize(0)
		return width

	} else if expr.Function.Fingerprint() == builtins.GetSizeY.Fingerprint() {
		_, height, _ := terminal.GetSize(0)
		return height

	} else if expr.Function.Fingerprint() == builtins.Random.Fingerprint() {
		max := evl.EvaluateExpression(expr.Arguments[0])
		return rand.Intn(max.(int))

	} else if expr.Function.Fingerprint() == builtins.Sleep.Fingerprint() {
		mills := evl.EvaluateExpression(expr.Arguments[0])
		time.Sleep(time.Duration(mills.(int)) * time.Millisecond)
		return nil

	} else if expr.Function.Fingerprint() == builtins.Version.Fingerprint() {
		return info.RECT_VERSION
	}

	locals := make(map[string]interface{})
	for i, arg := range expr.Arguments {
		parameter := expr.Function.Parameters[i]
		argument := evl.EvaluateExpression(arg)
		locals[parameter.Fingerprint()] = argument
	}

	evl.PushTheseLocals(locals)

	body := evl.Functions[expr.Function.Fingerprint()].Body
	result := evl.EvaluateStatement(body)

	evl.PopLocals()

	return result
}

func (evl *Evaluator) EvaluateTypeCallExpression(expr boundnodes.BoundTypeCallExpressionNode) interface{} {
	// get value of variable
	varValue := evl.Read(expr.Variable)

	// check if the given functions fingerprint matches the GetLength() function's fingerprint
	if expr.Function.Fingerprint() == builtins.GetLength.Fingerprint() {
		// return the length
		return len(varValue.(string))

	} else if expr.Function.Fingerprint() == builtins.Substring.Fingerprint() {
		// get index and length of the substring
		index := evl.EvaluateExpression(expr.Arguments[0])
		length := evl.EvaluateExpression(expr.Arguments[1])

		// create the substring
		return varValue.(string)[index.(int) : index.(int)+length.(int)]
	} else {
		print.PrintCF(print.Red, "Unknown type function! [%s]", expr.Function.Fingerprint())
		os.Exit(-1)

		return nil
	}
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
			if value == nil {
				return "cringe"
			}

			print.PrintCF(print.Red, "No Conversion! (cringe) [%s -> %s]", expr.Expression.Type().Fingerprint(), expr.ToType.Fingerprint())
			os.Exit(-1)
		}

		// string -> bool
	} else if expr.ToType.Fingerprint() == builtins.Bool.Fingerprint() {
		switch value.(type) {
		case string:
			val, _ := strconv.ParseBool(value.(string))
			return val
		default:
			print.PrintCF(print.Red, "No Conversion! (cringe) [%s -> %s]", expr.Expression.Type().Fingerprint(), expr.ToType.Fingerprint())
			os.Exit(-1)
		}
	} else if expr.ToType.Fingerprint() == builtins.Int.Fingerprint() {
		switch value.(type) {
		case string:
			val, _ := strconv.Atoi(value.(string))
			return val
		default:
			print.PrintCF(print.Red, "No Conversion! (cringe) [%s -> %s]", expr.Expression.Type().Fingerprint(), expr.ToType.Fingerprint())
			os.Exit(-1)
		}
	} else if expr.ToType.Fingerprint() == builtins.Float.Fingerprint() {
		switch value.(type) {
		case string:
			val, _ := strconv.ParseFloat(value.(string), 32)
			return val
		default:
			print.PrintCF(print.Red, "No Conversion! (cringe) [%s -> %s]", expr.Expression.Type().Fingerprint(), expr.ToType.Fingerprint())
			os.Exit(-1)
		}
	}

	print.PrintC(print.Red, "idek how this happened (conversion went completely wrong)")
	os.Exit(-1)
	return value
}
