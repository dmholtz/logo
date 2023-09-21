package builder

import (
	"fmt"
	"math/rand"

	. "github.com/dmholtz/logo"
)

type RandomFormulaBuilder struct {
	Scope []string
}

func NewRandomFormulaBuilder(numVariables int) *RandomFormulaBuilder {
	if numVariables < 1 {
		panic("numVariables must be positive")
	}

	scope := BuildScope(numVariables)
	return &RandomFormulaBuilder{
		Scope: scope,
	}
}

// Randomly returns a unary or a binary operator
func randomOperator() LogicNode {
	opType := rand.Intn(5)
	if opType < 4 {
		return &BinaryOp{Op: OpType(opType)}
	} else {
		return &NotOp{}
	}
}

func (rfb *RandomFormulaBuilder) RandomVariable() LogicNode {
	return Var(rfb.Scope[rand.Intn(len(rfb.Scope))])
}

func (rfb *RandomFormulaBuilder) Build(numOperators int) LogicNode {
	if numOperators < 1 {
		panic("numOperators must be positive")
	}

	// choose random operators
	operators := make([]LogicNode, 0)
	for i := 0; i < numOperators; i++ {
		operators = append(operators, randomOperator())
	}

	// grow the formula tree to the right
	formula := operators[0]
	op := formula
	for i := 1; i < numOperators; i++ {
		switch op1 := op.(type) {
		case *BinaryOp:
			op1.X = rfb.RandomVariable()
			op1.Y = operators[i]
			op = op1.Y
		case *NotOp:
			op1.X = operators[i]
			op = op1.X
		default:
			panic(fmt.Sprintf("Unsupported operator type=%T of operator=%s", op1, op1))
		}
	}

	// assign variables to the leaves
	switch op1 := op.(type) {
	case *BinaryOp:
		op1.X = rfb.RandomVariable()
		op1.Y = rfb.RandomVariable()
	case *NotOp:
		op1.X = rfb.RandomVariable()
	default:
		panic(fmt.Sprintf("Unsupported operator type=%T of operator=%s", op1, op1))
	}

	return formula
}
