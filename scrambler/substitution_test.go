package scrambler

import (
	"testing"

	. "github.com/dmholtz/logo"
	bf "github.com/dmholtz/logo/brute_force"
	"github.com/stretchr/testify/assert"
)

func TestSubstituteByImplies(t *testing.T) {
	t.Run("!A OR B is substituted by A->B", func(t *testing.T) {
		f := BinaryOp{X: Not(Var("A")), Y: Var("B"), Op: OrOp}

		// assert that the substitution is successful
		result, subOk := SubstituteByImplies(&f)
		assert.True(t, subOk)

		// assert that the result is an implies node
		impliesNode, typeOk := result.(*BinaryOp)
		assert.True(t, typeOk)
		assert.Equal(t, impliesNode.Op, IfOp)

		// assert that the implies node has the correct subnodes
		assert.Equal(t, f.X, Not(impliesNode.X))
		assert.Equal(t, f.Y, impliesNode.Y)

		// assert that result yields a semantically equivalent formula
		assert.True(t, bf.IsEquiv(&f, result))
	})
	t.Run("B OR !A is substituted by A->B", func(t *testing.T) {
		f := BinaryOp{X: Var("B"), Y: Not(Var("A")), Op: OrOp}

		// assert that the substitution is successful
		result, subOk := SubstituteByImplies(&f)
		assert.True(t, subOk)

		// assert that the result is an implies node
		impliesNode, typeOk := result.(*BinaryOp)
		assert.True(t, typeOk)
		assert.Equal(t, impliesNode.Op, IfOp)

		// assert that the implies node has the correct subnodes
		assert.Equal(t, f.X, impliesNode.Y)
		assert.Equal(t, f.Y, Not(impliesNode.X))

		// assert that result yields a semantically equivalent formula
		assert.True(t, bf.IsEquiv(&f, result))
	})
	t.Run("A AND !B is not substituted", func(t *testing.T) {
		f := BinaryOp{X: Var("A"), Y: Not(Var("B")), Op: AndOp}

		// assert that the substitution is not successful
		result, subOk := SubstituteByImplies(&f)
		assert.False(t, subOk)

		// assert that the result is the same as the input
		assert.Equal(t, &f, result)
	})
}

func TestSubstituteByIff(t *testing.T) {
	t.Run("(A->B) AND (B->A) is substituted by (A<->B)", func(t *testing.T) {
		f := BinaryOp{
			X:  &BinaryOp{X: Var("A"), Y: Var("B"), Op: IfOp},
			Y:  &BinaryOp{X: Var("B"), Y: Var("A"), Op: IfOp},
			Op: AndOp,
		}

		// assert that the substitution is successful
		result, subOk := SubstituteByIff(&f)
		assert.True(t, subOk)

		// assert that the result is an iff node
		iffNode, typeOk := result.(*BinaryOp)
		assert.True(t, typeOk)
		assert.Equal(t, iffNode.Op, IffOp)

		// assert that the iff node has the correct subnodes
		fLeft := f.X.(*BinaryOp)
		assert.Equal(t, fLeft.X, iffNode.X)
		assert.Equal(t, fLeft.Y, iffNode.Y)

		// assert that result yields a semantically equivalent formula
		assert.True(t, bf.IsEquiv(&f, result))
	})
	t.Run("A AND B is not substituted", func(t *testing.T) {
		f := BinaryOp{X: Var("A"), Y: Var("B"), Op: AndOp}

		// assert that the substitution is not successful
		result, subOk := SubstituteByIff(&f)
		assert.False(t, subOk)

		// assert that the result is the same as the input
		assert.Equal(t, &f, result)
	})
}

func TestRemoveImplies(t *testing.T) {
	t.Run("A->B is substituted by !A OR B", func(t *testing.T) {
		f := BinaryOp{X: Var("A"), Y: Var("B"), Op: IfOp}

		// assert that the substitution is successful
		result, subOk := RemoveImplies(&f)
		assert.True(t, subOk)

		// assert that the result is an implies node
		orNode, typeOk := result.(*BinaryOp)
		assert.True(t, typeOk)
		assert.Equal(t, orNode.Op, OrOp)

		// assert that the implies node has the correct subnodes
		assert.Equal(t, Not(f.X), orNode.X)
		assert.Equal(t, f.Y, orNode.Y)

		// assert that result yields a semantically equivalent formula
		assert.True(t, bf.IsEquiv(&f, result))
	})
	t.Run("A OR B is not substituted", func(t *testing.T) {
		f := BinaryOp{X: Var("A"), Y: Var("B"), Op: OrOp}

		// assert that the substitution is not successful
		result, subOk := RemoveImplies(&f)
		assert.False(t, subOk)

		// assert that the result is the same as the input
		assert.Equal(t, &f, result)
	})
}

func TestRemoveIff(t *testing.T) {
	t.Run("A<->B is substituted by ((!A AND !B) OR (A AND B))", func(t *testing.T) {
		f := BinaryOp{X: Var("A"), Y: Var("B"), Op: IffOp}

		// assert that the substitution is successful
		result, subOk := RemoveIff(&f)
		assert.True(t, subOk)

		// assert that the result is an implies node
		orNode, typeOk := result.(*BinaryOp)
		assert.True(t, typeOk)
		assert.Equal(t, orNode.Op, OrOp)

		// assert that the implies node has the correct subnodes
		andNode1, typeOk := orNode.X.(*BinaryOp)
		assert.True(t, typeOk)
		assert.Equal(t, andNode1.Op, AndOp)
		assert.Equal(t, Not(f.X), andNode1.X)
		assert.Equal(t, Not(f.Y), andNode1.Y)

		andNode2, typeOk := orNode.Y.(*BinaryOp)
		assert.True(t, typeOk)
		assert.Equal(t, andNode2.Op, AndOp)
		assert.Equal(t, f.X, andNode2.X)
		assert.Equal(t, f.Y, andNode2.Y)

		// assert that result yields a semantically equivalent formula
		assert.True(t, bf.IsEquiv(&f, result))
	})
	t.Run("A->B is not substituted", func(t *testing.T) {
		f := BinaryOp{X: Var("A"), Y: Var("B"), Op: IfOp}

		// assert that the substitution is not successful
		result, subOk := RemoveIff(&f)
		assert.False(t, subOk)

		// assert that the result is the same as the input
		assert.Equal(t, &f, result)
	})
}
