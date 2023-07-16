package scrambler

import (
	"testing"

	. "github.com/dmholtz/logo"
	bf "github.com/dmholtz/logo/brute_force"
	"github.com/stretchr/testify/assert"
)

func TestMultiplyOut(t *testing.T) {
	t.Run("A & (B | C) is multiplied out to (A & B) | (A & C)", func(t *testing.T) {
		f := And(Var("A"), Or(Var("B"), Var("C")))

		// assert that multiplication is successful
		result, ok := MultiplyOut(f)
		assert.True(t, ok)

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A & (B | C | D) is multiplied out to (A & B) | (A & C) | (A & D)", func(t *testing.T) {
		f := And(Var("A"), NewDisjunction(Var("B"), Var("C"), Var("D")))

		// assert that multiplication is successful
		result, ok := MultiplyOut(f)
		assert.True(t, ok)

		// assert that the disjunction has the correct number of disjuncts
		disjunction, ok := result.(*NaryOp)
		assert.True(t, ok)
		assert.Equal(t, OrOp, disjunction.Op)
		assert.Equal(t, 3, len(disjunction.Clauses))

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A | (B & C) is multiplied out to (A | B) & (A | C)", func(t *testing.T) {
		f := Or(Var("A"), And(Var("B"), Var("C")))

		// assert that multiplication is successful
		result, ok := MultiplyOut(f)
		assert.True(t, ok)

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A | (B & C & D) is multiplied out to (A | B) & (A | C) & (A | D)", func(t *testing.T) {
		f := Or(Var("A"), NewConjunction(Var("B"), Var("C"), Var("D")))

		// assert that multiplication is successful
		result, ok := MultiplyOut(f)
		assert.True(t, ok)

		// assert that the conjunction has the correct number of conjuncts
		conjunction, ok := result.(*NaryOp)
		assert.True(t, ok)
		assert.Equal(t, AndOp, conjunction.Op)
		assert.Equal(t, 3, len(conjunction.Clauses))

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A & (B & C) is not multiplied out", func(t *testing.T) {
		f := And(Var("A"), And(Var("B"), Var("C")))

		// assert that multiplication is unsuccessful
		result, ok := MultiplyOut(f)
		assert.False(t, ok)

		// assert that the result is the same as the original expression
		assert.Equal(t, f, result)
	})

}
