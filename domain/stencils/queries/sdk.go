package queries

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/messages"
)

// Query represents the query
type Query interface {
	Message() messages.Message
	Layer() layers.Layer
	HasParams() bool
	Params() symbols.Symbols
}
