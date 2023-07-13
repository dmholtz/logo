package logo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSat(t *testing.T) {
	t.Run("top is satisfiable", func(t *testing.T) {
		assert.True(t, IsSat(Top()))
	})
	t.Run("bottom is not satisfiable", func(t *testing.T) {
		assert.False(t, IsSat(Bottom()))
	})
	t.Run("A & !A is not satisfiable", func(t *testing.T) {
		assert.False(t, IsSat(And(Var("A"), Not(Var("A")))))
	})
	t.Run("A | !A is satisfiable", func(t *testing.T) {
		assert.True(t, IsSat(Or(Var("A"), Not(Var("A")))))
	})
	t.Run("(A <-> B) & B is satisfiable", func(t *testing.T) {
		assert.True(t, IsSat(And(Iff(Var("A"), Var("B")), Var("B"))))
	})
	t.Run("(A <-> B) & A & !B is not satisfiable", func(t *testing.T) {
		assert.False(t, IsSat(And(Iff(Var("A"), Var("B")), And(Var("A"), Not(Var("B"))))))
	})
	t.Run("large formula is rejected", func(t *testing.T) {
		clauses := []LogicNode{}
		for i := 0; i < 32; i++ {
			clauses = append(clauses, Var(fmt.Sprintf("x%d", i+1)))
		}
		conjunction := &Conjunction{Conjuncts: clauses}
		assert.Panics(t, func() { IsSat(conjunction) })
	})

}
