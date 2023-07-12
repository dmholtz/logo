package logo

import "fmt"

type And struct {
	X, Y LogicNode
}

func (a And) Eval(assignment Assignment) bool {
	return a.X.Eval(assignment) && a.Y.Eval(assignment)
}

func (a And) String() string {
	return fmt.Sprintf("(%s & %s)", a.X.String(), a.Y.String())
}

type Or struct {
	X, Y LogicNode
}

func (o Or) Eval(assignment Assignment) bool {
	return o.X.Eval(assignment) || o.Y.Eval(assignment)
}

func (o Or) String() string {
	return fmt.Sprintf("(%s | %s)", o.X.String(), o.Y.String())
}

type Implication struct {
	X, Y LogicNode
}

func (i Implication) Eval(assignment Assignment) bool {
	return !i.X.Eval(assignment) || i.Y.Eval(assignment)
}

func (i Implication) String() string {
	return fmt.Sprintf("(%s -> %s)", i.X.String(), i.Y.String())
}

type Equivalence struct {
	X, Y LogicNode
}

func (e Equivalence) Eval(assignment Assignment) bool {
	return (!e.X.Eval(assignment) && !e.Y.Eval(assignment)) || (e.X.Eval(assignment) && e.Y.Eval(assignment))
}

func (e Equivalence) String() string {
	return fmt.Sprintf("(%s <-> %s)", e.X.String(), e.Y.String())
}
