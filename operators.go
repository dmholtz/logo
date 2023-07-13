package logo

import "fmt"

type OpType int

const (
	AndOp OpType = iota
	OrOp
	ImplicationOp
	EquivalenceOp
)

type BinaryOp struct {
	X, Y LogicNode
	Op   OpType
}

func (b BinaryOp) Eval(assignment Assignment) bool {
	switch b.Op {
	case AndOp:
		return b.X.Eval(assignment) && b.Y.Eval(assignment)
	case OrOp:
		return b.X.Eval(assignment) || b.Y.Eval(assignment)
	case ImplicationOp:
		return !b.X.Eval(assignment) || b.Y.Eval(assignment)
	case EquivalenceOp:
		return (!b.X.Eval(assignment) && !b.Y.Eval(assignment)) || (b.X.Eval(assignment) && b.Y.Eval(assignment))
	default:
		panic(fmt.Sprintf("Unknown OpType=%d", b.Op))
	}
}

func (b BinaryOp) Scope() map[string]struct{} {
	scope := make(map[string]struct{})
	for k, v := range b.X.Scope() {
		scope[k] = v
	}
	for k, v := range b.Y.Scope() {
		scope[k] = v
	}
	return scope
}

func (b BinaryOp) String() string {
	switch b.Op {
	case AndOp:
		return fmt.Sprintf("(%s & %s)", b.X.String(), b.Y.String())
	case OrOp:
		return fmt.Sprintf("(%s | %s)", b.X.String(), b.Y.String())
	case ImplicationOp:
		return fmt.Sprintf("(%s -> %s)", b.X.String(), b.Y.String())
	case EquivalenceOp:
		return fmt.Sprintf("(%s <-> %s)", b.X.String(), b.Y.String())
	default:
		panic(fmt.Sprintf("Unknown OpType=%d", b.Op))
	}
}
