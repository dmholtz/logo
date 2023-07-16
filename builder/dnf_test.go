package builder

import (
	"testing"

	bf "github.com/dmholtz/logo/brute_force"

	"github.com/stretchr/testify/assert"
)

func TestDnfBuilderSetup(t *testing.T) {
	t.Run("non-positive numVars panics", func(t *testing.T) {
		assert.Panics(t, func() { NewDnfBuilder(0, 2, 2) })
	})
	t.Run("non-positive numConjunctions panics", func(t *testing.T) {
		assert.Panics(t, func() { NewDnfBuilder(1, 0, 2) })
	})
	t.Run("non-trivial numClauses panics", func(t *testing.T) {
		assert.Panics(t, func() { NewDnfBuilder(1, 1, 1) })
	})
	t.Run("len(Scope) equal to numVars", func(t *testing.T) {
		numVars := 3
		dnfBuilder := NewDnfBuilder(numVars, 2, 2)
		assert.Equal(t, numVars, len(dnfBuilder.Scope))
	})
}

func TestRandomConjunction(t *testing.T) {
	t.Run("conjunction has numClauses", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(3, 2, 2)
		conjunction := dnfBuilder.randomConjunction()
		assert.Equal(t, dnfBuilder.NumClauses, len(conjunction.Clauses))
	})
	t.Run("conjunction scope is subset of builder scope", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(3, 2, 4)
		conjunction := dnfBuilder.randomConjunction()
		for varName := range conjunction.Scope() {
			assert.Contains(t, dnfBuilder.Scope, varName)
		}
	})
}

func TestSatConjunction(t *testing.T) {
	t.Run("conjunction has numClauses", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(3, 2, 2)
		conjunction := dnfBuilder.randomConjunction()
		assert.Equal(t, dnfBuilder.NumClauses, len(conjunction.Clauses))
	})
	t.Run("conjunction scope is subset of builder scope", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(3, 2, 4)
		conjunction := dnfBuilder.satConjunction()
		for varName := range conjunction.Scope() {
			assert.Contains(t, dnfBuilder.Scope, varName)
		}
	})
	t.Run("smallest possible sat conjunction is satisfiable", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(1, 1, 2)
		conjunction := dnfBuilder.satConjunction()
		assert.True(t, bf.IsSat(conjunction))
	})
	t.Run("20 random sat conjunctions are satisfiable", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(6, 2, 4)
		for i := 0; i < 20; i++ {
			conjunction := dnfBuilder.satConjunction()
			if !bf.IsSat(conjunction) {
				t.Log(conjunction.String())
				break
			}
			assert.True(t, bf.IsSat(conjunction))
		}
	})
	t.Run("large sat conjunction is satisfiable", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(5, 2, 20)
		conjunction := dnfBuilder.satConjunction()
		assert.True(t, bf.IsSat(conjunction))
	})
}

func TestUnsatConjunction(t *testing.T) {
	t.Run("conjunction has numClauses", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(3, 2, 2)
		conjunction := dnfBuilder.randomConjunction()
		assert.Equal(t, dnfBuilder.NumClauses, len(conjunction.Clauses))
	})
	t.Run("conjunction scope is subset of builder scope", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(3, 2, 4)
		conjunction := dnfBuilder.unsatConjunction()
		for varName := range conjunction.Scope() {
			assert.Contains(t, dnfBuilder.Scope, varName)
		}
	})
	t.Run("smallest possible unsat conjunction is unsatisfiable", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(1, 1, 2)
		conjunction := dnfBuilder.unsatConjunction()
		assert.False(t, bf.IsSat(conjunction))
	})
	t.Run("20 random unsat conjunctions are unsatisfiable", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(6, 2, 4)
		for i := 0; i < 20; i++ {
			conjunction := dnfBuilder.unsatConjunction()
			assert.False(t, bf.IsSat(conjunction))
		}
	})
}

func TestSatDnf(t *testing.T) {
	t.Run("dnf has numConjunctions", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(3, 2, 2)
		dnf := dnfBuilder.BuildSat()
		assert.Equal(t, dnfBuilder.NumConjunctions, len(dnf.Clauses))
	})
	t.Run("dnf scope is subset of builder scope", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(3, 2, 4)
		dnf := dnfBuilder.BuildSat()
		for varName := range dnf.Scope() {
			assert.Contains(t, dnfBuilder.Scope, varName)
		}
	})
	t.Run("20 random sat dnf are satisfiable", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(6, 4, 4)
		for i := 0; i < 20; i++ {
			dnf := dnfBuilder.BuildSat()
			assert.True(t, bf.IsSat(&dnf))
		}
	})
}

func TestUnsatDnf(t *testing.T) {
	t.Run("dnf has numConjunctions", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(3, 2, 2)
		dnf := dnfBuilder.BuildUnsat()
		assert.Equal(t, dnfBuilder.NumConjunctions, len(dnf.Clauses))
	})
	t.Run("dnf scope is subset of builder scope", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(3, 2, 4)
		dnf := dnfBuilder.BuildUnsat()
		for varName := range dnf.Scope() {
			assert.Contains(t, dnfBuilder.Scope, varName)
		}
	})
	t.Run("20 random unsat dnf are unsatisfiable", func(t *testing.T) {
		dnfBuilder := NewDnfBuilder(6, 4, 4)
		for i := 0; i < 20; i++ {
			dnf := dnfBuilder.BuildUnsat()
			assert.False(t, bf.IsSat(&dnf))
		}
	})
}
