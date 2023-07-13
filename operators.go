package logo

import "fmt"

type OpType int

// OpType is an enumeration of the different types of binary logical operators.
const (
	AndOp OpType = iota // X AND Y
	OrOp                // X OR Y
	IfOp                // IF X THEN Y
	IffOp               // X IFF Y
)

type NotOp struct {
	X LogicNode
}

type BinaryOp struct {
	X, Y LogicNode
	Op   OpType
}

// Not returns a LogicNode that represents the logical NOT of x.
func Not(x LogicNode) LogicNode {
	return &NotOp{X: x}
}

// And returns a LogicNode that represents the logical AND of x and y.
func And(x, y LogicNode) LogicNode {
	return &BinaryOp{X: x, Y: y, Op: AndOp}
}

// Or returns a LogicNode that represents the logical OR of x and y.
func Or(x, y LogicNode) LogicNode {
	return &BinaryOp{X: x, Y: y, Op: OrOp}
}

// Implies returns a LogicNode that represents the logical formula IF X THEN Y.
func Implies(x, y LogicNode) LogicNode {
	return &BinaryOp{X: x, Y: y, Op: IfOp}
}

// Iff returns a LogicNode that represents the logical formula X IFF Y.
func Iff(x, y LogicNode) LogicNode {
	return &BinaryOp{X: x, Y: y, Op: IffOp}
}

func (n NotOp) Eval(assignment Assignment) bool {
	return !n.X.Eval(assignment)
}

func (n NotOp) Scope() map[string]struct{} {
	return n.X.Scope()
}

func (n NotOp) String() string {
	return "!" + n.X.String()
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
