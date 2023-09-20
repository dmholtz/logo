package scrambler

import (
	"testing"

	. "github.com/dmholtz/logo"
	"github.com/stretchr/testify/assert"
)

// replaceVariable replaces a variable node with the variable D and leaves all other nodes unchanged
func replaceVariable(node LogicNode) (LogicNode, bool) {
	switch node.(type) {
	case *Variable:
		return Var("D"), true
	default:
		return node, false
	}
}

func TestTraversal(t *testing.T) {

	t.Run("Traverse a formula and apply a function to each node", func(t *testing.T) {
		f := NewConjunction(And(Var("A"), Or(Var("B"), Var("C"))), Not(Var("D")), Bottom())
		result := Traverse(f, replaceVariable)

		// assert that the variable D is the only variable in the formula
		assert.Equal(t, 1, len(result.Scope()))
		_, ok := result.Scope()["D"]
		assert.True(t, ok)
	})
}

func TestTraversalWithProbability(t *testing.T) {

	t.Run("Traverse a formula and apply a function to each node with a probability of 0.5", func(t *testing.T) {
		f := NewConjunction(And(Var("A"), Or(Var("B"), Var("C"))), Not(Var("D")), Bottom())
		result := TraverseProbabilistic(f, replaceVariable, 0.5)

		// assert that the variable D is in scope
		assert.True(t, len(result.Scope()) <= 4 && len(result.Scope()) >= 1)
		_, ok := result.Scope()["D"]
		assert.True(t, ok)
	})

	t.Run("Panic if the probability is less than 0", func(t *testing.T) {
		f := NewConjunction(And(Var("A"), Or(Var("B"), Var("C"))), Not(Var("D")), Bottom())
		assert.Panics(t, func() { TraverseProbabilistic(f, replaceVariable, -0.1) })
	})

	t.Run("Panic if the probability is greater than 1", func(t *testing.T) {
		f := NewConjunction(And(Var("A"), Or(Var("B"), Var("C"))), Not(Var("D")), Bottom())
		assert.Panics(t, func() { TraverseProbabilistic(f, replaceVariable, 1.1) })
	})
}
