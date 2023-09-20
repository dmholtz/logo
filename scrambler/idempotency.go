package scrambler

import . "github.com/dmholtz/logo"

// RemoveIdempotency removes duplicate clauses from a disjunction or conjunction
func RemoveIdempotency(f LogicNode) (LogicNode, bool) {
	switch operator := f.(type) {
	case *NaryOp:
		clauses := make(map[string]LogicNode, 0)
		for _, clause := range operator.Clauses {
			clauses[clause.String()] = clause
		}

		clauseSlice := make([]LogicNode, 0)
		for _, clause := range clauses {
			clauseSlice = append(clauseSlice, clause)
		}

		return &NaryOp{Op: operator.Op, Clauses: clauseSlice}, true
	}
	return f, false
}
