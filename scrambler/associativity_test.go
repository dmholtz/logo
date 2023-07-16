package scrambler

import (
	"testing"

	. "github.com/dmholtz/logo"
	bf "github.com/dmholtz/logo/brute_force"

	"github.com/stretchr/testify/assert"
)

func TestCombineAnd(t *testing.T) {
	t.Run("(A & B) & C is combined to (A & B & C) ", func(t *testing.T) {
		f := And(And(Var("A"), Var("B")), Var("C"))

		// assert that combination is successful
		result, ok := CombineAnd(f)
		assert.True(t, ok)

		// assert that the result is a conjunction of correct length
		conjunction, ok := result.(*NaryOp)
		assert.True(t, ok)
		assert.Equal(t, 3, len(conjunction.Clauses))

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A & (B & C & D) is combined to (A & B & C & D) ", func(t *testing.T) {
		f := And(Var("A"), NewConjunction(Var("B"), Var("C"), Var("D")))

		// assert that combination is successful
		result, ok := CombineAnd(f)
		assert.True(t, ok)

		// assert that the result is a conjunction of correct length
		conjunction, ok := result.(*NaryOp)
		assert.True(t, ok)
		assert.Equal(t, 4, len(conjunction.Clauses))

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A & (B & C) & D is combined to (A & B & C & D) ", func(t *testing.T) {
		f := NewConjunction(Var("A"), NewConjunction(Var("B"), Var("C")), Var("D"))

		// assert that combination is successful
		result, ok := CombineAnd(f)
		assert.True(t, ok)

		// assert that the result is a conjunction of correct length
		conjunction, ok := result.(*NaryOp)
		assert.True(t, ok)
		assert.Equal(t, 4, len(conjunction.Clauses))

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("(A | C) & (B | C) is combined to (A | C) & (B | C) ", func(t *testing.T) {
		f := And(Or(Var("A"), Var("C")), Or(Var("B"), Var("C")))

		// assert that combination is successful
		result, ok := CombineAnd(f)
		assert.True(t, ok)

		// assert that the result is a conjunction of correct length
		conjunction, ok := result.(*NaryOp)
		assert.True(t, ok)
		assert.Equal(t, 2, len(conjunction.Clauses))

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A | B is not combined", func(t *testing.T) {
		f := Or(Var("A"), Var("B"))

		// assert that combination is unsuccessful
		result, ok := CombineAnd(f)
		assert.False(t, ok)

		// assert that the result is the same as the original expression
		assert.Equal(t, f, result)
	})
}

func TestCombineOr(t *testing.T) {
	t.Run("(A | B) | C is combined to (A | B | C) ", func(t *testing.T) {
		f := Or(Or(Var("A"), Var("B")), Var("C"))

		// assert that combination is successful
		result, ok := CombineOr(f)
		assert.True(t, ok)

		// assert that the result is a disjunction of correct length
		disjunction, ok := result.(*Disjunction)
		assert.True(t, ok)
		assert.Equal(t, 3, len(disjunction.Disjuncts))

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A | (B | C | D) is combined to (A | B | C | D) ", func(t *testing.T) {
		f := Or(Var("A"), NewDisjunction(Var("B"), Var("C"), Var("D")))

		// assert that combination is successful
		result, ok := CombineOr(f)
		assert.True(t, ok)

		// assert that the result is a disjunction of correct length
		disjunction, ok := result.(*Disjunction)
		assert.True(t, ok)
		assert.Equal(t, 4, len(disjunction.Disjuncts))

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A | (B | C) | D is combined to (A | B | C | D) ", func(t *testing.T) {
		f := NewDisjunction(Var("A"), NewDisjunction(Var("B"), Var("C")), Var("D"))

		// assert that combination is successful
		result, ok := CombineOr(f)
		assert.True(t, ok)

		// assert that the result is a disjunction of correct length
		disjunction, ok := result.(*Disjunction)
		assert.True(t, ok)
		assert.Equal(t, 4, len(disjunction.Disjuncts))

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("(A & C) | (B & C) is combined to (A & C) | (B & C)", func(t *testing.T) {
		f := Or(And(Var("A"), Var("C")), And(Var("B"), Var("C")))

		// assert that combination is successful
		result, ok := CombineOr(f)
		assert.True(t, ok)

		// assert that the result is a disjunction of correct length
		disjunction, ok := result.(*Disjunction)
		assert.True(t, ok)
		assert.Equal(t, 2, len(disjunction.Disjuncts))

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A & B is not combined", func(t *testing.T) {
		f := And(Var("A"), Var("B"))

		// assert that combination is unsuccessful
		result, ok := CombineOr(f)
		assert.False(t, ok)

		// assert that the result is the same as the original expression
		assert.Equal(t, f, result)
	})
}

func TestSplit(t *testing.T) {
	t.Run("A & B & C is split to (A & (B & C))", func(t *testing.T) {
		f := NewConjunction(Var("A"), Var("B"), Var("C"))

		// assert that split is successful
		result, ok := SplitNary(f)
		assert.True(t, ok)

		// assert that the result is a AND operator
		andOp, ok := result.(*BinaryOp)
		assert.True(t, ok)
		assert.Equal(t, andOp.Op, AndOp)

		// assert that the tree grows to the right
		andOpRight, ok := andOp.Y.(*BinaryOp)
		assert.True(t, ok)
		assert.Equal(t, andOpRight.Op, AndOp)

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("(A&) is split to A", func(t *testing.T) {
		f := NewConjunction(Var("A"))

		// assert that split is successful
		result, ok := SplitNary(f)
		assert.True(t, ok)

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("(&) is split to Top", func(t *testing.T) {
		f := NewConjunction()

		// assert that split is successful
		result, ok := SplitNary(f)
		assert.True(t, ok)

		// assert that the result is Top()
		assert.Equal(t, Top(), result)

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A | B | C is split to (A | (B | C))", func(t *testing.T) {
		f := NewDisjunction(Var("A"), Var("B"), Var("C"))

		// assert that split is successful
		result, ok := SplitNary(f)
		assert.True(t, ok)

		// assert that the result is a OR operator
		orOp, ok := result.(*BinaryOp)
		assert.True(t, ok)
		assert.Equal(t, orOp.Op, OrOp)

		// assert that the tree grows to the right
		orOpRight, ok := orOp.Y.(*BinaryOp)
		assert.True(t, ok)
		assert.Equal(t, orOpRight.Op, OrOp)

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("(A|) is split to A", func(t *testing.T) {
		f := NewDisjunction(Var("A"))

		// assert that split is successful
		result, ok := SplitNary(f)
		assert.True(t, ok)

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("(|) is split to Bottom", func(t *testing.T) {
		f := NewDisjunction()

		// assert that split is successful
		result, ok := SplitNary(f)
		assert.True(t, ok)

		// assert that the result is Bottom()
		assert.Equal(t, Bottom(), result)

		// assert that the result is equivalent to the original expression
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("!A is not split", func(t *testing.T) {
		f := Not(Var("A"))

		// assert that split is unsuccessful
		result, ok := SplitNary(f)
		assert.False(t, ok)

		// assert that the result is the same as the original expression
		assert.Equal(t, f, result)
	})
}
