package logo

import "fmt"

type And struct {
	X, Y LogicNode
}

func (a And) Eval(assignment Assignment) bool {
	return a.X.Eval(assignment) && a.Y.Eval(assignment)
}

func (a And) Scope() map[string]struct{} {
	scope := make(map[string]struct{})
	for k, v := range a.X.Scope() {
		scope[k] = v
	}
	for k, v := range a.Y.Scope() {
		scope[k] = v
	}
	return scope
}

func (a And) String() string {
	return fmt.Sprintf("(%s & %s)", a.X.String(), a.Y.String())
}

type Or struct {
	X, Y LogicNode
}

func (o Or) Eval(assignment Assignment) bool {
	return o.X.Eval(assignment) || o.Y.Eval(assignment)
}

func (o Or) Scope() map[string]struct{} {
	scope := make(map[string]struct{})
	for k, v := range o.X.Scope() {
		scope[k] = v
	}
	for k, v := range o.Y.Scope() {
		scope[k] = v
	}
	return scope
}

func (o Or) String() string {
	return fmt.Sprintf("(%s | %s)", o.X.String(), o.Y.String())
}

type Implication struct {
	X, Y LogicNode
}

func (i Implication) Eval(assignment Assignment) bool {
	return !i.X.Eval(assignment) || i.Y.Eval(assignment)
}

func (i Implication) Scope() map[string]struct{} {
	scope := make(map[string]struct{})
	for k, v := range i.X.Scope() {
		scope[k] = v
	}
	for k, v := range i.Y.Scope() {
		scope[k] = v
	}
	return scope
}

func (i Implication) String() string {
	return fmt.Sprintf("(%s -> %s)", i.X.String(), i.Y.String())
}

type Equivalence struct {
	X, Y LogicNode
}

func (e Equivalence) Eval(assignment Assignment) bool {
	return (!e.X.Eval(assignment) && !e.Y.Eval(assignment)) || (e.X.Eval(assignment) && e.Y.Eval(assignment))
}

func (e Equivalence) Scope() map[string]struct{} {
	scope := make(map[string]struct{})
	for k, v := range e.X.Scope() {
		scope[k] = v
	}
	for k, v := range e.Y.Scope() {
		scope[k] = v
	}
	return scope
}

func (e Equivalence) String() string {
	return fmt.Sprintf("(%s <-> %s)", e.X.String(), e.Y.String())
}
