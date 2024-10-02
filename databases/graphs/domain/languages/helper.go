package languages

func fetchGrammarInput() []byte {
	return []byte(`
		v1;
		> .root;
		# .SPACE .TAB .EOL;

		root: .schema
			---
				valid: "
					v1;
					name: myName;

					son;
					father;
					grandFather;
					grandGrandFather;

					father(son): .son .father
								| .son .myExternal[father]
								| .father .grandFather
								| .grandFather .grandGrandFather
								---
									mySuite[.son .grandGrandFather]:
										> (.son .father .external[grandFather] .grandGrandFather);
										(.son .father .grandFather .grandGrandFather);
										!(.son .father .grandFather .grandGrandFather);
									;
								;

					grandFather(grandSon): .son .grandFather
										| .father .grandGrandFather
										;
				";
			;

		schema: .head .instructionPoints .connectionBlocks
			---
				valid: "
					v1;
					name: myName;

					son;
					father;
					grandFather;
					grandGrandFather;

					father[0,3](son+): .son .father
								| .father .grandFather
								| .grandFather .grandGrandFather
								---
									mySuite[.son .grandGrandFather]:
										> (.son .father .external[grandFather] .grandGrandFather);
										(.son .father .grandFather .grandGrandFather);
										!(.son .father .grandFather .grandGrandFather);
									;
								;

					grandFather(grandSon[2,]): .son .grandFather
										| .father .grandGrandFather
										;
				";
			;

		head: .versionInstruction .nameInstruction
			---
				valid: "
						v1;
						name: myName;
				";
			;

		nameInstruction: .NAME .COLON .variableName .SEMI_COLON
						---
							valid: "
								name: myName;
							";
						;

		versionInstruction: .LL_V .numbers .SEMI_COLON
						---
							versionOneNumber: "v1;";
							versionWithMultipleNumbers: "v123;";
						;

		connectionBlocks: .connectionBlock+
						---
							valid: "
								father(son): .son .father
											| .father .grandFather
											| .grandFather .grandGrandFather
											---
												mySuite[.son .grandGrandFather]:
													(.son .father .grandFather .grandGrandFather);
													!(.son .father .grandFather .grandGrandFather);
												;
											;

								grandFather(grandSon): .son .grandFather
													| .father .grandGrandFather
													;
							";
						;

		connectionBlock: .connectionName .COLON .links .pointSuitesBlock? .SEMI_COLON
						---
							withoutSuite: "
								father(son): .son .father
										   | .father .grandFather
										   | .grandFather .grandGrandFather
										   ;
							";

							withSuites: "
									father(son): .son .father
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

		connectionName: .nameWithCardinality .variableNameInBracket?
					---
						name: "father";
						withCardinality: "father+";
						nameWithReverse: "father(son)";
					;

		variableNameInBracket: .PARENTHESIS_OPEN .nameWithCardinality .PARENTHESIS_CLOSE
							---
								valid: "(myVariable)";
								withCardinality: "(myVariable+)";
								withBracketCardinality: "(myVariable[0,])";
							;

		nameWithCardinality: .variableName .cardinality?
							---
								name: "son";
								withCardinality: "son+";
								withBracketCardinality: "son[3,]";
							;

		cardinality: .cardinalityBracketOption
					| .INTERROGATION_POINT
					| .STAR
					| .PLUS
					---
							minMax: "[0,23]";
							minNoMax: "[0,]";
							specific: "[22]";
							interrogationPoint: "?";
							star: "*";
							plus: "+";
					;

		cardinalityBracketOption: .BRACKET_OPEN .cardinalityNumberOptions .BRACKET_CLOSE
								---
										minMax: "[0,23]";
										minNoMax: "[0,]";
										specific: "[22]";
								;

		cardinalityNumberOptions: .minMax
								| .minComma
								| .numbers
								---
										minMax: "0,23";
										minNoMax: "0,";
										specific: "22";
								;


		minMax: .minComma .numbers
			  ---
			  		valid: "0,23";
			  ;

		minComma: .numbers .COMMA
				---
					zero: "0,";
					nonZero: "23,";
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
					single: "
							mySuite[.son .grandGrandFather]:
								(.son .father .grandFather .grandGrandFather);
								!(.son .father .grandFather .grandGrandFather);
							;
						";

					multiple: "
							first[.son .grandGrandFather]:
								> (.son .father .external[grandFather] .grandGrandFather);
								(.son .father .grandFather .grandGrandFather);
								!(.son .father .grandFather .grandGrandFather);
							;

							second[.son .grandGrandFather]:
								> (.son .father .external[grandFather] .grandGrandFather);
								(.son .father .grandFather .grandGrandFather);
								!(.son .father .grandFather .grandGrandFather);
							;
						";
				;

		pointSuite: .variableName .BRACKET_OPEN .pointReference[2] .BRACKET_CLOSE .COLON .suiteInstructions .SEMI_COLON
				---
					valid: "
						mySuite[.son .external[grandGrandFather]]:
							> (.son .father .external[grandFather] .grandGrandFather);
							!(.son .father .external[grandFather] .grandGrandFather);
							(.son .father .external[grandFather] .grandGrandFather);
						;
					";
				;

		suiteInstructions: .suiteOptimalInstruction? .suiteUnlimitedOptionInstruction+
						---
							withOptimalLink: "
								> (.son .father .external[grandFather] .grandGrandFather);
								!(.son .father .external[grandFather] .grandGrandFather);
								(.son .father .external[grandFather] .grandGrandFather);
							";

							withoutOptimalLink: "
								!(.son .father .external[grandFather] .grandGrandFather);
								(.son .father .external[grandFather] .grandGrandFather);
								(.son .father .external[grandFather] .grandGrandFather);
								(.son .father .external[grandFather] .grandGrandFather);
							";
						;

		suiteOptimalInstruction: .suiteOptimalLink .SEMI_COLON;

		suiteUnlimitedOptionInstruction: .suiteUnlimitedOption .SEMI_COLON;

		suiteUnlimitedOption: .suiteInvalidLink
							| .suiteReferencesInParenthesis
							---
								invalidLink: "!(.son .father .external[grandFather] .grandGrandFather)";
								validLink: "(.son .father .external[grandFather] .grandGrandFather)";
							;

		suiteOptimalLink: .GREATHER_THAN .suiteReferencesInParenthesis
						---
							valid: "> (.son .father .external[grandFather] .grandGrandFather)";
						;

		suiteInvalidLink: .EXCLAMATION_POINT .suiteReferencesInParenthesis
						---
							valid: "!(.son .father .external[grandFather] .grandGrandFather)";
						;

		suiteReferencesInParenthesis: .PARENTHESIS_OPEN .pointReferences .PARENTHESIS_CLOSE
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

		weight: .PARENTHESIS_OPEN .positiveNumbers .PARENTHESIS_CLOSE
				---
					zero: "(0)";
					multiple: "(45)";
					float: !"(34.98)";
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

		instructionPoints: .instructionPoint+
						---
							single: "
								son;
							";

							multiple: "
								son;
								father;
								grandFather;
								grandGrandFather;
							";
						;

		instructionPoint: .variableName .SEMI_COLON
						---
							valid: "son;";
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

		positiveNumbers: .oneNumber+
						---
							oneNumber: "1";
							numberWithAllNumbers: "1234567890";
							negativeNumberWithAllNumbers: !"-1234567890";
							oneLettter: !"a";
						;

		numbers: .oneNumber+
				---
					single: "0";
					multiple: "345";
				;

		oneNumber: .N_ZERO
					| .N_ONE
					| .N_TWO
					| .N_THREE
					| .N_FOUR
					| .N_FIVE
					| .N_SIX
					| .N_SEVEN
					| .N_HEIGHT
					| .N_NINE
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
		COMMA: ",";
		SEMI_COLON: ";";
		PARENTHESIS_OPEN: "(";
		PARENTHESIS_CLOSE: ")";
		BRACKET_OPEN: "[";
		BRACKET_CLOSE: "]";
		PIPE: "|";
		HYPHEN: "-";
		EXCLAMATION_POINT: "!";
		GREATHER_THAN: ">";
		INTERROGATION_POINT: "?";
		STAR: "*";
		PLUS: "+";

		NAME: "name";
	`)
}
