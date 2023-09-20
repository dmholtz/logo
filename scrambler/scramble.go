package scrambler

import (
	. "github.com/dmholtz/logo"
)

// Simplify removes double negations, combines operators, and removes idempotency
func Simplify(f LogicNode) LogicNode {
	f = Traverse(f, RemoveDoubleNegation)
	f = Traverse(f, Combine)
	f = Traverse(f, RemoveIdempotency)
	return f
}

// SubstituteArrows inserts implication and equivalence operators where possible
func SubstituteArrows(f LogicNode) LogicNode {
	f = Traverse(f, SplitNary)
	f = Traverse(f, SubstituteByImplies)
	f = Traverse(f, SubstituteByIff)
	return f
}

// DeMorganIteration applies iteratively applies DeMorgan's laws to a formula
func DeMorganIteration(f LogicNode) LogicNode {
	f = Traverse(f, RemoveIff)
	f = Traverse(f, RemoveImplies)
	f = Simplify(f)
	f = Traverse(f, SplitNary)

	for i := 0; i < 5; i++ {
		f = TraverseProbabilistic(f, DeMorganExpandEager, 0.5)
		f = Traverse(f, RemoveDoubleNegation)
	}

	return f
}
