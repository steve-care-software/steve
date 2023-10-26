package queries

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/messages"
)

// Builder represents a query builder
type Builder interface {
	Create() Builder
	WithMessage(message messages.Message) Builder
	WithLayer(layer layers.Layer) Builder
	WithParams(params symbols.Symbols) Builder
	Now() (Query, error)
}

// Query represents the query
type Query interface {
	Message() messages.Message
	Layer() layers.Layer
	HasParams() bool
	Params() symbols.Symbols
}
