package logo

import "fmt"

type OpType int

const (
	AndOp OpType = iota // X AND Y
	OrOp                // X OR Y
	IfOp                // IF X THEN Y
	IffOp               // X IFF Y
)

type BinaryOp struct {
	X, Y LogicNode
	Op   OpType
}

func And(x, y LogicNode) LogicNode {
	return &BinaryOp{X: x, Y: y, Op: AndOp}
}

func Or(x, y LogicNode) LogicNode {
	return &BinaryOp{X: x, Y: y, Op: OrOp}
}

func If(x, y LogicNode) LogicNode {
	return &BinaryOp{X: x, Y: y, Op: IfOp}
}

func Iff(x, y LogicNode) LogicNode {
	return &BinaryOp{X: x, Y: y, Op: IffOp}
}

func (b BinaryOp) Eval(assignment Assignment) bool {
	switch b.Op {
	case AndOp:
		return b.X.Eval(assignment) && b.Y.Eval(assignment)
	case OrOp:
		return b.X.Eval(assignment) || b.Y.Eval(assignment)
	case IfOp:
		return !b.X.Eval(assignment) || b.Y.Eval(assignment)
	case IffOp:
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
	case IfOp:
		return fmt.Sprintf("(%s -> %s)", b.X.String(), b.Y.String())
	case IffOp:
		return fmt.Sprintf("(%s <-> %s)", b.X.String(), b.Y.String())
	default:
		panic(fmt.Sprintf("Unknown OpType=%d", b.Op))
	}
}
