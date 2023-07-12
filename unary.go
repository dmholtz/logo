package logo

type Not struct {
	X LogicNode
}

func (n Not) Eval(assignment Assignment) bool {
	return !n.X.Eval(assignment)
}
