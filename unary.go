package logo

type Not struct {
	X LogicNode
}

func (n Not) Eval(assignment Assignment) bool {
	return !n.X.Eval(assignment)
}

func (n Not) Scope() map[string]struct{} {
	return n.X.Scope()
}

func (n Not) String() string {
	return "!" + n.X.String()
}
