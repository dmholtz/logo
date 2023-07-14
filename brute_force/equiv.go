package bruteforce

import . "github.com/dmholtz/logo"

// IsEquiv returns true iff the given formulas f and g are equivalent.
// It does so by checking whether the formula (f <-> g) is a tautology.
//
// The runtime of this approach is exponential and thus only feasible
// for small formulas.
func IsEquiv(f, g LogicNode) bool {
	return IsTaut(Iff(f, g))
}
