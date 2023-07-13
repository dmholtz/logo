package logo

import "fmt"

// IsSat returns true iff the given formula f is satisfiable.
// It does so by evaluating the formula for all possible assignments.
// The runtime of this approach is exponential and thus only feasible
// for small formulas.
func IsSat(f LogicNode) bool {
	scope := f.Scope()

	if len(scope) == 0 {
		return f.Eval(Assignment{})
	}
	if len(scope) > 31 {
		panic(fmt.Sprintf("Too many variables in formula f=%s: %d > 31", f, len(scope)))
	}

	// codeded_assignment is a bit vector that encodes an assignment
	var coded_assignment uint32 = 0
	for ; coded_assignment < (1 << len(scope)); coded_assignment++ {
		assignment := make(Assignment)

		var bit_idx uint32 = 0
		for varName := range scope {
			logicValue := ((coded_assignment >> bit_idx) & 1) == 1
			assignment[varName] = logicValue
			bit_idx++
		}

		if f.Eval(assignment) {
			return true
		}
	}
	return false
}
