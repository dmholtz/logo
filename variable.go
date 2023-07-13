package logo

import "fmt"

type Variable struct {
	Name string
}

func (lv Variable) Eval(assignment Assignment) bool {
	val, ok := assignment[lv.Name]
	if !ok {
		panic(fmt.Sprintf("Variable=%s not in scope of assignment=%v", lv.Name, assignment))
	}
	return val
}

func (v Variable) Scope() map[string]struct{} {
	return map[string]struct{}{v.Name: {}}
}

func (lv Variable) String() string {
	return lv.Name
}
