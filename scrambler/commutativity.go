package scrambler

import (
	"math/rand"

	. "github.com/dmholtz/logo"
)

func Commute(f LogicNode) (LogicNode, bool) {
	switch f1 := f.(type) {
	case *BinaryOp:
		// only AND, OR, and IFF are commutative
		if f1.Op == AndOp || f1.Op == OrOp || f1.Op == IffOp {
			f1.X, f1.Y = f1.Y, f1.X
			return f1, true
		}
	case *NaryOp:
		// only AND and OR are commutative
		if f1.Op == AndOp || f1.Op == OrOp {
			operands := make([]LogicNode, len(f1.Clauses))
			copy(operands, f1.Clauses)
			for i, j := range rand.Perm(len(operands)) {
				f1.Clauses[i] = operands[j]
			}
			return f1, true
		}
	}
	return f, false
}
