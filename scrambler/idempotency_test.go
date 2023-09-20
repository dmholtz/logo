package scrambler

import (
	"testing"

	. "github.com/dmholtz/logo"
	bf "github.com/dmholtz/logo/brute_force"
	"github.com/stretchr/testify/assert"
)

func TestRemoveIdempotency(t *testing.T) {
	t.Run("remove duplicate variables", func(t *testing.T) {
		f := NewConjunction(Var("A"), Var("A"), Var("B"), Var("B"), Var("C"), Var("C"))

		// assert that duplicate variables have been removed
		result, ok := RemoveIdempotency(f)
		assert.True(t, ok)

		// assert that the scope of the result is the same as the input
		assert.Equal(t, f.Scope(), result.Scope())

		// assert that the operand has not changed
		naryOp, ok := result.(*NaryOp)
		assert.True(t, ok)
		assert.Equal(t, f.Op, naryOp.Op)

		// assert that the number of clauses in the result is correct
		assert.Equal(t, 3, len(naryOp.Clauses))

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})

	t.Run("remove duplicate literals", func(t *testing.T) {
		f := NewConjunction(Var("A"), Not(Var("A")), Var("B"), Not(Var("B")), Var("C"), Not(Var("C")), Var("A"), Not(Var("B")))

		// assert that duplicate literals have been removed
		result, ok := RemoveIdempotency(f)
		assert.True(t, ok)

		// assert that the scope of the result is the same as the input
		assert.Equal(t, f.Scope(), result.Scope())

		// assert that the operand has not changed
		naryOp, ok := result.(*NaryOp)
		assert.True(t, ok)
		assert.Equal(t, f.Op, naryOp.Op)

		// assert that the number of clauses in the result is correct
		assert.Equal(t, 6, len(naryOp.Clauses))

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})

	t.Run("remove duplicate clauses", func(t *testing.T) {
		f := NewDisjunction(And(Var("A"), Var("B")), And(Var("A"), Var("B")), Or(Var("A"), Var("C")), And(Var("A"), Var("C")), Implies(Var("A"), Var("D")), Implies(Var("A"), Var("D")))

		// assert that duplicate clauses have been removed
		result, ok := RemoveIdempotency(f)
		assert.True(t, ok)

		// assert that the scope of the result is the same as the input
		assert.Equal(t, f.Scope(), result.Scope())

		// assert that the operand has not changed
		naryOp, ok := result.(*NaryOp)
		assert.True(t, ok)

		// assert that the number of clauses in the result is correct
		assert.Equal(t, 4, len(naryOp.Clauses))

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})

	t.Run("leave other operators than NaryOp unchanged", func(t *testing.T) {
		f := And(Var("A"), Var("B"))

		// assert that the input is unchanged
		result, ok := RemoveIdempotency(f)
		assert.False(t, ok)
		assert.Equal(t, f, result)
	})
}
