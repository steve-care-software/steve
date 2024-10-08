package scripts

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access/permissions"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads/access/writes"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections"
	connection_headers "github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/headers"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/headers/names"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/headers/names/cardinalities"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references/externals"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/suites"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/suites/expectations"
	applications_parser "github.com/steve-care-software/steve/parsers/applications"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
)

// NewAdapterBuilder creates a new adapter builder
func NewAdapterBuilder() AdapterBuilder {
	parserAppBuilder := applications_parser.NewBuilder()
	builder := NewBuilder()
	schemaBuilder := schemas.NewBuilder()
	headBuilder := heads.NewBuilder()
	accessBuilder := access.NewBuilder()
	permissionsBuilder := permissions.NewBuilder()
	writeBuilder := writes.NewBuilder()
	connectionsBuilder := connections.NewBuilder()
	connectionBuilder := connections.NewConnectionBuilder()
	suitesBuilder := suites.NewBuilder()
	suiteBuilder := suites.NewSuiteBuilder()
	expectationsBuilder := expectations.NewBuilder()
	expectationBuilder := expectations.NewExpectationBuilder()
	linksBuilder := links.NewBuilder()
	linkBuilder := links.NewLinkBuilder()
	referencesBuilder := references.NewBuilder()
	referenceBuilder := references.NewReferenceBuilder()
	externalBuilder := externals.NewBuilder()
	connectionHeaderBuilder := connection_headers.NewBuilder()
	nameBuilder := names.NewBuilder()
	cardinalityBuilder := cardinalities.NewBuilder()
	return createAdapterBuilder(
		parserAppBuilder,
		builder,
		schemaBuilder,
		headBuilder,
		accessBuilder,
		permissionsBuilder,
		writeBuilder,
		connectionsBuilder,
		connectionBuilder,
		suitesBuilder,
		suiteBuilder,
		expectationsBuilder,
		expectationBuilder,
		linksBuilder,
		linkBuilder,
		referencesBuilder,
		referenceBuilder,
		externalBuilder,
		connectionHeaderBuilder,
		nameBuilder,
		cardinalityBuilder,
	)
}

// NewBuilder creates a new script builder
func NewBuilder() Builder {
	return createBuilder()
}

// AdapterBuilder represents an adapter builder
type AdapterBuilder interface {
	Create() AdapterBuilder
	WithGramar(gramar grammars.Grammar) AdapterBuilder
	Now() (Adapter, error)
}

// Adapter represents the script adapter
type Adapter interface {
	ToScript(input []byte) (Script, []byte, error)
}

// Builder represents the script builder
type Builder interface {
	Create() Builder
	WithSchema(schema schemas.Schema) Builder
	Now() (Script, error)
}

// Script represents a script
type Script interface {
	IsSchema() bool
	Schema() schemas.Schema
}
