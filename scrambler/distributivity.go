package scrambler

import . "github.com/dmholtz/logo"

// MultiplyOut applies the distributive law to an expression of the form
// A & (B | C) and outputs (A & B) | (A & C)
func MultiplyOut(f LogicNode) (LogicNode, bool) {
	if binaryOp, ok := f.(*BinaryOp); ok {
		switch binaryOp.Op {
		case AndOp:
			if orOp, ok := binaryOp.Y.(*BinaryOp); ok && orOp.Op == OrOp {
				return Or(And(binaryOp.X, orOp.X), And(binaryOp.X, orOp.Y)), true
			}
			if disjunction, ok := binaryOp.Y.(*NaryOp); ok && disjunction.Op == OrOp {
				result := NewDisjunction()
				for _, clause := range disjunction.Clauses {
					result.Clauses = append(result.Clauses, And(binaryOp.X, clause))
				}
				return result, true
			}
		case OrOp:
			if andOp, ok := binaryOp.Y.(*BinaryOp); ok && andOp.Op == AndOp {
				return And(Or(binaryOp.X, andOp.X), Or(binaryOp.X, andOp.Y)), true
			}
			if conjunction, ok := binaryOp.Y.(*NaryOp); ok && conjunction.Op == AndOp {
				result := NewConjunction()
				for _, clause := range conjunction.Clauses {
					result.Clauses = append(result.Clauses, Or(binaryOp.X, clause))
				}
				return result, true
			}
		}
	}
	return f, false
}
