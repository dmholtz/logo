package bruteforce

import (
	. "github.com/dmholtz/logo"
)

// IsTaut returns true iff the given formula f is a tautology.
// It does so by negating the formula and checking whether the
// negated formula is not satisfiable.
//
// The runtime of this approach is exponential and thus only feasible
// for small formulas.
func IsTaut(f LogicNode) bool {
	negated := Not(f)
	return !IsSat(negated)
}
