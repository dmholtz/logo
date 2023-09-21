package scrambler

import (
	"testing"

	. "github.com/dmholtz/logo"
	bf "github.com/dmholtz/logo/brute_force"
	"github.com/stretchr/testify/assert"
)

func TestCommute(t *testing.T) {
	t.Run("commute switches operands of AND", func(t *testing.T) {
		f := And(Var("A"), Var("B"))

		// assert that commutation is successful
		result, ok := Commute(f)
		assert.True(t, ok)

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})

	t.Run("commute switches operands of OR", func(t *testing.T) {
		f := Or(Var("A"), Var("B"))

		// assert that commutation is successful
		result, ok := Commute(f)
		assert.True(t, ok)

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})

	t.Run("commute switches operands of Iff", func(t *testing.T) {
		f := Iff(Var("A"), Var("B"))

		// assert that commutation is successful
		result, ok := Commute(f)
		assert.True(t, ok)

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})

	t.Run("commute permutes the operands of NaryOp such as conjunction", func(t *testing.T) {
		f := NewConjunction(Var("A"), Var("B"), Var("C"), Var("D"))

		// assert that commutation is successful
		result, ok := Commute(f)
		assert.True(t, ok)

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})

	t.Run("commute does not switch operands of Implies", func(t *testing.T) {
		f := Implies(Var("A"), Var("B"))

		// assert that commutation is not successful
		_, ok := Commute(f)
		assert.False(t, ok)
	})
}
