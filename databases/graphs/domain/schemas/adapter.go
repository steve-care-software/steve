package schemas

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections"
	connection_headers "github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/headers"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/headers/names"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/headers/names/cardinalities"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links/references"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/links/references/externals"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/suites"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/connections/suites/expectations"
	"github.com/steve-care-software/steve/databases/graphs/domain/schemas/headers"
	"github.com/steve-care-software/steve/parsers/domain/asts"
	"github.com/steve-care-software/steve/parsers/domain/asts/instructions"
	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/queries"
)

type adapter struct {
	astAdapter              asts.Adapter
	queryAdapter            queries.Adapter
	builder                 Builder
	headerBuilder           headers.Builder
	connectionsBuilder      connections.Builder
	connectionBuilder       connections.ConnectionBuilder
	suitesBuilder           suites.Builder
	suiteBuilder            suites.SuiteBuilder
	expectationBuilder      expectations.Builder
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

func createAdapter(
	astAdapter asts.Adapter,
	queryAdapter queries.Adapter,
	builder Builder,
	headerBuilder headers.Builder,
	connectionsBuilder connections.Builder,
	connectionBuilder connections.ConnectionBuilder,
	suitesBuilder suites.Builder,
	suiteBuilder suites.SuiteBuilder,
	expectationBuilder expectations.Builder,
	linksBuilder links.Builder,
	linkBuilder links.LinkBuilder,
	referencesBuilder references.Builder,
	referenceBuilder references.ReferenceBuilder,
	externalBuilder externals.Builder,
	connectionHeaderBuilder connection_headers.Builder,
	nameBuilder names.Builder,
	cardinalityBuilder cardinalities.Builder,
	grammar grammars.Grammar,
) Adapter {
	out := adapter{
		astAdapter:              astAdapter,
		queryAdapter:            queryAdapter,
		builder:                 builder,
		headerBuilder:           headerBuilder,
		connectionsBuilder:      connectionsBuilder,
		connectionBuilder:       connectionBuilder,
		suitesBuilder:           suitesBuilder,
		suiteBuilder:            suiteBuilder,
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

// ToSchema converts an input to a Schema instance
func (app *adapter) ToSchema(input []byte) (Schema, []byte, error) {
	ast, retRemaining, err := app.astAdapter.ToAST(app.grammar, input)
	if err != nil {
		return nil, nil, err
	}

	root := ast.Root()
	retSchema, err := app.schema(root)
	if err != nil {
		return nil, nil, err
	}

	return retSchema, retRemaining, nil
}

func (app *adapter) schema(element instructions.Element) (Schema, error) {
	tokens, err := app.elementToTokens(element)
	if err != nil {
		return nil, err
	}

	retHead, err := app.head(tokens)
	if err != nil {
		return nil, err
	}

	retPoints, err := app.points(tokens)
	if err != nil {
		return nil, err
	}

	retConnections, err := app.connections(tokens)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithHeader(retHead).
		WithPoints(retPoints).
		WithConnections(retConnections).
		Now()
}

func (app *adapter) head(tokens instructions.Tokens) (headers.Header, error) {
	retVersionBytes, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		head[0][0]->versionInstruction[0][0]->numbers;
	`))

	if err != nil {
		return nil, err
	}

	retNameBytes, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		head[0][0]->nameInstruction[0][0]->variableName;
	`))

	if err != nil {
		return nil, err
	}

	version, err := strconv.Atoi(string(retVersionBytes))
	if err != nil {
		return nil, err
	}

	return app.headerBuilder.Create().
		WithName(string(retNameBytes)).
		WithVersion(uint(version)).
		Now()
}

func (app *adapter) query(tokens instructions.Tokens, script []byte) ([]byte, error) {
	query, remaining, err := app.queryAdapter.ToQuery(script)
	if err != nil {
		return nil, err
	}

	if len(remaining) > 0 {
		str := fmt.Sprintf("the script (%s) contains a remaining (%s)", script, remaining)
		return nil, errors.New(str)
	}

	chain := query.Chain()
	retElement, err := tokens.Select(chain)
	if err != nil {
		return nil, err
	}

	return retElement.Value(), nil
}

func (app *adapter) points(tokens instructions.Tokens) ([]string, error) {
	fmt.Printf("\n%v\n", tokens)

	panic(errors.New("stop"))
	return nil, nil
}

func (app *adapter) connections(tokens instructions.Tokens) (connections.Connections, error) {
	return nil, nil
}

func (app *adapter) tokenToFirstInstructionTokens(token instructions.Token) (instructions.Tokens, error) {
	retElement, err := token.Elements().Fetch(0)
	if err != nil {
		return nil, err
	}

	return app.elementToTokens(retElement)
}

func (app *adapter) elementToTokens(element instructions.Element) (instructions.Tokens, error) {
	if element.IsConstant() {
		str := fmt.Sprintf("the element (name: %s) was expected to contain an Instruction", element.Name())
		return nil, errors.New(str)
	}

	return element.Instruction().Tokens(), nil
}
