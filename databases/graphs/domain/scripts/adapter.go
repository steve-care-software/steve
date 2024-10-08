package scripts

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access/permissions"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/commons/heads/access/writes"
	"github.com/steve-care-software/steve/databases/graphs/domain/scripts/schemas"
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
}

func createAdapter(
	parserAppBuilder applications_parser.Builder,
	grammar grammars.Grammar,
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

) Adapter {
	out := adapter{
		parserAppBuilder:        parserAppBuilder,
		grammar:                 grammar,
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
				builder := app.builder.Create()
				if ins, ok := mp["schema"]; ok {
					builder.WithSchema(ins[0].(schemas.Schema))
				}

				return builder.Now()
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
								return input, nil
							},
							TokenList: &elements.TokenList{
								MapFn: func(elementName string, mp map[string][]any) (any, error) {
									return app.schemaBuilder.Create().
										WithConnections(mp["connectionBlocks"][0].(connections.Connections)).
										WithHead(mp["head"][0].(heads.Head)).
										WithPoints(mp["instructionPoints"][0].([]string)).
										Now()
								},
								List: map[string]elements.SelectedTokenList{
									"head": {
										SelectorScript: []byte(`
											v1;
											name: mySelector;
											head[0][0];
										`),
										Node: app.nodeHead(),
									},
									"instructionPoints": {
										SelectorScript: []byte(`
											v1;
											name: mySelector;
											instructionPoints[0][0]->instructionPoint[0];
										`),
										Node: &elements.Node{
											Token: &elements.Token{
												ListFn: func(list []any) (any, error) {
													output := []string{}
													for _, oneElement := range list {
														output = append(output, oneElement.(string))
													}

													return output, nil
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
									"connectionBlocks": {
										SelectorScript: []byte(`
											v1;
											name: mySelector;
											connectionBlocks[0][0]->connectionBlock;
										`),
										Node: &elements.Node{
											TokenList: &elements.TokenList{
												MapFn: func(elementName string, mp map[string][]any) (any, error) {
													return mp["connectionBlock"][0], nil
												},
												List: map[string]elements.SelectedTokenList{
													"connectionBlock": {
														SelectorScript: []byte(`
															v1;
															name: mySelector;
															connectionBlock[0];
														`),
														Node: &elements.Node{
															Token: &elements.Token{
																ListFn: func(list []any) (any, error) {
																	output := []connections.Connection{}
																	for _, oneElement := range list {
																		output = append(output, oneElement.(connections.Connection))
																	}

																	return app.connectionsBuilder.Create().
																		WithList(output).
																		Now()
																},
																Next: &elements.Element{
																	ElementFn: func(input any) (any, error) {
																		return input, nil
																	},
																	TokenList: &elements.TokenList{
																		MapFn: func(elementName string, mp map[string][]any) (any, error) {
																			builder := app.connectionBuilder.Create().WithHeader(mp["connectionHeader"][0].(connection_headers.Header)).WithLinks(mp["links"][0].(links.Links))
																			if ins, ok := mp["pointSuitesBlock"]; ok {
																				builder.WithSuites(ins[0].(suites.Suites))
																			}

																			return builder.Now()
																		},
																		List: map[string]elements.SelectedTokenList{
																			"connectionHeader": {
																				SelectorScript: []byte(`
																					v1;
																					name: mySelector;
																					connectionHeader[0][0];
																				`),
																				Node: &elements.Node{
																					Element: &elements.Element{
																						ElementFn: func(input any) (any, error) {
																							return input, nil
																						},
																						TokenList: &elements.TokenList{
																							MapFn: func(elementName string, mp map[string][]any) (any, error) {
																								builder := app.connectionHeaderBuilder.Create().
																									WithName(mp["nameWithCardinality"][0].(names.Name))

																								if reverse, ok := mp["variableNameInParenthesis"]; ok {
																									builder.WithReverse(reverse[0].(names.Name))
																								}

																								return builder.Now()
																							},
																							List: map[string]elements.SelectedTokenList{
																								"nameWithCardinality": {
																									SelectorScript: []byte(`
																										v1;
																										name: mySelector;
																										nameWithCardinality[0][0];
																									`),
																									Node: app.nodeNameWithCardinality(),
																								},
																								"variableNameInParenthesis": {
																									SelectorScript: []byte(`
																										v1;
																										name: mySelector;
																										variableNameInParenthesis[0][0]->nameWithCardinality[0][0];
																									`),
																									Node: app.nodeNameWithCardinality(),
																								},
																							},
																						},
																					},
																				},
																			},
																			"links": {
																				SelectorScript: []byte(`
																					v1;
																					name: mySelector;
																					links[0][0];
																				`),
																				Node: &elements.Node{
																					Element: &elements.Element{
																						ElementFn: func(input any) (any, error) {
																							return input, nil
																						},
																						TokenList: &elements.TokenList{
																							MapFn: func(elementName string, mp map[string][]any) (any, error) {
																								referencesList := mp["link"][0].(references.References).List()
																								link, err := app.linkBuilder.Create().
																									WithOrigin(referencesList[0]).
																									WithTarget(referencesList[1]).
																									Now()

																								if err != nil {
																									return nil, err
																								}

																								list := []links.Link{
																									link,
																								}

																								if ins, ok := mp["pipeJourney"]; ok {
																									list = append(list, ins[0].([]links.Link)...)
																								}

																								return app.linksBuilder.Create().
																									WithList(list).
																									Now()
																							},
																							List: map[string]elements.SelectedTokenList{
																								"link": {
																									SelectorScript: []byte(`
																										v1;
																										name: mySelector;
																										link[0][0]->pointReference[0];
																									`),
																									Node: app.nodePointReference(),
																								},
																								"pipeJourney": {
																									SelectorScript: []byte(`
																										v1;
																										name: mySelector;
																										pipeJourney[0];
																									`),
																									Node: &elements.Node{
																										Token: &elements.Token{
																											ListFn: func(list []any) (any, error) {
																												output := []links.Link{}
																												for _, oneElement := range list {
																													output = append(output, oneElement.(links.Link))
																												}

																												return output, nil
																											},
																											Next: &elements.Element{
																												ElementFn: func(input any) (any, error) {
																													return input, nil
																												},
																												TokenList: &elements.TokenList{
																													MapFn: func(elementName string, mp map[string][]any) (any, error) {
																														referencesList := mp["link"][0].(references.References).List()
																														return app.linkBuilder.Create().
																															WithOrigin(referencesList[0]).
																															WithTarget(referencesList[1]).
																															Now()
																													},
																													List: map[string]elements.SelectedTokenList{
																														"link": {
																															SelectorScript: []byte(`
																																v1;
																																name: mySelector;
																																link[0][0]->pointReference[0];
																															`),
																															Node: app.nodePointReference(),
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
																			"pointSuitesBlock": {
																				SelectorScript: []byte(`
																					v1;
																					name: mySelector;
																					pointSuitesBlock[0][0]->pointSuites[0][0]->pointSuite[0];
																				`),
																				Node: &elements.Node{
																					Token: &elements.Token{
																						ListFn: func(list []any) (any, error) {
																							output := []suites.Suite{}
																							for _, oneElement := range list {
																								output = append(output, oneElement.(suites.Suite))
																							}

																							return app.suitesBuilder.Create().
																								WithList(output).
																								Now()
																						},
																						Next: &elements.Element{
																							ElementFn: func(input any) (any, error) {
																								return input, nil
																							},
																							TokenList: &elements.TokenList{
																								MapFn: func(elementName string, mp map[string][]any) (any, error) {
																									referencesList := mp["pointReference"][0].(references.References).List()
																									link, err := app.linkBuilder.Create().
																										WithOrigin(referencesList[0]).
																										WithTarget(referencesList[1]).
																										Now()

																									if err != nil {
																										return nil, err
																									}

																									return app.suiteBuilder.Create().
																										WithName(mp["variableName"][0].(string)).
																										WithLink(link).
																										WithExpectations(mp["suiteInstructions"][0].(expectations.Expectations)).
																										Now()
																								},
																								List: map[string]elements.SelectedTokenList{
																									"variableName": {
																										SelectorScript: []byte(`
																											v1;
																											name: mySelector;
																											variableName[0][0];
																										`),
																										Node: app.nodeByteToString(),
																									},
																									"pointReference": {
																										SelectorScript: []byte(`
																											v1;
																											name: mySelector;
																											pointReference[0];
																										`),
																										Node: app.nodePointReference(),
																									},
																									"suiteInstructions": {
																										SelectorScript: []byte(`
																											v1;
																											name: mySelector;
																											suiteInstructions[0][0]->suiteOptionInstruction[0];
																										`),
																										Node: &elements.Node{
																											Token: &elements.Token{
																												ListFn: func(list []any) (any, error) {
																													output := []expectations.Expectation{}
																													for _, oneElement := range list {
																														output = append(output, oneElement.(expectations.Expectation))
																													}

																													return app.expectationsBuilder.Create().
																														WithList(output).
																														Now()
																												},
																												Next: &elements.Element{
																													ElementFn: func(input any) (any, error) {
																														return input, nil
																													},
																													TokenList: &elements.TokenList{
																														MapFn: func(elementName string, mp map[string][]any) (any, error) {
																															return mp["suiteOption"][0], nil
																														},
																														List: map[string]elements.SelectedTokenList{
																															"suiteOption": {
																																SelectorScript: []byte(`
																																	v1;
																																	name: mySelector;
																																	suiteOption[0][0];
																																`),
																																Node: &elements.Node{
																																	Element: &elements.Element{
																																		ElementFn: func(input any) (any, error) {
																																			return input, nil
																																		},
																																		TokenList: &elements.TokenList{
																																			MapFn: func(elementName string, mp map[string][]any) (any, error) {
																																				builder := app.expectationBuilder.Create()
																																				if ins, ok := mp["suiteInvalidLink"]; ok {
																																					builder.IsFail().
																																						WithReferences(ins[0].(references.References))
																																				}

																																				if ins, ok := mp["suiteReferencesInParenthesis"]; ok {
																																					builder.WithReferences(ins[0].(references.References))
																																				}

																																				return builder.Now()
																																			},
																																			List: map[string]elements.SelectedTokenList{
																																				"suiteInvalidLink": {
																																					SelectorScript: []byte(`
																																						v1;
																																						name: mySelector;
																																						suiteInvalidLink[0][0]->suiteReferencesInParenthesis[0][0]->pointReferences[0][0]->pointReference[0];
																																					`),
																																					Node: app.nodePointReference(),
																																				},
																																				"suiteReferencesInParenthesis": {
																																					SelectorScript: []byte(`
																																						v1;
																																						name: mySelector;
																																						suiteReferencesInParenthesis[0][0]->pointReferences[0][0]->pointReference[0];
																																					`),
																																					Node: app.nodePointReference(),
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

func (app *adapter) nodeByteToString() *elements.Node {
	return &elements.Node{
		Element: &elements.Element{
			ElementFn: func(input any) (any, error) {
				return string(input.([]byte)), nil
			},
		},
	}
}

func (app *adapter) nodeByteToUint() *elements.Node {
	return &elements.Node{
		Element: &elements.Element{
			ElementFn: func(input any) (any, error) {
				value, err := strconv.Atoi(string(input.([]byte)))
				if err != nil {
					return nil, err
				}
				return uint(value), nil
			},
		},
	}
}

func (app *adapter) nodeNameWithCardinality() *elements.Node {
	return &elements.Node{
		Element: &elements.Element{
			ElementFn: func(input any) (any, error) {
				return input, nil
			},
			TokenList: &elements.TokenList{
				MapFn: func(elementName string, mp map[string][]any) (any, error) {

					builder := app.nameBuilder.Create().
						WithName(mp["variableName"][0].(string))

					if ins, ok := mp["cardinality"]; ok {
						builder.WithCardinality(ins[0].(cardinalities.Cardinality))
					} else {
						ins, err := app.cardinalityBuilder.Create().
							WithMin(1).
							WithMax(1).
							Now()

						if err != nil {
							return nil, err
						}

						builder.WithCardinality(ins)
					}

					return builder.Now()
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
					"cardinality": {
						SelectorScript: []byte(`
						v1;
						name: mySelector;
						cardinality[0][0]->cardinalityBracketOption[0][0]->cardinalityNumberOptions[0][0];
					`),
						Node: &elements.Node{
							Element: &elements.Element{
								ElementFn: func(input any) (any, error) {
									return input, nil
								},
								TokenList: &elements.TokenList{
									MapFn: func(elementName string, mp map[string][]any) (any, error) {
										if value, ok := mp["minMax"]; ok {
											return value[0], nil
										}

										builder := app.cardinalityBuilder.Create()
										if value, ok := mp["minComma"]; ok {
											builder.WithMin(value[0].(uint))
										}

										if value, ok := mp["numbers"]; ok {
											builder.WithMin(value[0].(uint)).
												WithMax(value[0].(uint))
										}

										return builder.Now()
									},
									List: map[string]elements.SelectedTokenList{
										"minMax": {
											SelectorScript: []byte(`
												v1;
												name: mySelector;
												minMax[0][0];
											`),
											Node: &elements.Node{
												Element: &elements.Element{
													ElementFn: func(input any) (any, error) {
														return input, nil
													},
													TokenList: &elements.TokenList{
														MapFn: func(elementName string, mp map[string][]any) (any, error) {
															return app.cardinalityBuilder.Create().
																WithMin(mp["minComma"][0].(uint)).
																WithMax(mp["numbers"][0].(uint)).
																Now()
														},
														List: map[string]elements.SelectedTokenList{
															"minComma": {
																SelectorScript: []byte(`
																	v1;
																	name: mySelector;
																	minComma[0][0]->numbers[0][0];
																`),
																Node: app.nodeByteToUint(),
															},
															"numbers": {
																SelectorScript: []byte(`
																	v1;
																	name: mySelector;
																	numbers[0][0];
																`),
																Node: app.nodeByteToUint(),
															},
														},
													},
												},
											},
										},
										"minComma": {
											SelectorScript: []byte(`
												v1;
												name: mySelector;
												minComma[0][0]->numbers[0][0];
											`),
											Node: app.nodeByteToUint(),
										},
										"numbers": {
											SelectorScript: []byte(`
												v1;
												name: mySelector;
												numbers[0][0];
											`),
											Node: app.nodeByteToUint(),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (app *adapter) nodeHead() *elements.Node {
	return &elements.Node{
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
	}
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
						Node: app.nodeReferences(),
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

func (app *adapter) nodeReferences() *elements.Node {
	return &elements.Node{
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
										Node: app.nodeReference(),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (app *adapter) nodeReference() *elements.Node {
	return &elements.Node{
		Token: &elements.Token{
			ListFn: func(list []any) (any, error) {
				return list, nil
			},
			Next: app.elementReference(),
		},
	}
}

func (app *adapter) elementReference() *elements.Element {
	return &elements.Element{
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
	}
}

func (app *adapter) nodePointReference() *elements.Node {
	return &elements.Node{
		Token: &elements.Token{
			ListFn: func(list []any) (any, error) {
				output := []references.Reference{}
				for _, oneElement := range list {
					output = append(output, oneElement.(references.Reference))
				}

				return app.referencesBuilder.Create().
					WithList(output).
					Now()
			},
			Next: &elements.Element{
				ElementFn: func(input any) (any, error) {
					return input, nil
				},
				TokenList: &elements.TokenList{
					MapFn: func(elementName string, mp map[string][]any) (any, error) {
						builder := app.referenceBuilder.Create()
						if ins, ok := mp["reference"]; ok {
							builder.WithInternal(ins[0].(string))
						}

						if ins, ok := mp["externalPointReference"]; ok {
							builder.WithExternal(ins[0].(externals.External))
						}

						return builder.Now()
					},
					List: map[string]elements.SelectedTokenList{
						"reference": {
							SelectorScript: []byte(`
								v1;
								name: mySelector;
								reference[0][0];
							`),
							Node: &elements.Node{
								Element: app.elementReference(),
							},
						},
						"externalPointReference": {
							SelectorScript: []byte(`
								v1;
								name: mySelector;
								externalPointReference[0][0];
							`),
							Node: &elements.Node{
								Element: &elements.Element{
									ElementFn: func(input any) (any, error) {
										return input, nil
									},
									TokenList: &elements.TokenList{
										MapFn: func(elementName string, mp map[string][]any) (any, error) {
											return app.externalBuilder.Create().
												WithSchema(mp["reference"][0].(string)).
												WithPoint(mp["variableName"][0].(string)).
												Now()
										},
										List: map[string]elements.SelectedTokenList{
											"reference": {
												SelectorScript: []byte(`
													v1;
													name: mySelector;
													reference[0][0];
												`),
												Node: &elements.Node{
													Element: app.elementReference(),
												},
											},
											"variableName": {
												SelectorScript: []byte(`
													v1;
													name: mySelector;
													variableName[0][0];
												`),
												Node: app.nodeByteToString(),
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
	}
}
