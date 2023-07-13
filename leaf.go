package logo

import "fmt"

type Leaf bool

func Bottom() Leaf {
	return false
}

func Top() Leaf {
	return true
}

func (l Leaf) Eval(assignment Assignment) bool {
	return bool(l)
}

func (l Leaf) Scope() map[string]struct{} {
	return map[string]struct{}{}
}

func (l Leaf) String() string {
	return fmt.Sprintf("%t", l)
}
