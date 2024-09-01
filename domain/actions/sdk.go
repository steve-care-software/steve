package actions

// Chain represents a chain of action
type Chain interface {
	Input() []byte // contain my grammar code
	Program() Program
}

// Program represents the program
type Program interface {
	Input() []byte // contains my program code
	Action() Action
}

// Action represents a program action
type Action interface {
	IsInterpret() bool
	IsTranspile() bool
	Transpile() Transpile
}

// Transpile represents a transpile
type Transpile interface {
	To()     // grammar code
	Bridge() // bridge code
	HasNext() bool
	Next() Chain
}
