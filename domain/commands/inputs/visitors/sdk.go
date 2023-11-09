package visitors

import (
	"github.com/steve-care-software/steve/domain/commands/inputs/visitors/administrators"
)

// Visitor represents visitor command
type Visitor interface {
	IsAdministrator() bool
	Administrator() administrators.Administrator
}
