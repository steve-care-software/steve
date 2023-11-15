package programs

import (
	bytes_programs "github.com/steve-care-software/steve/vms/bytes/programs"
	"github.com/steve-care-software/steve/vms/bytes/results/hash"
	reduce_programs "github.com/steve-care-software/steve/vms/layers/scopes/instructions/scopes/assignments/scopes/assignables/scopes/reduces/programs"
	query_programs "github.com/steve-care-software/steve/vms/queries/programs"
)

// ProgramBuilder represents the program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithQuery(query query_programs.Program) ProgramBuilder
	WithCompare(compare bytes_programs.Programs) ProgramBuilder
	WithLength(length bytes_programs.Programs) ProgramBuilder
	WithJoin(join bytes_programs.Programs) ProgramBuilder
	WithReduce(reduce reduce_programs.Program) ProgramBuilder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	Hash() hash.Hash
	IsQuery() bool
	Query() query_programs.Program
	IsCompare() bool
	Compare() bytes_programs.Programs
	IsLength() bool
	Length() bytes_programs.Programs
	IsJoin() bool
	Join() bytes_programs.Programs
	IsReduce() bool
	Reduce() reduce_programs.Program
}
