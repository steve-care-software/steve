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

func createAdapter(
	astAdapter asts.Adapter,
	queryAdapter queries.Adapter,
	builder Builder,
	headerBuilder headers.Builder,
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
	retVersionElements, _, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		head[0][0]->versionInstruction[0][0]->numbers;
	`))

	if err != nil {
		return nil, err
	}

	retNameElements, _, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		head[0][0]->nameInstruction[0][0]->variableName;
	`))

	if err != nil {
		return nil, err
	}

	version, err := strconv.Atoi(string(retVersionElements.Value()))
	if err != nil {
		return nil, err
	}

	return app.headerBuilder.Create().
		WithName(string(retNameElements.Value())).
		WithVersion(uint(version)).
		Now()
}

func (app *adapter) points(tokens instructions.Tokens) ([]string, error) {
	retElements, _, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		instructionPoints[0][0]->instructionPoint;
	`))

	if err != nil {
		return nil, err
	}

	output := []string{}
	elementsList := retElements.List()
	for _, oneElement := range elementsList {
		retTokens, err := app.elementToTokens(oneElement)
		if err != nil {
			return nil, err
		}

		_, retElement, err := app.query(retTokens, []byte(`
			v1;
			name: mySelector;
			variableName[0][0];
		`))

		if err != nil {
			return nil, err
		}

		output = append(output, string(retElement.Value()))
	}

	return output, nil
}

func (app *adapter) connections(tokens instructions.Tokens) (connections.Connections, error) {
	retElements, _, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		connectionBlocks[0][0]->connectionBlock;
	`))

	if err != nil {
		return nil, err
	}

	output := []connections.Connection{}
	elementsList := retElements.List()
	for _, oneElement := range elementsList {
		retTokens, err := app.elementToTokens(oneElement)
		if err != nil {
			return nil, err
		}

		retConnection, err := app.connection(retTokens)
		if err != nil {
			return nil, err
		}

		output = append(output, retConnection)
	}

	return app.connectionsBuilder.Create().
		WithList(output).
		Now()
}

func (app *adapter) connection(tokens instructions.Tokens) (connections.Connection, error) {
	retHeader, err := app.header(tokens)
	if err != nil {
		return nil, err
	}

	retLinks, err := app.links(tokens)
	if err != nil {
		return nil, err
	}

	builder := app.connectionBuilder.Create().
		WithHeader(retHeader).
		WithLinks(retLinks)

	retSuites, err := app.suites(tokens)
	if err == nil {
		builder.WithSuites(retSuites)
	}

	return builder.Now()

}

func (app *adapter) suites(tokens instructions.Tokens) (suites.Suites, error) {
	retElements, _, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		pointSuitesBlock[0][0]->pointSuites[0][0]->pointSuite;
	`))

	if err != nil {
		return nil, err
	}

	output := []suites.Suite{}
	list := retElements.List()
	for _, oneElement := range list {
		retTokens, err := app.elementToTokens(oneElement)
		if err != nil {
			return nil, err
		}

		retLink, err := app.suite(retTokens)
		if err != nil {
			return nil, err
		}

		output = append(output, retLink)
	}

	return app.suitesBuilder.Create().
		WithList(output).
		Now()
}

func (app *adapter) suite(tokens instructions.Tokens) (suites.Suite, error) {
	_, retNameElement, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		variableName[0][0];
	`))

	if err != nil {
		return nil, err
	}

	retReferenceElements, _, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		pointReference;
	`))

	if err != nil {
		return nil, err
	}

	retLink, err := app.link(retReferenceElements)
	if err != nil {
		return nil, err
	}

	retExpectations, err := app.expectations(tokens)
	if err != nil {
		return nil, err
	}

	return app.suiteBuilder.Create().
		WithName(string(retNameElement.Value())).
		WithLink(retLink).
		WithExpectations(retExpectations).
		Now()
}

func (app *adapter) expectations(tokens instructions.Tokens) (expectations.Expectations, error) {
	retElements, _, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		suiteInstructions[0][0]->suiteOptionInstruction;
	`))

	if err != nil {
		return nil, err
	}

	output := []expectations.Expectation{}
	list := retElements.List()
	for _, oneElement := range list {
		retTokens, err := app.elementToTokens(oneElement)
		if err != nil {
			return nil, err
		}

		retExpectation, err := app.expectation(retTokens)
		if err != nil {
			return nil, err
		}

		output = append(output, retExpectation)
	}

	return app.expectationsBuilder.Create().
		WithList(output).
		Now()
}

