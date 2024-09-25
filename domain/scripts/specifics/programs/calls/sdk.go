package calls

// Call represents a call
type Call interface {
	IsProgram() bool
	Program() ProgramCall
	IsFunc() bool
	Func() FuncCall
}

// ProgramCall represents a program call
type ProgramCall interface {
	Name() string
	Input() string
}

// FuncCall represents a func call
type FuncCall interface {
	Name() string
	Parameters() FuncCallParameters
	IsEngine() bool
}

// FuncCallParameters represents a func call parameter
type FuncCallParameters interface {
	List() []FuncCallParameter
}

// FuncCallParameter represents a func call parameter
type FuncCallParameter interface {
	Current() string
	Local() string
}
