package logo

type LogicNode interface {
	Eval(assignment Assignment) bool
	Scope() map[string]struct{}
	String() string
}

type Assignment map[string]bool
