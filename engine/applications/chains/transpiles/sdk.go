package transpiles

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/programs"
)

// Application represents the transpile application
type Application interface {
	Execute(source programs.Program, target grammars.Grammar, bridge transpiles.Transpile) ([]byte, error)
}
