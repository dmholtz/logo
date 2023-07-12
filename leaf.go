package logo

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
