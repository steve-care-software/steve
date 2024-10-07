package scripts

import (
	"errors"

	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access/permissions"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access/writes"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections"
	connection_headers "github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/headers"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/headers/names"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/headers/names/cardinalities"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/links"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/links/references"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/links/references/externals"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/suites"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas/connections/suites/expectations"
	applications_parser "github.com/steve-care-software/steve/parsers/applications"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
)

type adapterBuilder struct {
	parserAppBuilder        applications_parser.Builder
	builder                 Builder
	headBuilder             heads.Builder
	accessBuilder           access.Builder
	permissionsBuilder      permissions.Builder
	permissionBuilder       permissions.PermissionBuilder
	writeBuilder            writes.Builder
	connectionsBuilder      connections.Builder
	connectionBuilder       connections.ConnectionBuilder
	suitesBuilder           suites.Builder
	suiteBuilder            suites.SuiteBuilder
	expectationsBuilder     expectations.Builder
	expectationBuilder      expectations.ExpectationBuilder
	linksBuilder            links.Builder
	linkBuilder             links.LinkBuilder
	referencesBuilder       references.Builder
	referenceBuilder        references.ReferenceBuilder
	externalBuilder         externals.Builder
	connectionHeaderBuilder connection_headers.Builder
	nameBuilder             names.Builder
	cardinalityBuilder      cardinalities.Builder
	grammar                 grammars.Grammar
}

func createAdapterBuilder(
	parserAppBuilder applications_parser.Builder,
	builder Builder,
	headBuilder heads.Builder,
	accessBuilder access.Builder,
	permissionsBuilder permissions.Builder,
	permissionBuilder permissions.PermissionBuilder,
	writeBuilder writes.Builder,
	connectionsBuilder connections.Builder,
	connectionBuilder connections.ConnectionBuilder,
	suitesBuilder suites.Builder,
	suiteBuilder suites.SuiteBuilder,
	expectationsBuilder expectations.Builder,
	expectationBuilder expectations.ExpectationBuilder,
	linksBuilder links.Builder,
	linkBuilder links.LinkBuilder,
	referencesBuilder references.Builder,
	referenceBuilder references.ReferenceBuilder,
	externalBuilder externals.Builder,
	connectionHeaderBuilder connection_headers.Builder,
	nameBuilder names.Builder,
	cardinalityBuilder cardinalities.Builder,
) AdapterBuilder {
	out := adapterBuilder{
		parserAppBuilder:        parserAppBuilder,
		builder:                 builder,
		headBuilder:             headBuilder,
		accessBuilder:           accessBuilder,
		permissionsBuilder:      permissionsBuilder,
		permissionBuilder:       permissionBuilder,
		writeBuilder:            writeBuilder,
		connectionsBuilder:      connectionsBuilder,
		connectionBuilder:       connectionBuilder,
		suitesBuilder:           suitesBuilder,
		suiteBuilder:            suiteBuilder,
		expectationsBuilder:     expectationsBuilder,
		expectationBuilder:      expectationBuilder,
		linksBuilder:            linksBuilder,
		linkBuilder:             linkBuilder,
		referencesBuilder:       referencesBuilder,
		referenceBuilder:        referenceBuilder,
		externalBuilder:         externalBuilder,
		connectionHeaderBuilder: connectionHeaderBuilder,
		nameBuilder:             nameBuilder,
		cardinalityBuilder:      cardinalityBuilder,
		grammar:                 nil,
	}

	return &out
}

// Create initializes the builder
func (app *adapterBuilder) Create() AdapterBuilder {
	return createAdapterBuilder(
		app.parserAppBuilder,
		app.builder,
		app.headBuilder,
		app.accessBuilder,
		app.permissionsBuilder,
		app.permissionBuilder,
		app.writeBuilder,
		app.connectionsBuilder,
		app.connectionBuilder,
		app.suitesBuilder,
		app.suiteBuilder,
		app.expectationsBuilder,
		app.expectationBuilder,
		app.linksBuilder,
		app.linkBuilder,
		app.referencesBuilder,
		app.referenceBuilder,
		app.externalBuilder,
		app.connectionHeaderBuilder,
		app.nameBuilder,
		app.cardinalityBuilder,
	)
}

// WithGramar adds a grammar to the builder
func (app *adapterBuilder) WithGramar(gramar grammars.Grammar) AdapterBuilder {
	app.grammar = gramar
	return app
}

// Now builds a new Adapter instance
func (app *adapterBuilder) Now() (Adapter, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build an Adapter instance")
	}

	return createAdapter(
		app.parserAppBuilder,
		app.grammar,
		app.builder,
		app.headBuilder,
		app.accessBuilder,
		app.permissionsBuilder,
		app.permissionBuilder,
		app.writeBuilder,
		app.connectionsBuilder,
		app.connectionBuilder,
		app.suitesBuilder,
		app.suiteBuilder,
		app.expectationsBuilder,
		app.expectationBuilder,
		app.linksBuilder,
		app.linkBuilder,
		app.referencesBuilder,
		app.referenceBuilder,
		app.externalBuilder,
		app.connectionHeaderBuilder,
		app.nameBuilder,
		app.cardinalityBuilder,
	), nil
}
