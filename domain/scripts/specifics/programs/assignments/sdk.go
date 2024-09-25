package assignments

import (
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/operations"
)

// Assignment represents an assignment
type Assignment interface {
	Variables() []string
	Operation() operations.Operation
	IsInitial() bool
}
