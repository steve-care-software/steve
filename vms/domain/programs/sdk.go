package programs

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithInstructions(instructions Instructions) Builder
	WithParameters(parameters []string) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Instructions() Instructions
	HasParameters() bool
	Parameters() []string
}

// InstructionsBuilder represents the instructions builder
type InstructionsBuilder interface {
	Create() InstructionsBuilder
	WithList(list []Instruction) InstructionsBuilder
	Now() (Instructions, error)
}

// Instructions represents instructions
type Instructions interface {
	List() []Instruction
}

// InstructionBuilder represents the instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithInit(init Init) InstructionBuilder
	WithDelete(del string) InstructionBuilder
	WithBack(back string) InstructionBuilder
	WithCommit(commit Commit) InstructionBuilder
	WithClear(clear string) InstructionBuilder
	WithRollback(rollback string) InstructionBuilder
	WithAssignment(assignment Assignment) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsInit() bool
	Init() Init
	IsDelete() bool
	Delete() string
	IsBack() bool
	Back() string
	IsCommit() bool
	Commit() Commit
	IsClear() bool
	Clear() string
	IsRollback() bool
	Rollback() string
	IsAssignment() bool
	Assignment() Assignment
}

// InitBuilder represents the init builder
type InitBuilder interface {
	Create() InitBuilder
	WithRoot(root Root) InitBuilder
	WithPath(path string) InitBuilder
	Now() (Init, error)
}

// Init represents an init instruction
type Init interface {
	Root() Root
	Path() string
}

// RootBuilder represents the root builder
type RootBuilder interface {
	Create() RootBuilder
	WithFees(fees uint16) RootBuilder
	WithAffiliate(affiliate uint16) RootBuilder
	Now() (Root, error)
}

// Root represents a root instruction
type Root interface {
	Fees() uint16
	Affiliate() uint16
}

// CommitBuilder represents the commit builder
type CommitBuilder interface {
	Create() CommitBuilder
	WithContext(context string) CommitBuilder
	WithMessage(message string) CommitBuilder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Context() string
	Message() string
}

// AssignmentBuilder represents the assignment builder
type AssignmentBuilder interface {
	Create() AssignmentBuilder
	WithName(name string) AssignmentBuilder
	WithAssignable(assignable Assignable) AssignmentBuilder
	Now() (Assignment, error)
}

// Assignment represents an assignmnet
type Assignment interface {
	Name() string
	Assignable() Assignable
}

// AssignableBuilder represents the assignable builder
type AssignableBuilder interface {
	Create() AssignableBuilder
	WithBegin(begin string) AssignableBuilder
	WithExists(exists string) AssignableBuilder
	WithTransact(trx Transact) AssignableBuilder
	WithQueue(queue string) AssignableBuilder
	Now() (Assignable, error)
}

// Assignable represents an assignable
type Assignable interface {
	IsBegin() bool
	Begin() string
	IsExists() bool
	Exists() string
	IsTransact() bool
	Transact() Transact
	IsQueue() bool
	Queue() string
}

// TransactBuilder represents a transact builder
type TransactBuilder interface {
	Create() TransactBuilder
	WithContext(context string) TransactBuilder
	WithInput(input []byte) TransactBuilder
	Now() (Transact, error)
}

// Transact represents a transact
type Transact interface {
	Context() string
	Input() []byte
}
