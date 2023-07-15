package scrambler

import . "github.com/dmholtz/logo"

// DeMorganExpand expands negated conjunction / disjunction into a disjunction / conjunction
// of negated clauses by applying DeMorgan's rule if possible.
func DeMorganExpand(f LogicNode) (LogicNode, bool) {
	if notNode, ok := f.(*NotOp); ok {
		switch child := notNode.X.(type) {
		case *BinaryOp:
			if child.Op == AndOp {
				return Or(Not(child.X), Not(child.Y)), true
			}
			if child.Op == OrOp {
				return And(Not(child.X), Not(child.Y)), true
			}
		case *Conjunction:
			disjunction := NewDisjunction()
			for _, clause := range child.Conjuncts {
				disjunction.Disjuncts = append(disjunction.Disjuncts, Not(clause))
			}
			return disjunction, true
		case *Disjunction:
			conjunction := NewConjunction()
			for _, clause := range child.Disjuncts {
				conjunction.Conjuncts = append(conjunction.Conjuncts, Not(clause))
			}
			return conjunction, true
		default:
		}
	}
	return f, false
}

// DeMorganExpandEager works like DeMorganExpand but is more eagerly, i.e., it may in addition
// apply double negation to expand formulas like (A & B) where DeMorganExpand would fail.
func DeMorganExpandEager(f LogicNode) (LogicNode, bool) {
	if _, isNotOp := f.(*NotOp); !isNotOp {
		// f is not a negation, apply double negation
		fPrime := Not(f)
		result, ok := DeMorganExpand(fPrime)
		if ok {
			return Not(result), true
		} else {
			return f, false
		}
	}
	return DeMorganExpand(f)
}

// DeMorganContract contracts a conjunction / disjunction of negated clauses into a negated
// disjunction / conjunction by applying DeMorgan's rule if possible.
func DeMorganContract(f LogicNode) (LogicNode, bool) {
	switch operator := f.(type) {
	case *BinaryOp:
		not1, isNot1 := operator.X.(*NotOp)
		not2, isNot2 := operator.Y.(*NotOp)
		if isNot1 && isNot2 {
			switch operator.Op {
			case AndOp:
				return Not(Or(not1.X, not2.X)), true
			case OrOp:
				return Not(And(not1.X, not2.X)), true
			}
		}
	case *Conjunction:
		disjunction := NewDisjunction()
		for _, clause := range operator.Conjuncts {
			if not, isNot := clause.(*NotOp); isNot {
				disjunction.Disjuncts = append(disjunction.Disjuncts, not.X)
			} else {
				return f, false
			}
		}
		return Not(disjunction), true
	case *Disjunction:
		conjunction := NewConjunction()
		for _, clause := range operator.Disjuncts {
			if not, isNot := clause.(*NotOp); isNot {
				conjunction.Conjuncts = append(conjunction.Conjuncts, not.X)
			} else {
				return f, false
			}
		}
		return Not(conjunction), true
	}
	return f, false
}

// DeMorganContractEager works like DeMorganContract but is more eagerly, i.e., it may in addition
// apply double negation on clauses to contract formulas like (A & !B) where DeMorganContract would fail.
func DeMorganContractEager(f LogicNode) (LogicNode, bool) {
	switch operator := f.(type) {
	case *BinaryOp:
		xOperand, yOperand := operator.X, operator.Y
		if not1, isNot := xOperand.(*NotOp); !isNot {
			xOperand = Not(xOperand)
		} else {
			xOperand = not1.X
		}
		if not2, isNot := yOperand.(*NotOp); !isNot {
			yOperand = Not(yOperand)
		} else {
			yOperand = not2.X
		}
		switch operator.Op {
		case AndOp:
			return Not(Or(xOperand, yOperand)), true
		case OrOp:
			return Not(And(xOperand, yOperand)), true
		}
	case *Conjunction:
		disjunction := NewDisjunction()
		for _, clause := range operator.Conjuncts {
			if not, isNot := clause.(*NotOp); isNot {
				disjunction.Disjuncts = append(disjunction.Disjuncts, not.X)
			} else {
				// apply double negation
				disjunction.Disjuncts = append(disjunction.Disjuncts, Not(clause))
			}
		}
		return Not(disjunction), true
	case *Disjunction:
		conjunction := NewConjunction()
		for _, clause := range operator.Disjuncts {
			if not, isNot := clause.(*NotOp); isNot {
				conjunction.Conjuncts = append(conjunction.Conjuncts, not.X)
			} else {
				// apply double negation
				conjunction.Conjuncts = append(conjunction.Conjuncts, Not(clause))
			}
		}
		return Not(conjunction), true
	}
	return f, false
}
