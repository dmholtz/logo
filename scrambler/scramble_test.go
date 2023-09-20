package scrambler

import (
	"testing"

	. "github.com/dmholtz/logo"
	bf "github.com/dmholtz/logo/brute_force"
	"github.com/stretchr/testify/assert"
)

func TestSimplify(t *testing.T) {
	t.Run("simplify yields an equivalent formula", func(t *testing.T) {
		f := NewConjunction(Not(Not(Var("A"))), Var("A"), Var("B"), Var("B"), Var("C"), Var("C"))

		result := Simplify(f)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
}

func TestSubstituteArrows(t *testing.T) {
	t.Run("SubstituteArrows yields an equivalent formula", func(t *testing.T) {
		f := And(Implies(Var("A"), Var("B")), Implies(Var("B"), Var("A")))

		result := SubstituteArrows(f)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
}

func TestDeMorganIteration(t *testing.T) {
	t.Run("DeMorganIteration yields an equivalent formula", func(t *testing.T) {
		f := Not(Or(Var("A"), Var("B")))

		result := DeMorganIteration(f)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
}
