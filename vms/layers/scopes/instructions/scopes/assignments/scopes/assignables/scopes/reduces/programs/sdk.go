package programs

import (
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
)

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithBytes(bytes bytes_programs.Program) Builder
	WithLength(length uint) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	Bytes() bytes_programs.Program
	Length() uint
}
