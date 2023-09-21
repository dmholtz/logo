package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScope(t *testing.T) {
	t.Run("Scope returns a slice of letters if numVariables is less than or equal to 26", func(t *testing.T) {
		assert.Equal(t, []string{"A"}, BuildScope(1))

		assert.Equal(t, []string{"A", "B", "C"}, BuildScope(3))

		longScope := BuildScope(26)
		assert.Equal(t, 26, len(longScope))
		assert.Equal(t, "A", longScope[0])
		assert.Equal(t, "Z", longScope[25])
	})
	t.Run("Scope returns a slice of x1, x2, etc. if numVariables is greater than 26", func(t *testing.T) {
		scope := BuildScope(27)

		assert.Equal(t, 27, len(scope))
		assert.Equal(t, "x1", scope[0])
		assert.Equal(t, "x27", scope[26])
	})
	t.Run("Scope panics if numVariables is less than 1", func(t *testing.T) {
		assert.Panics(t, func() { BuildScope(0) })
	})
}
