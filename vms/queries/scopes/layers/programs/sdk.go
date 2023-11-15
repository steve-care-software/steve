package programs

import (
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
	instruction_programs "github.com/steve-care-software/steve/vms/queries/scopes/layers/scopes/instructions/programs"
)

// Builder represents a programs builder
type Builder interface {
	Create() Builder
	WithList(list []Program) Builder
	Now() (Programs, error)
}

// Programs represents programs
type Programs interface {
	Hash() hash.Hash
	List() []Program
}

// ProgramBuilder represents a program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithInstructions(instructions instruction_programs.Programs) ProgramBuilder
	WithSignature(signature Signature) ProgramBuilder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Instructions() instruction_programs.Programs
	Signature() Signature
}

// SignatureBuilder represents a builder
type SignatureBuilder interface {
	Create() SignatureBuilder
	WithBytes(bytes string) SignatureBuilder
	WithParams(params []string) SignatureBuilder
	WithReturns(ret Kind) SignatureBuilder
	WithDependencies(dependencies []string) SignatureBuilder
	Now() (Signature, error)
}

// Signature represents the layer signature
type Signature interface {
	Hash() hash.Hash
	Bytes() string
	Params() []string
	Returns() Kind
	HasDependencies() bool
	Dependencies() []string
}

// KindBuilder represents the kind builder
type KindBuilder interface {
	Create() KindBuilder
	WithExecute(exec []string) KindBuilder
	IsContinue() bool
	IsPrompt() bool
	Now() (Kind, error)
}

// Kind represents the kind
type Kind interface {
	Hash() hash.Hash
	IsContinue() bool
	IsPrompt() bool
	IsExecute() bool
	Execute() []string
}
