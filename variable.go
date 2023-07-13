package logo

import "fmt"

// Variable represents a propositional variable.
type Variable struct {
	Name string
}

func Var(name string) LogicNode {
	return &Variable{Name: name}
}

func (v Variable) Eval(assignment Assignment) bool {
	val, ok := assignment[v.Name]
	if !ok {
		panic(fmt.Sprintf("Variable=%s not in scope of assignment=%v", v.Name, assignment))
	}
	return val
}

func (v Variable) Scope() map[string]struct{} {
	return map[string]struct{}{v.Name: {}}
}

func (v Variable) String() string {
	return v.Name
}