func (app *adapter) expectation(tokens instructions.Tokens) (expectations.Expectation, error) {
	retFail, _, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		suiteOption[0][0]->suiteInvalidLink[0][0]->suiteReferencesInParenthesis[0][0]->pointReferences[0][0]->pointReference;
	`))

	builder := app.expectationBuilder.Create()
	if err == nil {
		retReferences, err := app.references(retFail)
		if err != nil {
			return nil, err
		}

		builder.IsFail().
			WithReferences(retReferences)
	}

	retSuccess, _, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		suiteOption[0][0]->suiteReferencesInParenthesis[0][0]->pointReferences[0][0]->pointReference;
	`))

	if err == nil {
		retReferences, err := app.references(retSuccess)
		if err != nil {
			return nil, err
		}

		builder.WithReferences(retReferences)
	}

	return builder.Now()
}

func (app *adapter) links(tokens instructions.Tokens) (links.Links, error) {
	retElements, _, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		links[0][0]->link[0][0]->pointReference;
	`))

	if err == nil {
		link, err := app.link(retElements)
		if err != nil {
			return nil, err
		}

		return app.linksBuilder.Create().WithList([]links.Link{
			link,
		}).Now()
	}

	retElements, _, err = app.query(tokens, []byte(`
		v1;
		name: mySelector;
		links[0][0]->pipeJourney;
	`))

	if err != nil {
		return nil, err
	}

	output := []links.Link{}
	list := retElements.List()
	for _, oneElement := range list {
		retTokens, err := app.elementToTokens(oneElement)
		if err != nil {
			return nil, err
		}

		retElements, _, err = app.query(retTokens, []byte(`
			v1;
			name: mySelector;
			links;
		`))

		if err != nil {
			return nil, err
		}

		retLink, err := app.link(retElements)
		if err != nil {
			return nil, err
		}

		output = append(output, retLink)
	}

	return app.linksBuilder.Create().
		WithList(output).
		Now()

}

func (app *adapter) link(elements instructions.Elements) (links.Link, error) {
	list := elements.List()
	if len(list) != 2 {
		str := fmt.Sprintf("%d references were expected in the link, %d provided", 2, len(list))
		return nil, errors.New(str)
	}

	builder := app.linkBuilder.Create()
	for idx, oneElement := range list {
		retTokens, err := app.elementToTokens(oneElement)
		if err != nil {
			return nil, err
		}

		retReference, err := app.reference(retTokens)
		if err != nil {
			return nil, err
		}

		if idx <= 0 {
			builder.WithOrigin(retReference)
			continue
		}

		builder.WithTarget(retReference)
	}

	return builder.Now()

}

func (app *adapter) references(elements instructions.Elements) (references.References, error) {
	output := []references.Reference{}
	list := elements.List()
	for _, oneElement := range list {
		retTokens, err := app.elementToTokens(oneElement)
		if err != nil {
			return nil, err
		}

		retReferemce, err := app.reference(retTokens)
		if err != nil {
			return nil, err
		}

		output = append(output, retReferemce)
	}

	return app.referencesBuilder.Create().
		WithList(output).
		Now()
}

func (app *adapter) reference(tokens instructions.Tokens) (references.Reference, error) {
	_, retElement, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		reference[0][0]->variableName[0][0];
	`))

	if err == nil {
		return app.referenceBuilder.Create().
			WithInternal(string(retElement.Name())).
			Now()
	}

	_, retSchemaElement, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		externalPointReference[0][0]->reference[0][0]->variableName[0][0];
	`))

	if err != nil {
		return nil, err
	}

	_, retPointElement, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		externalPointReference[0][0]->variableName[0][0];
	`))

	if err != nil {
		return nil, err
	}

	external, err := app.externalBuilder.Create().
		WithSchema(string(retSchemaElement.Value())).
		WithPoint(string(retPointElement.Value())).
		Now()

	if err != nil {
		return nil, err
	}

	return app.referenceBuilder.Create().
		WithExternal(external).
		Now()

}

func (app *adapter) header(tokens instructions.Tokens) (connection_headers.Header, error) {
	name, err := app.name(tokens)
	if err != nil {
		return nil, err
	}

	builder := app.connectionHeaderBuilder.Create().WithName(name)
	reverse, err := app.reverse(tokens)
	if err == nil {
		builder.WithReverse(reverse)
	}

	return builder.Now()
}

func (app *adapter) reverse(tokens instructions.Tokens) (names.Name, error) {
	_, retElement, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		connectionHeader[0][0]->variableNameInParenthesis[0][0]->nameWithCardinality[0][0];
	`))

	if err != nil {
		return nil, err
	}

	builder := app.nameBuilder.Create().
		WithName(string(retElement.Value()))

	retCardinality, err := app.cardinality(tokens)
	if err == nil {
		builder.WithCardinality(retCardinality)
	}

	return builder.Now()
}

func (app *adapter) name(tokens instructions.Tokens) (names.Name, error) {
	_, retElement, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		connectionHeader[0][0]->nameWithCardinality[0][0]->variableName[0][0];
	`))

	if err != nil {
		return nil, err
	}

	builder := app.nameBuilder.Create().
		WithName(string(retElement.Value()))

	retCardinality, err := app.cardinality(tokens)
	if err == nil {
		builder.WithCardinality(retCardinality)
	}

	return builder.Now()
}

