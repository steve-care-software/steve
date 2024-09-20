package transpiles

import (
	"github.com/steve-care-software/steve/domain/pipelines/transpiles"
	"github.com/steve-care-software/steve/domain/programs"
	"github.com/steve-care-software/steve/domain/programs/grammars"
)

// Application represents the transpile application
type Application interface {
	Execute(source programs.Program, target grammars.Grammar, bridge transpiles.Transpile) ([]byte, error)
}
