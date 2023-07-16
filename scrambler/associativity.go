package scrambler

import (
	. "github.com/dmholtz/logo"
)

// CombineAnd combines a tree of nested ANDs or ORs into a single n-ary operator by applying the associativity rule.
// Associativity: (A & B) & C = A & (B & C) = A & B & C
// Associativity: (A | B) | C = A | (B | C) = A | B | C
func Combine(f LogicNode) (LogicNode, bool) {
	naryOp := NaryOp{}

	extendClauses := func(node LogicNode, opType OpType) {
		switch n := node.(type) {
		case *BinaryOp:
			if n.Op == opType {
				naryOp.Clauses = append(naryOp.Clauses, n.X)
				naryOp.Clauses = append(naryOp.Clauses, n.Y)
				return
			}
		case *NaryOp:
			if n.Op == opType {
				naryOp.Clauses = append(naryOp.Clauses, n.Clauses...)
				return
			}
		}
		naryOp.Clauses = append(naryOp.Clauses, node)
	}

	switch operator := f.(type) {
	case *BinaryOp:
		if operator.Op == AndOp || operator.Op == OrOp {
			naryOp.Op = operator.Op
			extendClauses(operator.X, naryOp.Op)
			extendClauses(operator.Y, naryOp.Op)
			return &naryOp, true
		}
	case *NaryOp:
		if operator.Op == AndOp || operator.Op == OrOp {
			naryOp.Op = operator.Op
			for _, clause := range operator.Clauses {
				extendClauses(clause, naryOp.Op)
			}
			return &naryOp, true
		}
	}
	return f, false
}

// SplitNary splits a n-ary conjunction or disjunction into a binary tree of conjunctions or disjunctions.
func SplitNary(f LogicNode) (LogicNode, bool) {
	switch operator := f.(type) {
	case *NaryOp:
		if len(operator.Clauses) == 0 && operator.Op == AndOp {
			return Top(), true
		}
		if len(operator.Clauses) == 0 && operator.Op == OrOp {
			return Bottom(), true
		}
		if len(operator.Clauses) == 1 {
			return operator.Clauses[0], true
		}

		// len(operator.Clauses) > 1
		if operator.Op == AndOp {
			rest, _ := SplitNary(NewConjunction(operator.Clauses[1:]...))
			return And(operator.Clauses[0], rest), true
		}
		if operator.Op == OrOp {
			rest, _ := SplitNary(NewDisjunction(operator.Clauses[1:]...))
			return Or(operator.Clauses[0], rest), true
		}
	}
	return f, false
}
