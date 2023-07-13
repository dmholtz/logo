package logo

import "strings"

type Conjunction struct {
	Conjuncts []LogicNode
}

func (c *Conjunction) Eval(assignment Assignment) bool {
	for _, conjunct := range c.Conjuncts {
		if !conjunct.Eval(assignment) {
			return false
		}
	}
	return true
}

func (c *Conjunction) Scope() map[string]struct{} {
	scope := make(map[string]struct{})
	for _, conjunct := range c.Conjuncts {
		for varName := range conjunct.Scope() {
			scope[varName] = struct{}{}
		}
	}
	return scope
}

func (c *Conjunction) String() string {
	// special case: empty conjunction is true
	if len(c.Conjuncts) == 0 {
		return "true"
	}

	// Build strings with strings.Builder efficiently
	// Source: https://pkg.go.dev/strings#Builder
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(c.Conjuncts[0].String())
	if len(c.Conjuncts) > 1 {
		for _, conjunct := range c.Conjuncts[1:] {
			sb.WriteString(" & ")
			sb.WriteString(conjunct.String())
		}
	}
	sb.WriteString(")")
	return sb.String()
}
