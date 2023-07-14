package builder

import (
	"fmt"
	"math/rand"

	. "github.com/dmholtz/logo"
)

type DnfBuilder struct {
	Scope           []string
	NumConjunctions int // number of conjunctions in the DNF
	NumClauses      int // number of clauses in each conjunction
}

func NewDnfBuilder(numVariables int, numConjunctions, numClauses int) *DnfBuilder {
	if numVariables < 1 {
		panic("numVariables must be positive")
	}
	if numConjunctions < 1 {
		panic("numConjunctions must be positive")
	}
	if numClauses < 2 {
		panic("non-trivial conjunction must have at least to clauses")
	}
	scope := make([]string, numVariables)
	for i := 0; i < numVariables; i++ {
		scope[i] = fmt.Sprintf("x%d", i+1)
	}
	return &DnfBuilder{
		Scope:           scope,
		NumConjunctions: numConjunctions,
		NumClauses:      numClauses,
	}
}

// BuildSat returns a disjunction of conjunctions that is satisfiable
func (b *DnfBuilder) BuildSat() Disjunction {
	conjunctions := []LogicNode{b.satConjunction()}
	for i := 1; i < b.NumConjunctions; i++ {
		conjunctions = append(conjunctions, b.randomConjunction())
	}
	return Disjunction{Disjuncts: conjunctions}
}

// BuildUnsat returns a disjunction of conjunctions that is not satisfiable
func (b *DnfBuilder) BuildUnsat() Disjunction {
	conjunctions := make([]LogicNode, 0)
	for i := 0; i < b.NumConjunctions; i++ {
		conjunctions = append(conjunctions, b.unsatConjunction())
	}
	return Disjunction{Disjuncts: conjunctions}
}

// randomVariable returns a random logic variable
func (b *DnfBuilder) randomVariable() Variable {
	return Variable{Name: b.Scope[rand.Intn(len(b.Scope))]}
}

// randomNegationWrap randomly wraps a logic variable in a negation
func (b *DnfBuilder) randomNegationWrap(variable Variable) (LogicNode, bool) {
	if rand.Intn(2) == 0 {
		return variable, true
	}
	return Not(variable), false
}

// randomLiteral returns a random literal, i.e., a logic variable or its negation
func (b *DnfBuilder) randomLiteral() LogicNode {
	literal, _ := b.randomNegationWrap(b.randomVariable())
	return literal
}

// randomConjunction returns a random conjunction (not necessarily satisfiable)
func (b *DnfBuilder) randomConjunction() *Conjunction {
	clauses := make([]LogicNode, 0)
	for i := 0; i < b.NumClauses; i++ {
		clauses = append(clauses, b.randomLiteral())
	}
	return &Conjunction{Conjuncts: clauses}
}

// satConjunction returns a conjunction that is satisfiable, i.e., a conjunction
// that never contains both a variable and its negation
func (b *DnfBuilder) satConjunction() *Conjunction {
	clauses := make([]LogicNode, 0)
	usedLiterals := make(map[string]bool)
	for i := 0; i < b.NumClauses; i++ {
		randVar := b.randomVariable()
		usedPositively, ok := usedLiterals[randVar.Name]
		if !ok {
			// first time we use this variable
			randLiteral, neg := b.randomNegationWrap(randVar)
			usedLiterals[randVar.Name] = neg
			clauses = append(clauses, randLiteral)
		} else {
			if usedPositively {
				clauses = append(clauses, randVar)
			} else {
				clauses = append(clauses, Not(&randVar))
			}
		}
	}
	return &Conjunction{Conjuncts: clauses}
}

// unsatConjunction returns a conjunction that is not satisfiable, i.e., a conjunction
// that contains both a variable and its negation
func (b *DnfBuilder) unsatConjunction() *Conjunction {
	randVar := b.randomVariable()
	negatedVar := Not(&Variable{Name: randVar.Name})
	clauses := []LogicNode{randVar, negatedVar}
	for i := 0; i < b.NumClauses-2; i++ {
		clauses = append(clauses, b.randomLiteral())
	}
	return &Conjunction{Conjuncts: clauses}
}
