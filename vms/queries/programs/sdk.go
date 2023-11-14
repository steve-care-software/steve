package programs

import (
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
)

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithLayer(layer string) Builder
	WithBytes(bytes bytes_programs.Program) Builder
	WithParams(params bytes_programs.Programs) Builder
	WithDependencies(dependencies []string) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Layer() string
	Bytes() bytes_programs.Program
	Params() bytes_programs.Programs
	HasDependencies() bool
	Dependencies() []string
}
