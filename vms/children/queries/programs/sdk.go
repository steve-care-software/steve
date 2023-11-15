package programs

import (
	bytes_programs "github.com/steve-care-software/steve/vms/children/bytes/programs"
	"github.com/steve-care-software/steve/vms/libraries/hash"
)

// Builder represents the program builder
type Builder interface {
	Create() Builder
	WithLayer(layer string) Builder
	WithBytes(bytes bytes_programs.Program) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Layer() string
	Bytes() bytes_programs.Program
}

// SuitesBuilder represents a suites builder
type SuitesBuilder interface {
	Create() SuitesBuilder
	WithList(list []Suite) SuitesBuilder
	Now() (Suites, error)
}

// Suites represents suites
type Suites interface {
	Hash() hash.Hash
	List() []Suite
}

// SuiteBuilder represents a suite builder
type SuiteBuilder interface {
	Create() SuiteBuilder
	WithName(name string) SuiteBuilder
	WithQuery(query Program) SuiteBuilder
	WithValid(valid []byte) SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Hash() hash.Hash
	Name() string
	Query() Program
	IsValid() bool
	Valid() []byte
}
