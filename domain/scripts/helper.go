package scripts

func grammarInput() []byte {
	return []byte(`
		v1;
		> .script;
		# .SPACE .TAB .EOL;

		script: .grammarDefinition
			  | .transpileDefinition
			  | .pipelineDefinition
			  | .rootDefinition
			  | .schema
			---
				grammar: "
								head:
									engine: v1;
									name: myGrammar;
									access: 
										read: .first;
										write: .first .second;
										review: .first;
									;
									compensation: 0.1, 0.23, 0.45;
								;

								entry: .myBlock;

								myBlock: .first ._second*
									---
										validString: \"this is some input\";
									;

								OTHER_RULE: [0, 1, 2, 3];
								MY_RULE: \"A\";
					";

				transpile: "
								head:
									engine: v1;
									name: myTranspile;
									access: 
										read: .first;
										write: .first .second;
										review: .first;
									;
									compensation: 0.1, 0.23, 0.45;
								;

								origin: .myOriginGrammar;

								myBlock: .myElement .myElement[0][1]->_myConstant->mySubElement[0]->MY_RULE[0] ._myConstant[3] .myTargetGrammar[_targetConst] .myTargetGrammar[TARGET_RULE[3]];
								mySecond: .myTargetGrammar[_targetConst] .myTargetGrammar[TARGET_RULE[3]]
									---
										myTest:
											input: \"this is some data\";
											expected: [2, 3, 4, 5];
										;

										myOtherTest:
											input: [2, 3, 4, 5];
											expected: \"this is the expected output\";
										;

										expectedImvalid:
											input: !\"this is some data\";
										;
									;
						";

				pipeline: "
								head:
									engine: v1;
									name: myPipeline;
									access: 
										read: .first;
										write: .first .second;
										review: .first;
									;
									compensation: 0.1, 0.23, 0.45;
								;
								recipe: .myTranspile;
								program: myProgramOutputVar [
									recipe: .myTranspile;
									program: myProgramOutputVar [
										recipe: .myTranspile;
									];
								]
								---
									myTest:
										input: \"this is some data\";
										expected: [2, 3, 4, 5];
									;

									myOtherTest:
										input: [2, 3, 4, 5];
										expected: \"this is the expected output\";
									;
								;
						";

				root: "
						head:
							engine: v1;
							name: myRoot;
							access: 
								read: .first;
								write: .first .second;
								review: .first;
							;
							compensation: 0.1, 0.23, 0.45;
						;

						.first;
						.second:
							.sub:
								.subSub
								---
									myTest:
										input: \"this is some data\";
										expected: [2, 3, 4, 5];
									;

									myOtherTest:
										input: ![2, 3, 4, 5];
									;
								;
							;
						;
						.third
						---
							myOtherTest:
								input: ![2, 3, 4, 5];
							;
						;
				";

				schema: "
							head:
									engine: v1;
									name: myName;
									access: 
										read: .first;
										write: .first .second;
										review: .first;
									;
									compensation: 0.1, 0.23, 0.45;
							;

							son;
							father;
							grandFather;
							grandGrandFather;
							myList: list[uint8];
							mySet: set[uint8];
							myMap: map[uint8];
							mySortedSet: sortedSet[uint8];

							father[son]: .son .father
										| .father .grandFather
										| .grandFather .grandGrandFather
										---
											mySuite[.son .grandGrandFather]:
												(.son .father .grandFather .grandGrandFather);
												!(.son .father .grandFather .grandGrandFather);
											;
										;

							grandFather[grandSon]: .son .grandFather
												| .father .grandGrandFather
												;
				";
			;

		rootDefinition: .head .rootElements
					---
						valid: "
								head:
									engine: v1;
									name: myRoot;
									access: 
										read: .first;
										write: .first .second;
										review: .first;
									;
									compensation: 0.1, 0.23, 0.45;
								;

								.first;
								.second:
									.sub:
										.subSub
										---
											myTest:
												input: \"this is some data\";
												expected: [2, 3, 4, 5];
											;

											myOtherTest:
												input: ![2, 3, 4, 5];
											;
										;
									;
								;
								.third
								---
									myOtherTest:
										input: ![2, 3, 4, 5];
									;
								;
						";
					;

		rootSubElement: .COLON .rootElements;

		rootElements: .rootElement+
				---
					valid: "
							.first;
							.second:
								.sub:
									.subSub
									---
										myTest:
											input: \"this is some data\";
											expected: [2, 3, 4, 5];
										;

										myOtherTest:
											input: ![2, 3, 4, 5];
										;
									;
								;
							;
							.third
							---
								myOtherTest:
									input: ![2, 3, 4, 5];
								;
							;
					";
				;

		rootElement: .reference .rootSubElement? .transpileSuites? .SEMI_COLON
					---
						simple: "
							.first;
						";

						withSubElement: "
							.first:
								.sub:
									.subSub
									---
										myTest:
											input: \"this is some data\";
											expected: [2, 3, 4, 5];
										;

										myOtherTest:
											input: ![2, 3, 4, 5];
										;
									;
								;
							;
						";

						withSuite: "
							.first
							---
								myTest:
									input: \"this is some data\";
									expected: [2, 3, 4, 5];
								;

								myOtherTest:
									input: ![2, 3, 4, 5];
								;
							;
						";
					;

		pipelineDefinition: .head .pipelineInstruction
							---
								valid: "
										head:
											engine: v1;
											name: myPipeline;
											access: 
												read: .first;
												write: .first .second;
												review: .first;
											;
											compensation: 0.1, 0.23, 0.45;
										;
										recipe: .myTranspile;
										program: myProgramOutputVar [
											recipe: .myTranspile;
											program: myProgramOutputVar [
												recipe: .myTranspile;
											];
										]
										---
											myTest:
												input: \"this is some data\";
												expected: [2, 3, 4, 5];
											;

											myOtherTest:
												input: [2, 3, 4, 5];
												expected: \"this is the expected output\";
											;
										;
								";
							;

		pipelineInstruction: .pipeLineRecipeDefinition .pipelineProgramDefinition?
							---
								recipe: "
									recipe: .myTranspile;
								";

								recipeWithProgram: "
										recipe: .myTranspile;
										program: myProgramOutputVar [
											recipe: .myTranspile;
											program: myProgramOutputVar [
												recipe: .myTranspile;
											];
										]
										---
											myTest:
												input: \"this is some data\";
												expected: [2, 3, 4, 5];
											;

											myOtherTest:
												input: [2, 3, 4, 5];
												expected: \"this is the expected output\";
											;
										;
								";

								
							;

		pipelineProgramDefinition: .PROGRAM .COLON .variableName .BRACKET_OPEN .pipelineInstruction .BRACKET_CLOSE .transpileSuites? .SEMI_COLON
								---
									recipe: "
										program: myProgramOutputVar [
											recipe: .myTranspile;
										];
									";

									recipeWithSuites: "
										program: myProgramOutputVar [
											recipe: .myTranspile;
										]
										---
											myTest:
												input: \"this is some data\";
												expected: [2, 3, 4, 5];
											;

											myOtherTest:
												input: [2, 3, 4, 5];
												expected: \"this is the expected output\";
											;
										;
									";

									recipeWithSubProgramWithSuites: "
										program: myProgramOutputVar [
											recipe: .myTranspile;
											program: myProgramOutputVar [
												recipe: .myTranspile;
											];
										]
										---
											myTest:
												input: \"this is some data\";
												expected: [2, 3, 4, 5];
											;

											myOtherTest:
												input: [2, 3, 4, 5];
												expected: \"this is the expected output\";
											;
										;
									";
									
								;

		pipeLineRecipeDefinition: .RECIPE .COLON .reference .SEMI_COLON
								---
									valid: "
										recipe: .myTranspile;
									";
								;

		originDefinition: .ORIGIN .COLON .reference .SEMI_COLON
						---
							valid: "origin: .myOriginGrammar;";
						;

		transpileDefinition: .head .originDefinition .transpileBlocks
							---
								valid: "
										head:
											engine: v1;
											name: myTranspile;
											access: 
												read: .first;
												write: .first .second;
												review: .first;
											;
											compensation: 0.1, 0.23, 0.45;
										;

										origin: .myOriginGrammar;

										myBlock: .myElement .MY_RULE[1] ._myConstant[3] .myTargetGrammar[_targetConst] .myTargetGrammar[TARGET_RULE[3]];
										mySecond: .myTargetGrammar[_targetConst] .myTargetGrammar[TARGET_RULE[3]]
											---
												myTest:
													input: \"this is some data\";
													expected: [2, 3, 4, 5];
												;

												myOtherTest:
													input: [2, 3, 4, 5];
													expected: \"this is the expected output\";
												;
											;
								";
							;

		transpileBlocks: .transpileBlock+
						---
							single: "
								myBlock: .myElement .MY_RULE[1] ._myConstant[3] .myTargetGrammar[_targetConst] .myTargetGrammar[TARGET_RULE[3]];
							";

							multiple: "
								myBlock: .myElement .MY_RULE[1] ._myConstant[3] .myTargetGrammar[_targetConst] .myTargetGrammar[TARGET_RULE[3]];
								mySecond: .myTargetGrammar[_targetConst] .myTargetGrammar[TARGET_RULE[3]]
									---
										myTest:
											input: \"this is some data\";
											expected: [2, 3, 4, 5];
										;

										myOtherTest:
											input: [2, 3, 4, 5];
											expected: \"this is the expected output\";
										;
									;
							";
						;

		transpileBlock: .variableName .COLON .astElements .pipeASTElements? .transpileSuites? .SEMI_COLON
						---
							oneInstruction: "
									myBlock: .myElement .MY_RULE[1] ._myConstant[3] .myTargetGrammar[_targetConst] .myTargetGrammar[TARGET_RULE[3]];
							";

							oneInstructionWithTests: "
									myBlock: .myElement .MY_RULE[1] ._myConstant[3] .myTargetGrammar[_targetConst] .myTargetGrammar[TARGET_RULE[3]]
									---
										myTest:
											input: \"this is some data\";
											expected: [2, 3, 4, 5];
										;

										myOtherTest:
											input: [2, 3, 4, 5];
											expected: \"this is the expected output\";
										;
									;
							";

							multipleInstruction: "
									myBlock: .myElement .MY_RULE[1] ._myConstant[3] .myTargetGrammar[_targetConst] .myTargetGrammar[TARGET_RULE[3]]
											| .myTargetGrammar[TARGET_RULE[3]]
											;
							";

							multipleInstructionWithTests: "
									myBlock: .myElement .MY_RULE[1] ._myConstant[3] .myTargetGrammar[_targetConst] .myTargetGrammar[TARGET_RULE[3]]
											| .myTargetGrammar[TARGET_RULE[3]]
											---
												myTest:
													input: \"this is some data\";
													expected: [2, 3, 4, 5];
												;

												myOtherTest:
													input: [2, 3, 4, 5];
													expected: \"this is the expected output\";
												;
											;
							";
						;

		transpileSuites: .suitePrefix .transpileSuite+
						---
							valid: "
									---
									expectedImvalid:
										input: !\"this is some data\";
									;

									myTest:
										input: \"this is some data\";
										expected: [2, 3, 4, 5];
									;

									myOtherTest:
										input: [2, 3, 4, 5];
										expected: \"this is the expected output\";
									;
							";
						;

		transpileSuite: .transpileSuiteInvalid
					  | .transpileSuiteValid
					---
						expectedInvalid: "
								myTest:
									input: !\"this is some data\";
								;
						";

						expectedValid: "
								myTest:
									input: \"this is some data\";
									expected: [2, 3, 4, 5];
								;
						";
					;

		transpileSuiteInvalid: .variableName .COLON .suiteInputInvalid .SEMI_COLON
							---
								string: "
										myTest:
											input: !\"this is some data\";
										;
								";

								bytes: "
										myTest:
											input: ![0, 1, 2];
										;
								";
							;

		transpileSuiteValid: .variableName .COLON .suiteInput .suiteExpected .SEMI_COLON
							---
								valid: "
										myTest:
											input: \"this is some data\";
											expected: [2, 3, 4, 5];
										;
								";
							;

		suiteExpected: .EXPECTED .COLON .blockSuiteValue .SEMI_COLON
					---
						string: "
							expected: \"this is some data\";
						";

						bytes: "
							expected: [2, 3, 4, 5];
						";
					;

		suiteInputInvalid: .INPUT .COLON .EXCLAMATION_POINT .blockSuiteValue .SEMI_COLON
				---
					string: "
						input: !\"this is some data\";
					";

					bytes: "
						input: ![2, 3, 4, 5];
					";
				;

		suiteInput: .INPUT .COLON .blockSuiteValue .SEMI_COLON
				---
					string: "
						input: \"this is some data\";
					";

					bytes: "
						input: [2, 3, 4, 5];
					";
				;

		pipeASTElements: .PIPE .astElements
					---
						origin: "| .MY_RULE ._myConstant[5] .myToken";
						target: "| .myTargetGrammar[MY_RULE[3]]";
					;

		astElements: .astElement+
				---
					valid: ".MY_RULE ._myConstant[5] .myToken";
				;

		astElement: .astTargetElement
				  | .selectorValue
				  ---
				  		origin: ".MY_RULE";
						target: ".myElement[0][1]->_myConstant->mySubElement[0]->MY_RULE[0]";
				  ;

		astTargetElement: .reference .BRACKET_OPEN .astTargetElementKeyname .BRACKET_CLOSE
						---
							ruleWithAmount: ".myTargetGrammar[MY_RULE[3]]";
							rule: ".myTargetGrammar[MY_RULE]";
							constantWithAmount: ".myTargetGrammar[_myConstant[3]]";
							constant: ".myTargetGrammar[_myConstant]";
						;

		astTargetElementKeyname: .constantElementName .specificAmount?
								---
									constant: "_myConstant";
									rule: "MY_RULE";
									constantWithAmount: "_myConstant[3]";
									ruleWithAmount: "MY_RULE[2]";
									block: !"myBlock";
								;

		selectorValue: .DOT .selectorElement .selectorSubElement*
					---
						one: ".myElement";
						oneWithIndexes: ".myElement[0][1]";
						complex: ".myElement[0][1]->mySubElement[0]->MY_RULE[0]->_myConstant";
					;

		selectorSubElement: .HYPHEN .GREATHER_THAN .selectorElement
						---
							simple: "->MY_RULE";
							simpleWithOneIndex: "->myToken[0]";
							simpleWithTwoIndex: "->myToken[1][2]";
						;

		selectorElement: .blockElementName .doubleSpecificValue?
						---
							simple: "MY_RULE";
							simpleWithOneIndex: "myToken[0]";
							simpleWithTwoIndex: "myToken[1][2]";
						;

		doubleSpecificValue: .specificAmount[1,2]
							---
								one: "[0]";
								two: "[0][1]";
							;

		grammarDefinition: .head .entryDefinition .omitDefinition? .blockDefinition+ .constantDefinition* .ruleDefinition+
						---
							simple: "
								head:
										engine: v1;
										name: myGrammar;
										access: 
											read: .first;
											write: .first .second;
											review: .first;
										;
										compensation: 0.1, 0.23, 0.45;
								;

								entry: .myBlock;

								myBlock: .first ._second*
									---
										validString: \"this is some input\";
									;

								OTHER_RULE: [0, 1, 2, 3];
								MY_RULE: \"A\";
							";

							withOmit: "
								head:
										engine: v1;
										name: myGrammar;
										access: 
											read: .first;
											write: .first .second;
											review: .first;
										;
										compensation: 0.1, 0.23, 0.45;
								;

								entry: .myBlock;
								omit: .SPACE .TAB. EOL;

								myBlock: .first ._second*
									---
										validString: \"this is some input\";
									;

								OTHER_RULE: [0, 1, 2, 3];
								MY_RULE: \"A\";
							";

							withConstant: "
								head:
										engine: v1;
										name: myGrammar;
										access: 
											read: .first;
											write: .first .second;
											review: .first;
										;
										compensation: 0.1, 0.23, 0.45;
								;

								entry: .myBlock;

								myBlock: .first ._second*
									---
										validString: \"this is some input\";
									;

								_myConstant: ._myConstant .MY_RULE[2];

								OTHER_RULE: [0, 1, 2, 3];
								MY_RULE: \"A\";
							";

							
						;

		omitDefinition: .OMIT .COLON .blockElementReference+ .SEMI_COLON
					---
						multiple: "
							omit: .SPACE .TAB. EOL;
						";

						single: "
							omit: ._myConstant;
						";
					;

		entryDefinition: .ENTRY .COLON .blockElementReference .SEMI_COLON
						---
							valid: "
								entry: .myElement;
							";
						;

		blockDefinition: .variableName .COLON .tokens .pipeTokens? .blockSuiteDefinition? .SEMI_COLON
						---
							oneLine: "
								myBlock: .first ._second*;
							";

							oneLineWithTests: "
								myBlock: .first ._second*
									---
										validString: \"this is some input\";
									;
							";


							multipleLine: "
								myBlock: .first ._second*
									   | .OTHER+ .yes
									   ;
							";

							multipleLineWithTests: "
								myBlock: .first ._second*
									   | .OTHER+ .yes
									   ---
											validString: \"this is some input\";
										;
							";
						;

		blockSuiteDefinition: .suitePrefix .blockSuites
							---
								valid: "
									---
									validString: \"this is some input\";
									invalidString: !\"this is some input\";
									validBytes: [0, 22, 33];
									invalidBytes: ![0, 22, 33];
								";
					;

		blockSuites: .blockSuite+
					---
						valid: "
							validString: \"this is some input\";
							invalidString: !\"this is some input\";
							validBytes: [0, 22, 33];
							invalidBytes: ![0, 22, 33];
						";
					;

		blockSuite: .variableName .COLON .EXCLAMATION_POINT? .blockSuiteValue .SEMI_COLON
				---
					stringValue: "
						myName: \"this is some input\";
					";

					stringValueInvalid: "
						myName: !\"this is some input\";
					";

					bytesValue: "
						myName: [0, 22, 33];
					";

					bytesValueInvalid: "
						myName: ![0, 22, 33];
					";
				;

		blockSuiteValue: .stringValue
					   | .bytesList
					   ---
					   		stringValue: "
								\"this is some value\"
							";

							bytesList: "[0, 22, 33]";
					   ;

		pipeTokens: .PIPE .tokens
				---
					valid: "
						| !.MY_RULE [._conatnt].MY_RULE[4,] ![._conatnt].myBlock !._constant ._myReference+
					";
				;

		tokens: .token+
			---
				valid: "
					!.MY_RULE [._conatnt].MY_RULE[4,] ![._conatnt].myBlock !._constant ._myReference+
				";
			;

		token: .EXCLAMATION_POINT? .escapeElement? .blockElementReference .cardinality?
			---
				reverse: "!.MY_RULE";
				escapeCardinality: "[._conatnt].MY_RULE[4,]";
				cardinality: ".myBlock*";
				reference: "._myReference";
			;

		escapeElement: .BRACKET_OPEN .blockElementReference .BRACKET_CLOSE
					---
						rule: "[.MY_RULE]";
						block: "[.myBlock]";
						constant: "[._myConstant]";
					;

		blockElementReference: .DOT .blockElementName
							---
								rule: ".MY_RULE";
								block: ".myBlock";
								constant: "._myConstant";
							;

		blockElementName: .constantElementName
						| .variableName
						---
							constant: "_myConstant";
							rule: "MY_RULE";
							block: "myBlock";
						;

		constantDefinition: .constantName .COLON .constantElements .SEMI_COLON
							---
								valid: "
									_myConstant: .MY_RULE .OTHER_RULE[4] ._myConstant ._myConstant[2];
								";
						  	;

		constantElements: .constantElement+
						---
							valid: "
								.MY_RULE .OTHER_RULE[4] ._myConstant ._myConstant[2]
							";
						;

		constantElement: .constantElementReference .specificAmount?
					  ---
					  		rule: ".MY_RULE";
							ruleWithAmount: ".MY_RULE[4]";
							constant: "._myConstant";
							constantWithAmount: "._myConstant[4]";
					  ;

		constantElementReference: .DOT .constantElementName
								---
									constant: "._myConstant";
									rule: ".MY_RULE";
								;

		constantElementName: .constantName
						   | .ruleName
						   ---
								constant: "_myConstant";
								rule: "MY_RULE";
								block: !"myBlock";
						   ;

		constantName: .UNDERSCORE .variableName
					---
						valid: "_myConstant";
					;

		specificAmount: .BRACKET_OPEN .numbers .BRACKET_CLOSE
					---
						valid: "[2]";
					;

		cardinality: .BRACKET_OPEN .cardinalityValue .BRACKET_CLOSE
				   | .PLUS
				   | .STAR
				   | .INTERROGATION_POINT
					---
						value: "[0, 23]";
						plus: "+";
						star: "*";
						interrogationPoint: "?";
					;

		cardinalityValue: .numbers .commaNumbers
						| .numbersComma
						| .numbers
						---
							minWithMax: "0, 23";
							minWithoutMax: "0,";
							specific: "4";
						;


		numbersComma: .numbers .COMMA
					---
						valid: "0,";
					;

		ruleDefinition: .ruleName .COLON .ruleValue .SEMI_COLON
						---
							bytesList: "
								OTHER_RULE: [0, 1, 2, 3];
							";

							string: "
								MY_RULE: \"A\";
							";
						;

		ruleName: .oneUpperCaseLetter .ruleNameCharacter+
				---
					oneCharacter: !"A";
					valid: "MY_RULE";
					beginWithUnderscore: !"_A";
				;
		
		ruleNameCharacter: .oneUpperCaseLetter
						 | .UNDERSCORE
						 ---
						 	letter: "A";
							underscore: "_";
						 ;

		ruleValue: .bytesList
				 | .stringValue
				 ---
				 		string: "\"this is 13 \\\" string values!\"";
						bytes: "[1, 2, 3]";
				 ;

		bytesList: .BRACKET_OPEN .numbers .commaNumbers* .BRACKET_CLOSE
				---
					oneNumber: "[0]";
					list: "[0, 45, 33, 22]";
				;

		commaNumbers: .COMMA .numbers
					---
						valid: ", 45";
					;

		stringValue: .QUOTE ![.BACKSLASH].QUOTE .QUOTE
					---
						stringInQuotes: "\"this is 13 \\\" string values!\"";
					;

		schema: .head .pointWithTypes .connectionBlocks
			---
				valid: "
					head:
							engine: v1;
							name: myName;
							access: 
								read: .first;
								write: .first .second;
								review: .first;
							;
							compensation: 0.1, 0.23, 0.45;
					;

					son;
					father;
					grandFather;
					grandGrandFather;
					myList: list[float32];
					mySet: set[uint8];
					mySortedSet: sortedSet[uint8];

					father[son]: .son .father
								| .father .grandFather
								| .grandFather .grandGrandFather
								---
									mySuite[.son .grandGrandFather]:
										(.son .father .grandFather .grandGrandFather);
										!(.son .father .grandFather .grandGrandFather);
									;
								;

					grandFather[grandSon]: .son .grandFather
										| .father .grandGrandFather
										;
				";
			;

		connectionBlocks: .connectionBlock+
						---
							valid: "
								father[son]: .son .father
											| .father .grandFather
											| .grandFather .grandGrandFather
											---
												mySuite[.son .grandGrandFather]:
													(.son .father .grandFather .grandGrandFather);
													!(.son .father .grandFather .grandGrandFather);
												;
											;

								grandFather[grandSon]: .son .grandFather
													| .father .grandGrandFather
													;
							";
						;

		connectionBlock: .connectionName .COLON .links .pointSuitesBlock? .SEMI_COLON
						---
							withoutSuite: "
								father[son]: .son .father
										   | .father .grandFather
										   | .grandFather .grandGrandFather
										   ;
							";

							withSuites: "
									father[son]: .son .father
												| .father .grandFather
												| .grandFather .grandGrandFather
												---
														first[.son .grandGrandFather]:
															(.son .father .grandFather .grandGrandFather);
															!(.son .father .grandFather .grandGrandFather);
														;

														second[.son .grandGrandFather]:
															(.son .father .grandFather .grandGrandFather);
															!(.son .father .grandFather .grandGrandFather);
														;
												;
							";
						;

		pointSuitesBlock: .suitePrefix .pointSuites
						---
							valid: "
									---
										first[.son .grandGrandFather]:
											(.son .father .grandFather .grandGrandFather);
											!(.son .father .grandFather .grandGrandFather);
										;

										second[.son .grandGrandFather]:
											(.son .father .grandFather .grandGrandFather);
											!(.son .father .grandFather .grandGrandFather);
										;
							";
						;

		pointSuites: .pointSuite+
				---
					none: !"";
					single: "
							mySuite[.son .grandGrandFather]:
								(.son .father .grandFather .grandGrandFather);
								!(.son .father .grandFather .grandGrandFather);
							;
						";

					multiple: "
							first[.son .grandGrandFather]:
								(.son .father .grandFather .grandGrandFather);
								!(.son .father .grandFather .grandGrandFather);
							;

							second[.son .grandGrandFather]:
								(.son .father .grandFather .grandGrandFather);
								!(.son .father .grandFather .grandGrandFather);
							;
						";
				;

		pointSuite: .variableName .BRACKET_OPEN .pointReference[2] .BRACKET_CLOSE .COLON .pointSuiteOptionLine+ .SEMI_COLON
				---
					valid: "
						mySuite[.son .external[grandGrandFather]]:
							(.son .father .grandFather .grandGrandFather);
							!(.son .father .grandFather .grandGrandFather);
						;
					";
				;

		pointSuiteOptionLine: .pointSuiteOption .SEMI_COLON
							---
								expectedInvalid: "!(.son .father .grandFather .grandGrandFather);";
								expectedValid: "(.son .father .grandFather .grandGrandFather);";
							;

		pointSuiteOption: .pointInvalidLink
						| .pointOptimalLink
						| .pointReferencesInParenthesis
						---
							optimal: "> (.son .father .external[grandFather] .grandGrandFather)";
							expectedInvalid: "!(.son .father .external[grandFather] .grandGrandFather)";
							expectedValid: "(.son .father .grandFather .grandGrandFather)";
						;

		pointOptimalLink: .GREATHER_THAN .pointReferencesInParenthesis
						---
							valid: "> (.son .father .external[grandFather] .grandGrandFather)";
						;

		pointInvalidLink: .EXCLAMATION_POINT .pointReferencesInParenthesis
						---
							valid: "!(.son .father .external[grandFather] .grandGrandFather)";
						;

		pointReferencesInParenthesis: .PARENTHESIS_OPEN .pointReferences .PARENTHESIS_CLOSE
									---
										valid: "(.son .external[father] .grandFather .grandGrandFather)";
									;

		suitePrefix: .HYPHEN[3]
				---
					valid: "---";
				;

		links: .journey .pipeJourney*
			---
				oneLine: ".son .father";
				multipleLine: "
					.son .father
					| .myExternal[father] .grandFather (67)
					| .grandFather .myExternal[grandGrandFather]
					| .myExternal[grandFather] .myExternal[grandGrandFather]
					| .myExternal[grandFather] .grandGrandFather
				";
			;

		connectionName: .variableName .variableNameInBracket?
					---
						name: "father";
						nameWithReverse: "father[son]";
					;

		variableNameInBracket: .BRACKET_OPEN .variableName .BRACKET_CLOSE
							---
								valid: "[myVariable]";
							;

		pipeJourney: .PIPE .journey
					---
						valid: "| .son .father";
					;

		journey: .pointReference[2] .weight?
				---
					withWeight: ".origin .target (45)";
					withoutWeight: ".origin .target";
					withWeightWithExternal: ".myExternal[origin] .target (45)";
				;

		weight: .PARENTHESIS_OPEN .numbers .PARENTHESIS_CLOSE
			---
				zero: "(0)";
				multiple: "(45)";
				float: !"(34.98)";
			;

		pointWithTypes: .pointWithType+
					---
						valid: "
							son;
							father;
							grandFather;
							grandGrandFather;
							myList: list[uint8];
							mySet: set[uint8];
							mySortedSet: sortedSet[uint8];
						";
					;

		pointWithType: .variableName .colonPointType? .SEMI_COLON
					---
						point: "myPoint;";
						pointWithType: "mySecondPoint: list[uint8];";
					;

		colonPointType: .COLON .schemaTypeOption
					---
						list: ": list[float32]";
						set: ": set[uint8]";
						sortedSet: ": sortedSet[uint8]";
						map: ": map[uint8]";
					;

		head: .HEAD .COLON .headOptions .SEMI_COLON
			---
				all: "
						head:
							engine: v1;
							name: myName;
							access: 
								read: .first;
								write: .first .second;
								review: .first;
							;
							compensation: 0.1, 0.23, 0.45;
						;
				";
			;

		headOptions: .engine .propertyName .headOptionalOption*
				  ---
				  		mandatory: "
							engine: v1;
							name: myName;
						";

						access: "
							engine: v1;
							name: myName;
							access: 
								read: .first;
								write: .first .second;
								review: .first;
							;
						";

						compensation: "
							engine: v1;
							name: myName;
							compensation: 0.1, 0.23, 0.45;
						";

						all: "
							engine: v1;
							name: myName;
							compensation: 0.1, 0.23, 0.45;
							access: 
								read: .first;
								write: .first .second;
								review: .first;
							;
						";

						nameAtTheEnd: !"
							engine: v1;
							compensation: 0.1, 0.23, 0.45;
							name: myName;
						";
				  ;

		headOptionalOption:	.access
							| .compensation
							---
									access: "
										access: 
											read: .first;
											write: .first .second;
											review: .first;
										;
									";

									compensation: "
										compensation: 0.1, 0.23, 0.45;
									";
							;

		engine: .ENGINE .COLON .version .SEMI_COLON
				---
					versionOneNumber: "engine: v1;";
					versionWithMultipleNumbers: "engine: v123;";
				;

		compensation: .COMPENSATION .COLON .threeFloatNumbersBetweenZeroAndOne .SEMI_COLON
					---
						valid: "compensation: 0.1, 0.23, 0.45;";
					;

		access: .ACCESS .COLON .permissionOptions .SEMI_COLON
			---
				allOptionsWithDuplicates: "
					access: 
						review: .first;
						write: .first .second;
						review: .first;
						review: .first .second .third;
						review: .first .second .third;
					;
				";

				allOptions: "
					access: 
						read: .first;
						write: .first .second;
						review: .first;
					;
				";

				readOption: "
					access: 
						read: .first;
					;
				";

				writeOption: "
					access: 
						write: .first .second;
					;
				";

				reviewOption: "
					access: 
						review: .first;
					;
				";
			;

		permissionOptions: .permissionOption+
						---
							reviewOption: "
								review: .first;
							";

							writeOption: "
								write: .first;
							";

							readOption: "
								read: .first;
							";

							threeOptions: "
								read: .first;
								write: .first .second;
								review: .first .second .third;
							";

							threeOptionsWithDuplicates: "
								review: .first;
								write: .first .second;
								review: .first;
								review: .first .second .third;
								review: .first .second .third;
							";
						;

		permissionOption: .propertyReview
						| .propertyWrite 
						| .propertyRead
						---
							review: "review: .first;";
							write: "write: .first .second;";
							read: "review: .first .second .third;";
						;

		propertyReview: .REVIEW .COLON .references .SEMI_COLON
					---
						oneRole: "review: .myRole;";
						multipleRoles: "review: .first .second .third;";
					;

		propertyWrite: .WRITE .COLON .references  .SEMI_COLON
					---
						oneRole: "write: .myRole;";
						multipleRoles: "write: .first .second .third;";
					;

		propertyRead: .READ .COLON .references .SEMI_COLON
					---
						oneRole: "read: .myRole;";
						multipleRoles: "read: .first .second .third;";
					;

		propertyName: .NAME .COLON .variableName .SEMI_COLON
					---
						valid: "name: myName;";
					;

		version: .LL_V .numbersExceptZero
				---
					versionZero: !"v0";
					versionOneNumber: "v1";
					versionWithMultipleNumbers: "v123";
				;

		references: .reference+
				  ---
				  		oneReference: ".myReference";
						multipleReferences: ".myReference .secondReference .thirdReference";
				  ;

		reference: .DOT .variableName
				---
					withDot: ".myReference";
					withoutDot: !"myReference";
					withoutVariable: !".";
				;

		variableName: .oneLowerCaseLetter .anyLetters*
					---
						oneLetter: "m";
						lowercaseLetters: "myvariable";
						multipleLetters: "myVariable";
						upperCaseFirstLetter: !"MyVariable";
						uppercaseLetters: !"MYVARIABLE";
					;

		anyLetters: .upperCaseLetters
				  | .lowerCaseLetters
				  ---
				  		upperCaseLetters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
						lowerCaseLetters: "abcdefghijklmnopqrstuvwxyz";
				  ;

		upperCaseLetters: .oneUpperCaseLetter+
						---
							upperCaseLetters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
							lowerCaseLetters: !"abcdefghijklmnopqrstuvwxyz";
						;
		
		lowerCaseLetters: .oneLowerCaseLetter+
						---
							lowerCaseLetters: "abcdefghijklmnopqrstuvwxyz";
							upperCaseLetters: !"ABCDEFGHIJKLMNOPQRSTUVWXYZ";
						;

		oneUpperCaseLetter: .UL_A
						  | .UL_B
						  | .UL_C
						  | .UL_D
						  | .UL_E
						  | .UL_F
						  | .UL_G
						  | .UL_H
						  | .UL_I
						  | .UL_J
						  | .UL_K
						  | .UL_L
						  | .UL_M
						  | .UL_N
						  | .UL_O
						  | .UL_P
						  | .UL_Q
						  | .UL_R
						  | .UL_S
						  | .UL_T
						  | .UL_U
						  | .UL_V
						  | .UL_W
						  | .UL_X
						  | .UL_Y
						  | .UL_Z
						  ---
						  	a: "A";
							b: "B";
							c: "C";
							d: "D";
							e: "E";
							f: "F";
							g: "G";
							h: "H";
							i: "I";
							j: "J";
							k: "K";
							l: "L";
							m: "M";
							n: "N";
							o: "O";
							p: "P";
							q: "Q";
							r: "R";
							s: "S";
							t: "T";
							u: "U";
							v: "V";
							w: "W";
							x: "X";
							y: "Y";
							z: "Z";
						  ;

		oneLowerCaseLetter: .LL_A
						  | .LL_B
						  | .LL_C
						  | .LL_D
						  | .LL_E
						  | .LL_F
						  | .LL_G
						  | .LL_H
						  | .LL_I
						  | .LL_J
						  | .LL_K
						  | .LL_L
						  | .LL_M
						  | .LL_N
						  | .LL_O
						  | .LL_P
						  | .LL_Q
						  | .LL_R
						  | .LL_S
						  | .LL_T
						  | .LL_U
						  | .LL_V
						  | .LL_W
						  | .LL_X
						  | .LL_Y
						  | .LL_Z
						  ---
						  	a: "a";
							b: "b";
							c: "c";
							d: "d";
							e: "e";
							f: "f";
							g: "g";
							h: "h";
							i: "i";
							j: "j";
							k: "k";
							l: "l";
							m: "m";
							n: "n";
							o: "o";
							p: "p";
							q: "q";
							r: "r";
							s: "s";
							t: "t";
							u: "u";
							v: "v";
							w: "w";
							x: "x";
							y: "y";
							z: "z";
						  ;

		threeFloatNumbersBetweenZeroAndOne: .floatNumberBetweenZeroAndOne .commaFloatNumberBetweenZeroAndOne[2]
										---
											three: "0.0, 0.12, 0.134";
											one: !"0.12";
											two: !"0.12, 23.32";
											lastTooBigThree: !"0.0, 0.12, 32.134";
										;

		commaFloatNumberBetweenZeroAndOne: .COMMA .floatNumberBetweenZeroAndOne
										---
											valid: ", 0.12";
											invalid: !", 32.98";
										; 

		floatNumberBetweenZeroAndOne: .N_ZERO .DOT .numbers
									---
										lessThanZero: "0.23";
										zero: "0.0";
										pointZero: !"4556.0";
									;

		numbers: .oneNumber+
				---
					oneNumber: "1";
					numberWithAllNumbers: "1234567890";
					negativeNumberWithAllNumbers: !"-1234567890";
					oneLettter: !"a";
				;

		numbersExceptZero: .oneNumberExceptZero+
				---
					oneNumber: "1";
					numberWithAllNumbers: !"1234567890";
					numberWithAllNumbersExceptZero: "123456789";
					negativeNumberWithAllNumbers: !"-1234567890";
					oneLettter: !"a";
					numberZero: !"0";
				;

		oneNumber: .N_ZERO
				 | .oneNumberExceptZero
				 ---
				   	zero: "0";
					one: "1";
					two: "2";
					three: "3";
					four: "4";
					five: "5";
					six: "6";
					seven: "7";
					height: "8";
					nine: "9";
				;

		oneNumberExceptZero: .N_ONE
							| .N_TWO
							| .N_THREE
							| .N_FOUR
							| .N_FIVE
							| .N_SIX
							| .N_SEVEN
							| .N_HEIGHT
							| .N_NINE
							---
								zero: !"0";
								one: "1";
								two: "2";
								three: "3";
								four: "4";
								five: "5";
								six: "6";
								seven: "7";
								height: "8";
								nine: "9";
							;

		pointReferences: .pointReference+
						---
							valid: ".mySchema[myPoint] .myPoint";
						;

		pointReference: .externalPointReference
					  | .reference
					  ---
							external: ".mySchema[myPoint]";
							internal: ".myPoint";
						;

		externalPointReference: .reference .BRACKET_OPEN .variableName .BRACKET_CLOSE
							---
								valid: ".mySchema[myPoint]";
							;

		schemaTypeOption: .reference
						| .typeOption
						---
							container: "set[float32]";
							primitive: "float32";
							reference: ".myGrammar";
						;

		typeOption: .containerType
				  | .primitiveType
					---
				   		container: "list[float32]";
						primitive: "float32";
				   ;

		containerType: .containerName .BRACKET_OPEN .typeOption .BRACKET_CLOSE
					---
						float: "list[float32]";
						int: "list[int32]";
						uint: "list[uint8]";
						string: "list[string]";
						bool: "list[bool]";
						listInLists: "list[list[list[float32]]]";
						setInLists: "list[sortedSet[set[float32]]]";
					;

		containerName: .LIST
					 | .SET
					 | .SORTED_SET
					 | .MAP
					 ---
					 	list: "list";
						set: "set";
						sortedSet: "sortedSet";
						map: "map";
					 ;

		primitiveType: .floatType
					| .intType
					| .STRING
					| .BOOL
					---
						float32: "float32";
						float64: "float64";
						int8: "uint8";
						int16: "uint16";
						int32: "uint32";
						int64: "uint64";
						uint8: "uint8";
						uint16: "uint16";
						uint32: "uint32";
						uint64: "uint64";
						string: "string";
						bool: "bool";
					;

		floatType: .FLOAT .floatSize
				---
					float32: "float32";
					float64: "float64";
				;

		intType: .intTypeOption .intSize
				---
					int8: "uint8";
					int16: "uint16";
					int32: "uint32";
					int64: "uint64";
					uint8: "uint8";
					uint16: "uint16";
					uint32: "uint32";
					uint64: "uint64";
				;

		intTypeOption: .INT
					| .UINT
					---
						int: "int";
						uint: "uint";
					;

		intSize: .N_HEIGHT
				| .sixteen
				| .thirtyTwo
				| .sixtyFour
				---
					height: "8";
					sixteen: "16";
					thirtyTwo: "32";
					sixtyFour: "64";
				;

		floatSize: .thirtyTwo
				| .sixtyFour
				---
					thirtyTwo: "32";
					sixtyFour: "64";
				;

		sixtyFour: .N_SIX .N_FOUR
				---
					valid: "64";
				;

		thirtyTwo: .N_THREE .N_TWO
				---
					valid: "32";
				;

		sixteen: .N_ONE .N_SIX
				---
					valid: "16";
				;
							
		SPACE: " ";
		TAB: "	";
		EOL: "
";

		N_ZERO: "0";
		N_ONE: "1";
		N_TWO: "2";
		N_THREE: "3";
		N_FOUR: "4";
		N_FIVE: "5";
		N_SIX: "6";
		N_SEVEN: "7";
		N_HEIGHT: "8";
		N_NINE: "9";

		LL_A: "a";
		LL_B: "b";
		LL_C: "c";
		LL_D: "d";
		LL_E: "e";
		LL_F: "f";
		LL_G: "g";
		LL_H: "h";
		LL_I: "i";
		LL_J: "j";
		LL_K: "k";
		LL_L: "l";
		LL_M: "m";
		LL_N: "n";
		LL_O: "o";
		LL_P: "p";
		LL_Q: "q";
		LL_R: "r";
		LL_S: "s";
		LL_T: "t";
		LL_U: "u";
		LL_V: "v";
		LL_W: "w";
		LL_X: "x";
		LL_Y: "y";
		LL_Z: "z";
		
		UL_A: "A";
		UL_B: "B";
		UL_C: "C";
		UL_D: "D";
		UL_E: "E";
		UL_F: "F";
		UL_G: "G";
		UL_H: "H";
		UL_I: "I";
		UL_J: "J";
		UL_K: "K";
		UL_L: "L";
		UL_M: "M";
		UL_N: "N";
		UL_O: "O";
		UL_P: "P";
		UL_Q: "Q";
		UL_R: "R";
		UL_S: "S";
		UL_T: "T";
		UL_U: "U";
		UL_V: "V";
		UL_W: "W";
		UL_X: "X";
		UL_Y: "Y";
		UL_Z: "Z";

		DOT: ".";
		COLON: ":";
		SEMI_COLON: ";";
		COMMA: ",";

		HEAD: "head";
		NAME: "name";
		READ: "read";
		WRITE: "write";
		REVIEW: "review";
		ACCESS: "access";
		COMPENSATION: "compensation";
		ENGINE: "engine";
		POINT: "point";
		LIST: "list";
		SET: "set";
		MAP: "map";
		SORTED_SET: "sortedSet";
		FLOAT: "float";
		INT: "int";
		UINT: "uint";
		STRING: "string";
		BOOL: "bool";
		ENTRY: "entry";
		OMIT: "omit";
		ORIGIN: "origin";
		INPUT: "input";
		EXPECTED: "expected";
		RECIPE: "recipe";
		PROGRAM: "program";

		BRACKET_OPEN: "[";
		BRACKET_CLOSE: "]";
		PARENTHESIS_OPEN: "(";
		PARENTHESIS_CLOSE: ")";
		PIPE: "|";
		HYPHEN: "-";
		EXCLAMATION_POINT: "!";
		GREATHER_THAN: ">";
		UNDERSCORE: "_";
		QUOTE: "\"";
		BACKSLASH: "\\";
		STAR: "*";
		PLUS: "+";
		INTERROGATION_POINT: "?";
		COMMERCIAL_A: "@";

		
	`)
}
