package inputs

import (
	"github.com/steve-care-software/steve/domain/commands/inputs/administrators"
	"github.com/steve-care-software/steve/domain/commands/inputs/visitors"
)

// Adapter represents a visitor's adapter
type Adapter interface {
	ToInput(bytes []byte) (Input, error)
	ToBytes(ins Input) ([]byte, error)
}

// Input represents an input
type Input interface {
	IsAdministrator() bool
	Administrator() administrators.Administrator
	IsVisitor() bool
	Visitor() visitors.Visitor
}
