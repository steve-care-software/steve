package initializations

import (
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/assignments"
	"github.com/steve-care-software/steve/domain/scripts/specifics/programs/containers"
)

// Initialization represents a variable initialization
type Initialization interface {
	Container() containers.Container
	Assignment() assignments.Assignment
}
