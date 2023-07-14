package bruteforce

import (
	"testing"

	. "github.com/dmholtz/logo"

	"github.com/stretchr/testify/assert"
)

func TestIsEquiv(t *testing.T) {
	t.Run("Top and Top are equivalent", func(t *testing.T) {
		assert.True(t, IsEquiv(Top(), Top()))
	})
	t.Run("Bottom and Top are not equivalent", func(t *testing.T) {
		assert.False(t, IsEquiv(Bottom(), Top()))
	})
	t.Run("A is equivalent to A", func(t *testing.T) {
		assert.True(t, IsEquiv(Var("A"), Var("A")))
	})
	t.Run("A is not equivalent to B", func(t *testing.T) {
		assert.False(t, IsEquiv(Var("A"), Var("B")))
	})
	t.Run("(A -> B) & (B -> A) is equivalent to (A <-> B)", func(t *testing.T) {
		assert.True(t, IsEquiv(And(Implies(Var("A"), Var("B")), Implies(Var("B"), Var("A"))), Iff(Var("A"), Var("B"))))
	})
	t.Run("double negation equivalence", func(t *testing.T) {
		assert.True(t, IsEquiv(Not(Not(Var("A"))), Var("A")))
	})
	t.Run("deMorgan equivalence", func(t *testing.T) {
		assert.True(t, IsEquiv(Not(Or(Var("A"), Var("B"))), And(Not(Var("A")), Not(Var("B")))))
	})
	t.Run("absorption equivalence", func(t *testing.T) {
		assert.True(t, IsEquiv(Or(Var("A"), And(Var("A"), Var("B"))), Var("A")))
	})
	t.Run("idempotence OR equivalence", func(t *testing.T) {
		assert.True(t, IsEquiv(Or(Var("A"), Var("A")), Var("A")))
	})
	t.Run("idempotence AND equivalence", func(t *testing.T) {
		assert.True(t, IsEquiv(And(Var("A"), Var("A")), Var("A")))
	})
}
