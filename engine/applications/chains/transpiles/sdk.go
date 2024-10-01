package transpiles

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/transpiles"
	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
)

// Application represents the transpile application
type Application interface {
	Execute(source asts.AST, target grammars.Grammar, bridge transpiles.Transpile) ([]byte, error)
}