func (app *adapter) cardinality(tokens instructions.Tokens) (cardinalities.Cardinality, error) {
	retMinMax, err := app.cardinalityMinMax(tokens)
	if err == nil {
		return retMinMax, nil
	}

	retMinNoMax, err := app.cardinalityMinNoMax(tokens)
	if err == nil {
		return retMinNoMax, nil
	}

	retSpecific, err := app.cardinalitySpecific(tokens)
	if err == nil {
		return retSpecific, nil
	}

	retSymbol, err := app.cardinalitySymbol(tokens)
	if err == nil {
		return retSymbol, nil
	}

	return app.cardinalityBuilder.Create().
		WithMin(uint(1)).
		WithMax(uint(1)).
		Now()
}

func (app *adapter) cardinalitySymbol(tokens instructions.Tokens) (cardinalities.Cardinality, error) {
	_, _, err := app.query(tokens, []byte(`
	v1;
	name: mySelector;
	cardinality[0][0]->INTERROGATION_POINT[0][0];
`))

	builder := app.cardinalityBuilder.Create()
	if err == nil {
		return builder.WithMin(0).WithMax(uint(1)).Now()
	}

	_, _, err = app.query(tokens, []byte(`
	v1;
	name: mySelector;
	cardinality[0][0]->STAR[0][0];
`))

	if err == nil {
		return builder.WithMin(0).Now()
	}

	_, _, err = app.query(tokens, []byte(`
	v1;
	name: mySelector;
	cardinality[0][0]->PLUS[0][0];
`))

	if err != nil {
		return nil, err
	}

	return builder.WithMin(1).Now()
}

func (app *adapter) cardinalitySpecific(tokens instructions.Tokens) (cardinalities.Cardinality, error) {
	_, retMinElement, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		cardinality[0][0]->cardinalityBracketOption[0][0]->cardinalityNumberOptions[0][0]->numbers[0][0];
	`))

	if err != nil {
		return nil, err
	}

	specific, err := strconv.Atoi(string(retMinElement.Value()))
	if err != nil {
		return nil, err
	}

	return app.cardinalityBuilder.Create().
		WithMin(uint(specific)).
		WithMax(uint(specific)).
		Now()
}

func (app *adapter) cardinalityMinNoMax(tokens instructions.Tokens) (cardinalities.Cardinality, error) {
	_, retMinElement, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		cardinality[0][0]->cardinalityBracketOption[0][0]->cardinalityNumberOptions[0][0]->minComma[0][0]->numbers[0][0];
	`))

	if err != nil {
		return nil, err
	}

	min, err := strconv.Atoi(string(retMinElement.Value()))
	if err != nil {
		return nil, err
	}

	return app.cardinalityBuilder.Create().WithMin(uint(min)).Now()
}

func (app *adapter) cardinalityMinMax(tokens instructions.Tokens) (cardinalities.Cardinality, error) {
	_, retMinElement, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		cardinality[0][0]->cardinalityBracketOption[0][0]->cardinalityNumberOptions[0][0]->minMax[0][0]->minComma[0][0]->numbers[0][0];
	`))

	if err != nil {
		return nil, err
	}

	min, err := strconv.Atoi(string(retMinElement.Value()))
	if err != nil {
		return nil, err
	}

	_, retMaxElement, err := app.query(tokens, []byte(`
		v1;
		name: mySelector;
		cardinality[0][0]->cardinalityBracketOption[0][0]->cardinalityNumberOptions[0][0]->minMax[0][0]->numbers[0][0];
	`))

	if err != nil {
		return nil, err
	}

	max, err := strconv.Atoi(string(retMaxElement.Value()))
	if err != nil {
		return nil, err
	}

	return app.cardinalityBuilder.Create().WithMin(uint(min)).WithMax(uint(max)).Now()
}

func (app *adapter) query(tokens instructions.Tokens, script []byte) (instructions.Elements, instructions.Element, error) {
	query, remaining, err := app.queryAdapter.ToQuery(script)
	if err != nil {
		return nil, nil, err
	}

	if len(remaining) > 0 {
		str := fmt.Sprintf("the script (%s) contains a remaining (%s)", script, remaining)
		return nil, nil, errors.New(str)
	}

	chain := query.Chain()
	return tokens.Select(chain)
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
