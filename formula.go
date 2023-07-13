package logo

// LogicNode represents a node in a propositional logic formula.
type LogicNode interface {
	// Eval evaluates the formula represented by the LogicNode given an assignment of truth values to variables.
	Eval(assignment Assignment) bool
	// Scope returns the set of variable names that occur in the formula represented by the LogicNode.
	Scope() map[string]struct{}
	// String returns a string representation of the formula represented by the LogicNode.
	String() string
}

// Assignment represents an assignment of truth values to variables.
type Assignment map[string]bool
