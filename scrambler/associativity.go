package scrambler

import . "github.com/dmholtz/logo"

// CombineAnd combines a tree of nested Ands into a single conjunction by applying the associativity rule.
// Associativity: (A & B) & C = A & (B & C) = A & B & C
func CombineAnd(f LogicNode) (LogicNode, bool) {
	conjunction := NewConjunction()

	extendConjunction := func(node LogicNode) {
		switch n := node.(type) {
		case *BinaryOp:
			if n.Op == AndOp {
				conjunction.Conjuncts = append(conjunction.Conjuncts, n.X)
				conjunction.Conjuncts = append(conjunction.Conjuncts, n.Y)
			} else {
				conjunction.Conjuncts = append(conjunction.Conjuncts, n)
			}
		case *Conjunction:
			conjunction.Conjuncts = append(conjunction.Conjuncts, n.Conjuncts...)
		default:
			conjunction.Conjuncts = append(conjunction.Conjuncts, n)
		}
	}

	switch operator := f.(type) {
	case *BinaryOp:
		if operator.Op == AndOp {
			extendConjunction(operator.X)
			extendConjunction(operator.Y)
			return conjunction, true
		}
	case *Conjunction:
		for _, conjunct := range operator.Conjuncts {
			extendConjunction(conjunct)
		}
		return conjunction, true
	}
	return f, false
}

// CombineOr combines a tree of nested Ors into a single disjunction by applying the associativity rule.
// Associativity: (A | B) | C = A | (B | C) = A | B | C
func CombineOr(f LogicNode) (LogicNode, bool) {
	disjunction := NewDisjunction()

	extendDisjunction := func(node LogicNode) {
		switch n := node.(type) {
		case *BinaryOp:
			if n.Op == OrOp {
				disjunction.Disjuncts = append(disjunction.Disjuncts, n.X)
				disjunction.Disjuncts = append(disjunction.Disjuncts, n.Y)
			} else {
				disjunction.Disjuncts = append(disjunction.Disjuncts, n)
			}
		case *Disjunction:
			disjunction.Disjuncts = append(disjunction.Disjuncts, n.Disjuncts...)
		default:
			disjunction.Disjuncts = append(disjunction.Disjuncts, n)
		}
	}

	switch operator := f.(type) {
	case *BinaryOp:
		if operator.Op == OrOp {
			extendDisjunction(operator.X)
			extendDisjunction(operator.Y)
			return disjunction, true
		}
	case *Disjunction:
		for _, disjunct := range operator.Disjuncts {
			extendDisjunction(disjunct)
		}
		return disjunction, true
	}
	return f, false
}
