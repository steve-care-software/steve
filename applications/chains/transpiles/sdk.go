package transpiles

import (
	"github.com/steve-care-software/steve/domain/programs"
	"github.com/steve-care-software/steve/domain/programs/grammars"
	"github.com/steve-care-software/steve/domain/scripts/specifics/transpiles"
)

// Application represents the transpile application
type Application interface {
	Execute(source programs.Program, target grammars.Grammar, bridge transpiles.Transpile) ([]byte, error)
}
