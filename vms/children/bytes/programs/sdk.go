package programs

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithList(list []Program) Builder
	Now() (Programs, error)
}

// Programs represents programs
type Programs interface {
	List() []Program
}

// ProgramBuilder represents a program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithVariable(variable string) ProgramBuilder
	WithValue(value []byte) ProgramBuilder
	Now() (Program, error)
}

// Program represents a bytes program
type Program interface {
	IsVariable() bool
	Variable() string
	IsValue() bool
	Value() []byte
}
