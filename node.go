package logo

type LogicNode interface {
	String() string
	Eval(assignment Assignment) bool
}

type Assignment map[string]bool
