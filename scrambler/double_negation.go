package scrambler

import . "github.com/dmholtz/logo"

// AddDoubleNegation adds a double negation to the given LogicNode
func AddDoubleNegation(f LogicNode) (LogicNode, bool) {
	return Not(Not(f)), true
}

// RemoveDoubleNegation removes a double negation from the given LogicNode if possible.
func RemoveDoubleNegation(f LogicNode) (LogicNode, bool) {
	if f1, ok := f.(*NotOp); ok {
		if f2, ok := f1.X.(*NotOp); ok {
			return f2.X, true
		}
	}
	return f, false
}
