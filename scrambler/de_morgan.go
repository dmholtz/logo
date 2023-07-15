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
			// TODO: implement
		case *Disjunction:
			// TODO: implement
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
	if binOp, ok := f.(*BinaryOp); ok {
		switch binOp.Op {
		case AndOp:
			if not1, ok := binOp.X.(*NotOp); ok {
				if not2, ok := binOp.Y.(*NotOp); ok {
					return Not(Or(not1.X, not2.X)), true
				}
			}
		case OrOp:
			if not1, ok := binOp.X.(*NotOp); ok {
				if not2, ok := binOp.Y.(*NotOp); ok {
					return Not(And(not1.X, not2.X)), true
				}
			}
		}
	} else if conj, ok := f.(*Conjunction); ok {
		conj = conj
		// TODO: implement
	} else if disj, ok := f.(*Disjunction); ok {
		disj = disj
		// TODO: implement
	}
	return f, false
}

// DeMorganContractEager works like DeMorganContract but is more eagerly, i.e., it may in addition
// apply double negation on clauses to contract formulas like (A & !B) where DeMorganContract would fail.
func DeMorganContractEager(f LogicNode) (LogicNode, bool) {
	if binOp, ok := f.(*BinaryOp); ok {
		xOperand, yOperand := binOp.X, binOp.Y
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
		switch binOp.Op {
		case AndOp:
			return Not(Or(xOperand, yOperand)), true
		case OrOp:
			return Not(And(xOperand, yOperand)), true
		}
	} else if conj, ok := f.(*Conjunction); ok {
		conj = conj
		// TODO: implement
	} else if disj, ok := f.(*Disjunction); ok {
		disj = disj
		// TODO: implement
	}
	return f, false
}
