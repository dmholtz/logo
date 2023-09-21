package builder

import (
	. "github.com/dmholtz/logo"
)

type EquivalentFormulaBuilder struct {
	Scope       []string
	BaseFormula LogicNode
	Appendix    LogicNode
}

func NewEquivalentFormulaBuilder(numVariables int) *EquivalentFormulaBuilder {
	if numVariables < 2 {
		panic("numVariables must be at least 2")
	}
	scope := BuildScope(numVariables)

	// build a random formula as a base formula
	rbf := NewRandomFormulaBuilder(numVariables - 1)
	baseFormula := rbf.Build(6)

	// build a random variable independent of the base formula
	appendix := Var(scope[len(scope)-1])

	return &EquivalentFormulaBuilder{
		Scope:       scope,
		BaseFormula: baseFormula,
		Appendix:    appendix,
	}
}

// Question() returns a random formula as a reference for later calls to Equivalent() and NotEquivalent()
func (efb *EquivalentFormulaBuilder) Question() LogicNode {
	return And(efb.Appendix, efb.BaseFormula)
}

// Equivalent() returns a formula that is equivalent to the reference formula
func (efb *EquivalentFormulaBuilder) Equivalent() LogicNode {
	return And(efb.BaseFormula, efb.Appendix)
}

// NotEquivalent() returns a formula that is not equivalent to the reference formula
func (efb *EquivalentFormulaBuilder) NotEquivalent() LogicNode {
	return Or(efb.BaseFormula, Not(efb.Appendix))
}
