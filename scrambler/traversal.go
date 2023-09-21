package scrambler

import (
	"fmt"

	"math/rand"

	. "github.com/dmholtz/logo"
)

// Traverse() traverses the formula tree and applies the transform function to each node
func Traverse(f LogicNode, transform func(LogicNode) (LogicNode, bool)) LogicNode {
	// apply the transform function to the current node
	f, _ = transform(f)

	// traverse the node's children recursively
	switch f1 := f.(type) {
	case *NotOp:
		f1.X = Traverse(f1.X, transform)
		return f1
	case *BinaryOp:
		f1.X = Traverse(f1.X, transform)
		f1.Y = Traverse(f1.Y, transform)
		return f1
	case *Variable:
		return f
	case Leaf:
		return f
	case *NaryOp:
		for i, c := range f1.Clauses {
			f1.Clauses[i] = Traverse(c, transform)
		}
		return f1
	default:
		panic(fmt.Sprintf("Unkown type=%T of subformula=%s", f1, f1))
	}
}

// TraverseProbabilistic() traverses the formula tree and applies the transform function to each node with a given probability
func TraverseProbabilistic(f LogicNode, transform func(LogicNode) (LogicNode, bool), probability float64) LogicNode {
	if probability < 0 || probability > 1 {
		panic(fmt.Sprintf("Probability must be between 0 and 1, but is %f", probability))
	}

	// apply the transform function to the current node with the given probability
	if rand.Float64() < probability {
		f, _ = transform(f)
	}

	switch f1 := f.(type) {
	case *NotOp:
		f1.X = TraverseProbabilistic(f1.X, transform, probability)
		return f1
	case *BinaryOp:
		f1.X = TraverseProbabilistic(f1.X, transform, probability)
		f1.Y = TraverseProbabilistic(f1.Y, transform, probability)
		return f1
	case *Variable:
		return f
	case Leaf:
		return f
	case *NaryOp:
		for i, c := range f1.Clauses {
			f1.Clauses[i] = TraverseProbabilistic(c, transform, probability)
		}
		return f1
	default:
		panic(fmt.Sprintf("Unkown type=%T of subformula=%s", f1, f1))
	}
}
