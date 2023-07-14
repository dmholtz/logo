package bruteforce

import (
	"fmt"

	. "github.com/dmholtz/logo"
)

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

	// iteration over map keys is non-deterministic, so we need to create a deterministic
	// mapping from bit indices to variable names
	bitIdxToVarName := make(map[uint32]string)
	var bitIdx uint32 = 0
	for varName := range scope {
		bitIdxToVarName[bitIdx] = varName
		bitIdx++
	}

	// codeded_assignment is a bit vector that encodes an assignment
	var coded_assignment uint32 = 0
	for ; coded_assignment < (1 << len(scope)); coded_assignment++ {
		assignment := make(Assignment)

		for i := uint32(0); i < uint32(len(scope)); i++ {
			varName := bitIdxToVarName[i]
			logicValue := ((coded_assignment >> i) & 1) == 1
			assignment[varName] = logicValue
		}

		if f.Eval(assignment) {
			return true
		}
	}
	return false
}
