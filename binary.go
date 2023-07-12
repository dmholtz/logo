package logo

type And struct {
	X, Y LogicNode
}

func (a And) Eval(assignment Assignment) bool {
	return a.X.Eval(assignment) && a.Y.Eval(assignment)
}

type Or struct {
	X, Y LogicNode
}

func (o Or) Eval(assignment Assignment) bool {
	return o.X.Eval(assignment) || o.Y.Eval(assignment)
}

type Implication struct {
	X, Y LogicNode
}

func (i Implication) Eval(assignment Assignment) bool {
	return !i.X.Eval(assignment) || i.Y.Eval(assignment)
}

type Equivalence struct {
	X, Y LogicNode
}

func (e Equivalence) Eval(assignment Assignment) bool {
	return (!e.X.Eval(assignment) && !e.Y.Eval(assignment)) || (e.X.Eval(assignment) && e.Y.Eval(assignment))
}
