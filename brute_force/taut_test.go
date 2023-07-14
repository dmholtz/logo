package bruteforce

import (
	"fmt"
	"testing"

	. "github.com/dmholtz/logo"

	"github.com/stretchr/testify/assert"
)

func TestIsTaut(t *testing.T) {
	t.Run("top is a tautology", func(t *testing.T) {
		assert.True(t, IsTaut(Top()))
	})
	t.Run("A | !A is a tautology", func(t *testing.T) {
		assert.True(t, IsTaut(Or(Var("A"), Not(Var("A")))))
	})
	t.Run("A & !A is not a tautology", func(t *testing.T) {
		assert.False(t, IsTaut(And(Var("A"), Not(Var("A")))))
	})
	t.Run("(A -> B) <-> (B -> A) is not a tautology", func(t *testing.T) {
		assert.False(t, IsTaut(And(Implies(Var("A"), Var("B")), Implies(Var("B"), Var("A")))))
	})
	t.Run("(A <-> B) <-> ((A -> B) & (B -> A)) is a tautology", func(t *testing.T) {
		assert.True(t, IsTaut(Iff(Iff(Var("A"), Var("B")), And(Implies(Var("A"), Var("B")), Implies(Var("B"), Var("A"))))))
	})

	t.Run("large formula is rejected", func(t *testing.T) {
		clauses := []LogicNode{}
		for i := 0; i < 32; i++ {
			clauses = append(clauses, Var(fmt.Sprintf("x%d", i+1)))
		}
		conjunction := &Conjunction{Conjuncts: clauses}
		assert.Panics(t, func() { IsTaut(conjunction) })
	})

}
