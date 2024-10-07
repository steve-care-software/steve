package scripts

import (
	"errors"
	"fmt"
	"strconv"

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
	"github.com/steve-care-software/steve/parsers/domain/walkers/elements"
)

type adapter struct {
	parserAppBuilder        applications_parser.Builder
	grammar                 grammars.Grammar
	builder                 Builder
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
}

func createAdapter(
	parserAppBuilder applications_parser.Builder,
	grammar grammars.Grammar,
	builder Builder,
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

) Adapter {
	out := adapter{
		parserAppBuilder:        parserAppBuilder,
		grammar:                 grammar,
		builder:                 builder,
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
	}

	return &out
}

// ToScript converts an input to script
func (app *adapter) ToScript(input []byte) (Script, []byte, error) {
	sequence := elements.Element{
		ElementFn: func(input any) (any, error) {
			return input, nil
		},
		TokenList: &elements.TokenList{
			MapFn: func(elementName string, mp map[string][]any) (any, error) {
				return nil, nil
			},
			List: map[string]elements.SelectedTokenList{
				"schema": {
					SelectorScript: []byte(`
						v1;
						name: mySelector;
						schema[0][0];
					`),
					Node: &elements.Node{
						Element: &elements.Element{
							ElementFn: func(input any) (any, error) {
								fmt.Printf("\n head: %v\n", input)
								return nil, nil
							},
							TokenList: &elements.TokenList{
								MapFn: func(elementName string, mp map[string][]any) (any, error) {
									fmt.Printf("\n schema: %s, %v\n", elementName, mp)
									return nil, nil
								},
								List: map[string]elements.SelectedTokenList{
									"head": {
										SelectorScript: []byte(`
											v1;
											name: mySelector;
											head[0][0];
										`),
										Node: &elements.Node{
											Element: &elements.Element{
												ElementFn: func(input any) (any, error) {
													return input, nil
												},
												TokenList: &elements.TokenList{
													MapFn: func(elementName string, mp map[string][]any) (any, error) {
														return app.headBuilder.Create().
															WithName(mp["propertyName"][0].(string)).
															WithVersion(mp["engine"][0].(uint)).
															WithAccess(mp["access"][0].(access.Access)).
															Now()
													},
													List: map[string]elements.SelectedTokenList{
														"engine": {
															SelectorScript: []byte(`
																v1;
																name: mySelector;
																engine[0][0]->version[0][0]->numbersExceptZero[0][0];
															`),
															Node: &elements.Node{
																Element: &elements.Element{
																	ElementFn: func(input any) (any, error) {
																		value, err := strconv.Atoi(string(input.([]byte)))
																		if err != nil {
																			return nil, err
																		}

																		return uint(value), nil
																	},
																},
															},
														},
														"propertyName": {
															SelectorScript: []byte(`
																v1;
																name: mySelector;
																propertyName[0][0]->variableName[0][0];
															`),
															Node: &elements.Node{
																Element: &elements.Element{
																	ElementFn: func(input any) (any, error) {
																		return string(input.([]byte)), nil
																	},
																},
															},
														},
														"access": {
															SelectorScript: []byte(`
																v1;
																name: mySelector;
																access[0][0]->roleOptions[0][0];
															`),
															Node: &elements.Node{
																Element: &elements.Element{
																	ElementFn: func(input any) (any, error) {
																		return input, nil
																	},
																	TokenList: &elements.TokenList{
																		MapFn: func(elementName string, mp map[string][]any) (any, error) {
																			builder := app.accessBuilder.Create().WithWrite(mp["roleOptionWrite"][0].(writes.Write))
																			if list, ok := mp["roleOptionRead"]; ok {
																				builder.WithRead(list[0].(permissions.Permission))
																			}

																			return builder.Now()
																		},
																		List: map[string]elements.SelectedTokenList{
																			"roleOptionRead": {
																				SelectorScript: []byte(`
																					v1;
																					name: mySelector;
																					roleOptionRead[0][0]->roleOptionSuffix[0][0]->referencesCompensation[0][0];
																				`),
																				Node: app.nodeReferencesCompensation(),
																			},
																			"roleOptionWrite": {
																				SelectorScript: []byte(`
																				v1;
																					name: mySelector;
																					roleOptionWrite[0][0];
																				`),
																				Node: &elements.Node{
																					Element: &elements.Element{
																						ElementFn: func(input any) (any, error) {
																							return input, nil
																						},
																						TokenList: &elements.TokenList{
																							MapFn: func(elementName string, mp map[string][]any) (any, error) {
																								builder := app.writeBuilder.Create().WithModify(mp["referencesCompensation"][0].(permissions.Permission))
																								if list, ok := mp["roleOptionReview"]; ok {
																									builder.WithReview(list[0].(permissions.Permission))
																								}

																								return builder.Now()
																							},
																							List: map[string]elements.SelectedTokenList{
																								"referencesCompensation": {
																									SelectorScript: []byte(`
																										v1;
																										name: mySelector;
																										referencesCompensation[0][0];
																									`),
																									Node: app.nodeReferencesCompensation(),
																								},

																								"roleOptionReview": {
																									SelectorScript: []byte(`
																										v1;
																										name: mySelector;
																										roleOptionReview[0][0]->roleOptionSuffix[0][0]->referencesCompensation[0][0];
																									`),
																									Node: app.nodeReferencesCompensation(),
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									//"instructionPoints": {},
									//"connectionBlocks":  {},
								},
							},
						},
					},
				},
			},
		},
	}

	parserApp, err := app.parserAppBuilder.Create().
		WithElement(sequence).
		Now()

	if err != nil {
		return nil, nil, err
	}

	retIns, retRemaining, err := parserApp.Execute(input, app.grammar)
	if err != nil {
		return nil, nil, err
	}

	if casted, ok := retIns.(Script); ok {
		return casted, retRemaining, nil
	}

	return nil, nil, errors.New("the returned parser instance was was expected to contain a Script instance")
}

func (app *adapter) nodeReferencesCompensation() *elements.Node {
	return &elements.Node{
		Element: &elements.Element{
			ElementFn: func(input any) (any, error) {
				return input, nil
			},
			TokenList: &elements.TokenList{
				MapFn: func(elementName string, mp map[string][]any) (any, error) {
					names := []string{}
					for _, oneName := range mp["references"][0].([]any) {
						names = append(names, oneName.(string))
					}
					builder := app.permissionBuilder.Create().
						WithNames(names)

					if list, ok := mp["floatNumberBetweenZeroAndOneInParenthesis"]; ok {
						builder.WithCompensation(list[0].(float64))
					}

					return builder.Now()
				},
				List: map[string]elements.SelectedTokenList{
					"references": {
						SelectorScript: []byte(`
						v1;
						name: mySelector;
						references[0][0];
					`),
						Node: &elements.Node{
							Element: &elements.Element{
								ElementFn: func(input any) (any, error) {
									return input, nil
								},
								TokenList: &elements.TokenList{
									MapFn: func(elementName string, mp map[string][]any) (any, error) {
										return mp["reference"][0], nil
									},
									List: map[string]elements.SelectedTokenList{
										"reference": {
											SelectorScript: []byte(`
											v1;
											name: mySelector;
											reference;
										`),
											Node: &elements.Node{
												TokenList: &elements.TokenList{
													MapFn: func(elementName string, mp map[string][]any) (any, error) {
														return mp["reference"][0], nil
													},
													List: map[string]elements.SelectedTokenList{
														"reference": {
															SelectorScript: []byte(`
															v1;
															name: mySelector;
															reference[0];
														`),
															Node: &elements.Node{
																Token: &elements.Token{
																	ListFn: func(list []any) (any, error) {
																		return list, nil
																	},
																	Next: &elements.Element{
																		ElementFn: func(input any) (any, error) {
																			return input, nil
																		},
																		TokenList: &elements.TokenList{
																			MapFn: func(elementName string, mp map[string][]any) (any, error) {
																				return mp["variableName"][0], nil
																			},
																			List: map[string]elements.SelectedTokenList{
																				"variableName": {
																					SelectorScript: []byte(`
																					v1;
																					name: mySelector;
																					variableName[0][0];
																				`),
																					Node: &elements.Node{
																						Element: &elements.Element{
																							ElementFn: func(input any) (any, error) {
																								return string(input.([]byte)), nil
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					"floatNumberBetweenZeroAndOneInParenthesis": {
						SelectorScript: []byte(`
						v1;
						name: mySelector;
						floatNumberBetweenZeroAndOneInParenthesis[0][0]->floatNumberBetweenZeroAndOne[0][0];
					`),
						Node: &elements.Node{
							Element: &elements.Element{
								ElementFn: func(input any) (any, error) {
									value, err := strconv.ParseFloat(string(input.([]byte)), 64)
									if err != nil {
										return nil, err
									}

									return value, nil
								},
							},
						},
					},
				},
			},
		},
	}
}
