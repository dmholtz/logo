package scrambler

import (
	"testing"

	. "github.com/dmholtz/logo"
	bf "github.com/dmholtz/logo/brute_force"
	"github.com/stretchr/testify/assert"
)

func TestAddDoubleNegation(t *testing.T) {
	t.Run("add double negation to an arbitrary node", func(t *testing.T) {
		f := And(Var("A"), Var("B"))

		// assert that double negation has been added
		result, ok := AddDoubleNegation(f)
		assert.True(t, ok)

		// assert that the result has a double negation
		neg1, type1Ok := result.(*NotOp)
		assert.True(t, type1Ok)
		neg2, type2Ok := neg1.X.(*NotOp)
		assert.True(t, type2Ok)

		// assert that the subnode of the inner negation is the input
		assert.Equal(t, f, neg2.X)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
}

func TestRemoveDoubleNegation(t *testing.T) {
	t.Run("remove present double negation", func(t *testing.T) {
		f := Not(Not(And(Var("A"), Var("B"))))

		// assert that double negation has been removed
		result, ok := RemoveDoubleNegation(f)
		assert.True(t, ok)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("fail to remove double negation if only single negation exists", func(t *testing.T) {
		f := Not(And(Var("a"), Var("b")))

		// assert that double negation has not been removed
		result, ok := RemoveDoubleNegation(f)
		assert.False(t, ok)

		// assert that the result is the same as the input
		assert.Equal(t, f, result)
	})
}
