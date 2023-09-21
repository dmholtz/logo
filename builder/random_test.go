package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomFormulaBuilder(t *testing.T) {
	t.Run("non-positive numVars panics", func(t *testing.T) {
		assert.Panics(t, func() { NewRandomFormulaBuilder(0) })
	})
}

func TestBuildRandom(t *testing.T) {
	t.Run("Build() panics if the number of operands is non-positive", func(t *testing.T) {
		builder := NewRandomFormulaBuilder(3)
		assert.Panics(t, func() { builder.Build(0) })
	})

	t.Run("Build() outputs a random formula", func(t *testing.T) {
		builder := NewRandomFormulaBuilder(3)
		_ = builder.Build(10)
	})
}
