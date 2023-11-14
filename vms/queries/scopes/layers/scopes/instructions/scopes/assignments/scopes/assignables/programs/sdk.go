package programs

import (
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
)

// ProgramBuilder represents the program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	/*WithQuery(query Query) AssignableBuilder
	WithReduce(reduce reduces.Reduce) AssignableBuilder*/
	WithCompare(compare bytes_programs.Programs) ProgramBuilder
	/*WithLength(length bytes_programs.Programs) AssignableBuilder
	WithJoin(join bytes_programs.Programs) AssignableBuilder
	WithValue(value bytes_programs.Programs) AssignableBuilder*/
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	/*IsQuery() bool
	Query() Query
	IsReduce() bool
	Reduce() reduces.Reduce*/
	IsCompare() bool
	Compare() bytes_programs.Programs
	/*IsLength() bool
	Length() bytes_programs.Programs
	IsJoin() bool
	Join() bytes_programs.Programs
	IsValue() bool
	Value() bytes_programs.Programs*/
}
