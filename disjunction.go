package logo

import "strings"

type Disjunction struct {
	Disjuncts []LogicNode
}

func (d *Disjunction) Eval(assignment Assignment) bool {
	for _, disjunct := range d.Disjuncts {
		if disjunct.Eval(assignment) {
			return true
		}
	}
	return false
}

func (d *Disjunction) String() string {
	// special case: empty disjunction is false
	if len(d.Disjuncts) == 0 {
		return "false"
	}

	// Build strings with strings.Builder efficiently
	// Source: https://pkg.go.dev/strings#Builder
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(d.Disjuncts[0].String())
	if len(d.Disjuncts) > 1 {
		for _, disjunct := range d.Disjuncts[1:] {
			sb.WriteString(" | ")
			sb.WriteString(disjunct.String())
		}
	}
	sb.WriteString(")")
	return sb.String()
}
