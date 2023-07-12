package logo

type LogicNode interface {
	Eval(assignment Assignment) bool
}

type Assignment map[string]bool
