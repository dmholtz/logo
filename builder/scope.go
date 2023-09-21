package builder

import "fmt"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// BuildScope returns a slice of strings of length numVariables
// If numVariables is greater than 26, the strings will be of the form "x1", "x2", etc.
// If numVariables is less than or equal to 26, the strings will be of the form "A", "B", etc.
func BuildScope(numVariables int) []string {
	if numVariables < 1 {
		panic("numVariables must be positive")
	}
	scope := make([]string, numVariables)
	for i := 0; i < numVariables; i++ {
		if numVariables <= len(alphabet) {
			scope[i] = string(alphabet[i])
		} else {
			scope[i] = fmt.Sprintf("x%d", i+1)
		}
	}
	return scope
}
