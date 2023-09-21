package builder

import (
	"testing"

	bf "github.com/dmholtz/logo/brute_force"
	"github.com/stretchr/testify/assert"
)

func TestNewEquivalentFormulaBuilder(t *testing.T) {
	t.Run("panic if numVars is less than 2", func(t *testing.T) {
		assert.Panics(t, func() { NewEquivalentFormulaBuilder(1, 6) })
	})

	t.Run("panic if numOperators is less than 2", func(t *testing.T) {
		assert.Panics(t, func() { NewEquivalentFormulaBuilder(3, 1) })
	})

	t.Run("Base formula has the proper scope", func(t *testing.T) {
		builder := NewEquivalentFormulaBuilder(3, 6)

		assert.Equal(t, 2, len(builder.BaseFormula.Scope()))
	})

	t.Run("Appendix has the proper scope", func(t *testing.T) {
		builder := NewEquivalentFormulaBuilder(3, 6)

		assert.Equal(t, 1, len(builder.Appendix.Scope()))
		assert.NotContains(t, builder.BaseFormula.Scope(), builder.Appendix.Scope())

	})
}

func TestEquivalent(t *testing.T) {
	t.Run("Equivalent() returns a formula that is equivalent to the reference formula", func(t *testing.T) {
		builder := NewEquivalentFormulaBuilder(5, 6)
		reference := builder.Question()
		equivalent := builder.Equivalent()

		assert.True(t, bf.IsEquiv(reference, equivalent))
	})
}

func TestNotEquivalent(t *testing.T) {
	t.Run("NotEquivalent() returns a formula that is not equivalent to the reference formula", func(t *testing.T) {
		builder := NewEquivalentFormulaBuilder(5, 6)
		reference := builder.Question()
		notEquivalent := builder.NotEquivalent()

		assert.False(t, bf.IsEquiv(reference, notEquivalent))
	})
}
