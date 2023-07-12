package logo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeaf(t *testing.T) {
	t.Run("top returns true", func(t *testing.T) {
		assert.Equal(t, Top().Eval(nil), true)
	})
	t.Run("bottom returns false", func(t *testing.T) {
		assert.Equal(t, Bottom().Eval(nil), false)
	})
}
