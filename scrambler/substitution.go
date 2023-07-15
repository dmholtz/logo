package scrambler

import (
	. "github.com/dmholtz/logo"
)

// SubstituteByImplies substitutes the logic formulas (!A OR B) or (B OR !A) with (A->B) if possible.
func SubstituteByImplies(f LogicNode) (LogicNode, bool) {
	if orNode, ok := f.(*BinaryOp); ok {
		if orNode.Op == OrOp {
			// case (!A OR B)
			if f2, ok := orNode.X.(*NotOp); ok {
				return Implies(f2.X, orNode.Y), true
			}
			// case (B OR !A)
			if f2, ok := orNode.Y.(*NotOp); ok {
				return Implies(f2.X, orNode.X), true
			}
		}
	}
	return f, false
}

// SubstituteByIff substitutes the logic formula (A->B) AND (B->A) with (A<->B) if possible.
//
// Caveat: SubstituteByImplies first on all subformulas before calling this function.
func SubstituteByIff(f LogicNode) (LogicNode, bool) {
	if andNode, ok := f.(*BinaryOp); ok {
		if andNode.Op == AndOp {
			if impliesNode1, ok := andNode.X.(*BinaryOp); ok {
				if impliesNode2, ok := andNode.Y.(*BinaryOp); ok {
					if impliesNode1.Op == IfOp && impliesNode2.Op == IfOp {
						implies1XSig := impliesNode1.X.String()
						implies1YSig := impliesNode1.Y.String()
						implies2XSig := impliesNode2.X.String()
						implies2YSig := impliesNode2.Y.String()
						// compare subformulas by their string representation (signature)
						if implies1XSig == implies2YSig && implies1YSig == implies2XSig {
							return Iff(impliesNode1.X, impliesNode1.Y), true
						}
					}
				}
			}
		}
	}
	return f, false
}

// RemoveImplies substitutes the logic formula (A->B) with (!A OR B) if possible.
func RemoveImplies(f LogicNode) (LogicNode, bool) {
	if impliesNode, ok := f.(*BinaryOp); ok {
		if impliesNode.Op == IfOp {
			return Or(Not(impliesNode.X), impliesNode.Y), true
		}
	}
	return f, false
}

// RemoveIff substitutes the logic formula (A<->B) with ((!A AND !B) OR (A AND B)) if possible.
func RemoveIff(f LogicNode) (LogicNode, bool) {
	if iffNode, ok := f.(*BinaryOp); ok {
		if iffNode.Op == IffOp {
			return Or(And(Not(iffNode.X), Not(iffNode.Y)), And(iffNode.X, iffNode.Y)), true
		}
	}
	return f, false
}
