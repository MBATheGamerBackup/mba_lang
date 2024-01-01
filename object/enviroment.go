package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	var env = NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	var store = make(map[string]Object)
	return &Environment{
		store: store,
		outer: nil,
	}
}

func (e Environment) Get(name string) (Object, bool) {
	var obj, ok = e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
