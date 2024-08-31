package programs

import (
	"github.com/steve-care-software/steve/domain/data/points/programs/grammars"
)

// Program represents the program
type Program interface {
	Name() string
	Description() string
	Grammar() grammars.Grammar
	Code() []byte
}
