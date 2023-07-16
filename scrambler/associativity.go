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

// SplitNary splits a n-ary conjunction or disjunction into a binary tree of conjunctions or disjunctions.
func SplitNary(f LogicNode) (LogicNode, bool) {
	switch operator := f.(type) {
	case *Conjunction:
		if len(operator.Conjuncts) > 1 {
			rest, _ := SplitNary(NewConjunction(operator.Conjuncts[1:]...))
			return And(operator.Conjuncts[0], rest), true
		}
		if len(operator.Conjuncts) == 1 {
			return operator.Conjuncts[0], true
		}
		if len(operator.Conjuncts) == 0 {
			return Top(), true
		}
	case *Disjunction:
		if len(operator.Disjuncts) > 1 {
			rest, _ := SplitNary(NewDisjunction(operator.Disjuncts[1:]...))
			return Or(operator.Disjuncts[0], rest), true
		}
		if len(operator.Disjuncts) == 1 {
			return operator.Disjuncts[0], true
		}
		if len(operator.Disjuncts) == 0 {
			return Bottom(), true
		}
	}
	return f, false
}
