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

type adapterFactory struct {
	parserAppBuilder        applications_parser.Builder
	grammarAdapter          grammars.Adapter
	builder                 Builder
	schemaBuilder           schemas.Builder
	headBuilder             heads.Builder
	accessBuilder           access.Builder
	permissionBuilder       permissions.Builder
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
	grammar                 []byte
}

func createAdapterFactory(
	parserAppBuilder applications_parser.Builder,
	grammarAdapter grammars.Adapter,
	builder Builder,
	schemaBuilder schemas.Builder,
	headBuilder heads.Builder,
	accessBuilder access.Builder,
	permissionBuilder permissions.Builder,
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
	grammar []byte,
) AdapterFactory {
	out := adapterFactory{
		parserAppBuilder:        parserAppBuilder,
		grammarAdapter:          grammarAdapter,
		builder:                 builder,
		schemaBuilder:           schemaBuilder,
		headBuilder:             headBuilder,
		accessBuilder:           accessBuilder,
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
		grammar:                 grammar,
	}

	return &out
}

// Create creates a new adapter
func (app *adapterFactory) Create() (Adapter, error) {
	grammar, _, err := app.grammarAdapter.ToGrammar(app.grammar)
	if err != nil {
		return nil, err
	}

	return createAdapter(
		app.parserAppBuilder,
		grammar,
		app.builder,
		app.schemaBuilder,
		app.headBuilder,
		app.accessBuilder,
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
