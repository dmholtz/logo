package scrambler

import (
	"testing"

	. "github.com/dmholtz/logo"
	bf "github.com/dmholtz/logo/brute_force"
	"github.com/stretchr/testify/assert"
)

func TestDeMorganExpand(t *testing.T) {
	t.Run("!(A & B) expands to !A | !B", func(t *testing.T) {
		f := Not(And(Var("A"), Var("B")))

		// assert that DeMorgan's law is applied
		result, ok := DeMorganExpand(f)
		assert.True(t, ok)

		// assert that the result is a binary op
		binOp, typeOk := result.(*BinaryOp)
		assert.True(t, typeOk)

		// assert that both subnodes are negated
		_, typeOk = binOp.X.(*NotOp)
		assert.True(t, typeOk)
		_, typeOk = binOp.Y.(*NotOp)
		assert.True(t, typeOk)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("!(A | B) expands to !A & !B", func(t *testing.T) {
		f := Not(Or(Var("A"), Var("B")))

		// assert that DeMorgan's law is applied
		result, ok := DeMorganExpand(f)
		assert.True(t, ok)

		// assert that the result is a binary op
		binOp, typeOk := result.(*BinaryOp)
		assert.True(t, typeOk)

		// assert that both subnodes are negated
		_, typeOk = binOp.X.(*NotOp)
		assert.True(t, typeOk)
		_, typeOk = binOp.Y.(*NotOp)
		assert.True(t, typeOk)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("!(A -> B) is not expanded", func(t *testing.T) {
		f := Not(Implies(Var("A"), Var("B")))

		// assert that DeMorgan's law is not applied
		result, ok := DeMorganExpand(f)
		assert.False(t, ok)

		// assert that the result is the same as the input
		assert.Equal(t, f, result)
	})
	t.Run("A & B is not expanded", func(t *testing.T) {
		f := And(Var("A"), Var("B"))

		// assert that DeMorgan's law is not applied
		result, ok := DeMorganExpand(f)
		assert.False(t, ok)

		// assert that the result is the same as the input
		assert.Equal(t, f, result)
	})
}

func TestDeMorganExpandEager(t *testing.T) {
	t.Run("behave like DeMorganExpand", func(t *testing.T) {
		f := Not(And(Var("A"), Var("B")))

		// assert that DeMorgan's law is applied
		result, ok := DeMorganExpandEager(f)
		assert.True(t, ok)

		// assert that the result is a binary op
		binOp, typeOk := result.(*BinaryOp)
		assert.True(t, typeOk)

		// assert that both subnodes are negated
		_, typeOk = binOp.X.(*NotOp)
		assert.True(t, typeOk)
		_, typeOk = binOp.Y.(*NotOp)
		assert.True(t, typeOk)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("(A & B) expands to !(!A | !B)", func(t *testing.T) {
		f := And(Var("A"), Var("B"))

		// assert that DeMorgan's law is applied
		result, ok := DeMorganExpandEager(f)
		assert.True(t, ok)

		// assert that the result is a negated binary op
		notOp, typeOk := result.(*NotOp)
		assert.True(t, typeOk)
		binOp, typeOk := notOp.X.(*BinaryOp)
		assert.True(t, typeOk)

		// assert that the binary op is a disjunction
		assert.Equal(t, OrOp, binOp.Op)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A -> B is not expanded", func(t *testing.T) {
		f := Implies(Var("A"), Var("B"))

		// assert that DeMorgan's law is not applied
		result, ok := DeMorganExpandEager(f)
		assert.False(t, ok)

		// assert that the result is the same as the input
		assert.Equal(t, f, result)
	})
}

func TestDeMorganContract(t *testing.T) {
	t.Run("(!A | !B) contracts to !(A & B)", func(t *testing.T) {
		f := Or(Not(Var("A")), Not(Var("B")))

		// assert that DeMorgan's law is applied
		result, ok := DeMorganContract(f)
		assert.True(t, ok)

		// assert that the result is a negated binary op
		notOp, typeOk := result.(*NotOp)
		assert.True(t, typeOk)
		binOp, typeOk := notOp.X.(*BinaryOp)
		assert.True(t, typeOk)

		// assert that the binary op is a conjunction
		assert.Equal(t, AndOp, binOp.Op)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("(!A & !B) contracts to !(A | B)", func(t *testing.T) {
		f := And(Not(Var("A")), Not(Var("B")))

		// assert that DeMorgan's law is applied
		result, ok := DeMorganContract(f)
		assert.True(t, ok)

		// assert that the result is a negated binary op
		notOp, typeOk := result.(*NotOp)
		assert.True(t, typeOk)
		binOp, typeOk := notOp.X.(*BinaryOp)
		assert.True(t, typeOk)

		// assert that the binary op is a disjunction
		assert.Equal(t, OrOp, binOp.Op)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("(A & !B) is not contracted", func(t *testing.T) {
		f := And(Var("A"), Not(Var("B")))

		// assert that DeMorgan's law is not applied
		result, ok := DeMorganContract(f)
		assert.False(t, ok)

		// assert that the result is the same as the input
		assert.Equal(t, f, result)
	})
	t.Run("(A! -> B) is not contracted", func(t *testing.T) {
		f := Implies(Not(Var("A")), Var("B"))

		// assert that DeMorgan's law is not applied
		result, ok := DeMorganContract(f)
		assert.False(t, ok)

		// assert that the result is the same as the input
		assert.Equal(t, f, result)
	})
}

func TestDeMorganContractEager(t *testing.T) {
	t.Run("behave like DeMorganContract", func(t *testing.T) {
		f := Or(Not(Var("A")), Not(Var("B")))

		// assert that DeMorgan's law is applied
		result, ok := DeMorganContractEager(f)
		assert.True(t, ok)

		// assert that the result is a negated binary op
		notOp, typeOk := result.(*NotOp)
		assert.True(t, typeOk)
		binOp, typeOk := notOp.X.(*BinaryOp)
		assert.True(t, typeOk)

		// assert that the binary op is a conjunction
		assert.Equal(t, AndOp, binOp.Op)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("(A & B) contracts to !(!A | !B)", func(t *testing.T) {
		f := And(Var("A"), Var("B"))

		// assert that DeMorgan's law is applied
		result, ok := DeMorganContractEager(f)
		assert.True(t, ok)

		// assert that the result is a negated binary op
		notOp, typeOk := result.(*NotOp)
		assert.True(t, typeOk)
		binOp, typeOk := notOp.X.(*BinaryOp)
		assert.True(t, typeOk)

		// assert that the binary op is a disjunction
		assert.Equal(t, OrOp, binOp.Op)

		// assert that the result is semantically equivalent to the input
		assert.True(t, bf.IsEquiv(f, result))
	})
	t.Run("A -> B is not contracted", func(t *testing.T) {
		f := Implies(Var("A"), Var("B"))

		// assert that DeMorgan's law is not applied
		result, ok := DeMorganContractEager(f)
		assert.False(t, ok)

		// assert that the result is the same as the input
		assert.Equal(t, f, result)
	})
}
