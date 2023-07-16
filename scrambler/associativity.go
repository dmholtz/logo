package scrambler

import (
	. "github.com/dmholtz/logo"
)

// CombineAnd combines a tree of nested Ands into a single conjunction by applying the associativity rule.
// Associativity: (A & B) & C = A & (B & C) = A & B & C
func CombineAnd(f LogicNode) (LogicNode, bool) {
	conjunction := NewConjunction()

	extendConjunction := func(node LogicNode) {
		switch n := node.(type) {
		case *BinaryOp:
			if n.Op == AndOp {
				conjunction.Clauses = append(conjunction.Clauses, n.X)
				conjunction.Clauses = append(conjunction.Clauses, n.Y)
			} else {
				conjunction.Clauses = append(conjunction.Clauses, n)
			}
		case *NaryOp:
			if n.Op == AndOp {
				conjunction.Clauses = append(conjunction.Clauses, n.Clauses...)
			} else {
				conjunction.Clauses = append(conjunction.Clauses, n)
			}
		default:
			conjunction.Clauses = append(conjunction.Clauses, n)
		}
	}

	switch operator := f.(type) {
	case *BinaryOp:
		if operator.Op == AndOp {
			extendConjunction(operator.X)
			extendConjunction(operator.Y)
			return conjunction, true
		}
	case *NaryOp:
		if operator.Op == AndOp {
			for _, conjunct := range operator.Clauses {
				extendConjunction(conjunct)
			}
			return conjunction, true
		}
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
				disjunction.Clauses = append(disjunction.Clauses, n.X)
				disjunction.Clauses = append(disjunction.Clauses, n.Y)
			} else {
				disjunction.Clauses = append(disjunction.Clauses, n)
			}
		case *NaryOp:
			if n.Op == OrOp {
				disjunction.Clauses = append(disjunction.Clauses, n.Clauses...)
			} else {
				disjunction.Clauses = append(disjunction.Clauses, n)
			}
		default:
			disjunction.Clauses = append(disjunction.Clauses, n)
		}
	}

	switch operator := f.(type) {
	case *BinaryOp:
		if operator.Op == OrOp {
			extendDisjunction(operator.X)
			extendDisjunction(operator.Y)
			return disjunction, true
		}
	case *NaryOp:
		if operator.Op == OrOp {
			for _, disjunct := range operator.Clauses {
				extendDisjunction(disjunct)
			}
			return disjunction, true
		}
	}
	return f, false
}

// SplitNary splits a n-ary conjunction or disjunction into a binary tree of conjunctions or disjunctions.
func SplitNary(f LogicNode) (LogicNode, bool) {
	switch operator := f.(type) {
	case *NaryOp:
		if operator.Op == AndOp {
			if len(operator.Clauses) > 1 {
				rest, _ := SplitNary(NewConjunction(operator.Clauses[1:]...))
				return And(operator.Clauses[0], rest), true
			}
			if len(operator.Clauses) == 1 {
				return operator.Clauses[0], true
			}
			if len(operator.Clauses) == 0 {
				return Top(), true
			}
		}
		if operator.Op == OrOp {
			if len(operator.Clauses) > 1 {
				rest, _ := SplitNary(NewDisjunction(operator.Clauses[1:]...))
				return Or(operator.Clauses[0], rest), true
			}
			if len(operator.Clauses) == 1 {
				return operator.Clauses[0], true
			}
			if len(operator.Clauses) == 0 {
				return Bottom(), true
			}
		}
	}
	return f, false
}
