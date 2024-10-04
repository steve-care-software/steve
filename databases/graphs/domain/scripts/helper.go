package scripts

func fetchGrammarInput() []byte {
	return []byte(`
		v1;
		> .reference;
		# .SPACE .TAB .EOL;

		operation: .assignable .operatorAssignable*
				 | .PARENTHESIS_OPEN .operation .PARENTHESIS_CLOSE
				---
					complex: "
							(
								(
									(myVariable < 23) && (myValue + 12)
								) || (
								 	myVariable * (
										(34 * 12 + 46) + 54 * 21
									)
								 ) && myValue
							)
					";
				;

		operatorAssignable: .operator .assignable;

		operator: .arithmeticOperator
				| .relationalOperator
				| .equalOperator
				| .logicalOperator
				;

		arithmeticOperator: .PLUS
						  | .HYPHEN
						  | .STAR
						  | .SLASH
						  | .PERCENT
						  ;

		relationalOperator: .SMALLER_THAN .EQUAL?
						  | .GREATHER_THAN .EQUAL?
						  ;

		equalOperator: .EXCLAMATION_POINT .EQUAL
					 | .EQUAL[2]
					 ;

		assignment: .assigneesCommaSeparated .assignmentSymbol .assignableOptions .SEMI_COLON
				---
					simple: "
						uint8 myVariable := 8;
					";

					complex: "
						uint8 myVariable, myAlreadyDeclared, map myMap, float32 myFloat :=
							8,
							32,
							myInt.(float32),
							[
								myKeyname: 34;
								again: [
									voila: true;
									again: [
										number: 23;
										again: 23;
									];
									third: true;
								];
							],
							myInt.(float32)
						;
					";

					querySelect: "
							select mySelect :=
								.mySchema[myPoint];
								.mySchema[mySecondPoint];
								condition:
									(
										.mySchema[myPoint]: @myVariable <> 
											.mySchema[secondPoint]: @myOtherVar
									) && (
											.mySchema[other]: @myVariable || (
												.mySchema[myPoint]: @myVariable && 
													.mySchema[secondPoint]: @myOtherVar
											)
										)
								;
							;
						";

					bridges: "
							bridges myBridge :=
								22:
									.myOriginSchema[myPoint]: @myOtherValue;
									.myTargetSchea[myPoint]: @myValue;
								;

								43:
									.myOriginSchema[myPoint]: @myOtherValue;
									.myTargetSchea[myPoint]: @myValue;
								;
							;
						";

					map: "
						map myMap := [
							myKeyname: 34;
							again: [
								voila: true;
								again: [
									number: 23;
									again: 23;
								];
								third: true;
							];
						];
					";

					casting: "
						float32 myFloat := myInt.(float32);
					";
				;

		assignmentSymbol: .COLON? .EQUAL
						---
							assignment: ":=";
							reAssignment: "=";
						;

		assigneesCommaSeparated: .assignee .commaAssignee*
								---
									single: "uint8 myVariable";
									list: "uint8 myVariable, myVariable, uint8 myThird";
								;

		commaAssignee: .COMMA .assignee
					---
						valid: ", myVariable";
					;

		assignee: .type? .variableName
				---
					variable: "myVariable";
					variableWithType: "uint8 myVariable";
				;

		assignableOptions: .assignable .commaAssignable*;
		commaAssignable: .COMMA .assignable;

		assignable: .queryInsertUpdate
				  | .queryDelete
				  | .querySelect
				  | .queryBridges
				  | .map
				  | .programCall
				  | .primitiveValue
				  | .variableCasting
				  | .variableName
				  | .PARENTHESIS_OPEN .operation .PARENTHESIS_CLOSE
					---

						queryDelete: "
								.mySchema[myPoint];
								.myOtherSchema[myOtherPoint];
							
						";

						queryInsertUpdate: "
								.mySchema[myPoint]:
									.mySubSchema[mySubPoint]: @myVariable;
									.mySubSchema[other]: @myVariable;
									.mySubSchema[third]: 
										.mySubSchema[mySubPoint]: @myVariable;
									;
								;
							
						";

						queryBridge: "
								22:
									.myOriginSchema[myPoint]: @myOtherValue;
									.myTargetSchea[myPoint]: @myValue;
								;

								43:
									.myOriginSchema[myPoint]: @myOtherValue;
									.myTargetSchea[myPoint]: @myValue;
								;
							
						";

						querySelect: "
								.mySchema[myPoint];
								.mySchema[mySecondPoint];
								condition:
									(
										.mySchema[myPoint]: @myVariable <> 
											.mySchema[secondPoint]: @myOtherVar
									) && (
											.mySchema[other]: @myVariable || (
												.mySchema[myPoint]: @myVariable && 
													.mySchema[secondPoint]: @myOtherVar
											)
										)
								;
							
						";

						programCall: "
							.first .second .third: [
								myKeyname: 34;
								again: [
									voila: true;
									again: [
										number: 23;
										again: 23;
									];
									third: true;
								];
							]
						";

						map: "
							[
								myKeyname: 34;
								again: [
									voila: true;
									again: [
										number: 23;
										again: 23;
									];
									third: true;
								];
							]
						";

						primitive: "34.0";
						variable: "myVariable";
						casting: "myVariable.(float32)";
					;

		variableCasting: .variableName .DOT .PARENTHESIS_OPEN .type .PARENTHESIS_CLOSE
						---
							valid: "
								myVariable.(set[uint8])
							";
						;

		programCall: .COMMERCIAL_A? .references .colonMap?
					---
						engine: "
							@.first .second .third
						";

						simple: "
							.first .second .third
						";

						withParams: "
							.first .second .third: [
								myKeyname: 34;
								again: [
									voila: true;
									again: [
										number: 23;
										again: 23;
									];
									third: true;
								];
							]
						";
					;

		colonMap: .COLON .map;

		map: .BRACKET_OPEN .keyValues .BRACKET_CLOSE
			---
				valid: "
					[
						myKeyname: 34;
						again: [
							voila: true;
							again: [
								number: 23;
								again: 23;
							];
							third: true;
						];
					]
				";
			;

		keyValues: .keyValue+
				---
					valid: "
						myKeyname: 34;
						again: [
							voila: true;
							again: [
								number: 23;
								again: 23;
							];
							third: true;
						];
					";
				;



		keyValue: .variableName .COLON .assignable .SEMI_COLON
				---
					simple: "
						myKeyname: 34;
					";

					complex: "
						myKeyname: [
							second: [
								third: 45;
							];
						];
					";
				;

		primitiveValue: .numericValue
					  | .boolValue
					  | .stringValue
					  ---
					  		numeric: "34.0";
							bool: "true";
							string: "\"this is 13 string values!\"";
					  ;

		numericValue: .floatValue
					| .intValue
					---
						negativeFloat: "-34.0";
						float: "34.0";

						negativeInt: "-34";
						int: "34";
					;

		boolValue: .TRUE
				| .FALSE
				---
					true: "true";
					false: "false";
				;

		stringValue: .QUOTE ![.BACKSLASH].QUOTE .QUOTE
					---
						stringInQuotes: "\"this is 13 \\\" string values!\"";
					;

		floatValue: .intValue .DOT .numbers
				---
					negative: "-34.0";
					positive: "34.0";
				;

		intValue: .HYPHEN? .numbers
				---
					negative: "-45";
					positive: "-45";
				;

		type: .engineType
			| .primitiveType
			| .containerType
			| .MAP
			---
				engine: "selector";
				primitive: "bool";
				container: "set[string]";
				complex: "set[list[set[sortedSet[string]]]]";
				map: "map";
			;

		containerType: .containerName .BRACKET_OPEN .type .BRACKET_CLOSE
					---
						complex: "set[list[set[sortedSet[string]]]]";
					;
		
		containerName: .LIST
					| .SET
					| .SORTED_SET
					---
						list: "list";
						set: "set";
						sortedSet: "sortedSet";
					;

		engineType: .SELECTOR
				  | .ROUTE
				  | .SELECT
				  | .INSERT
				  | .UPDATE
				  | .DELETE
				  | .BRIDGES
				  ---
						selector: "selector";
						select: "select";
						route: "route";
						insert: "insert";
						update: "update";
						delete: "delete";
						bridges: "bridges";
				  ;
		
		primitiveType: .numericType
					| .BOOL
					| .STRING
					---
						numeric: "int16";
						bool: "bool";
						string: "string";
					;

		numericType: .floatType
					| .intType
					| .uintType
					---
						int: "int8";
						uint: "uint16";
						float: "float32";
					;

		floatType: .FLOAT .floatSize
				---
					thirtyTwo: "float32";
					sixtyFour: "float64";
				;
		
		floatSize: .thirtyTwo
				| .sixtyFour
				---
					thirtyTwo: "32";
					sixtyFour: "64";
				;

		intType: .INT .intSize
				---
					height: "int8";
					sixteen: "int16";
					thirtyTwo: "int32";
					sixtyFour: "int64";
				;

		uintType: .UINT .intSize
				---
					height: "uint8";
					sixteen: "uint16";
					thirtyTwo: "uint32";
					sixtyFour: "uint64";
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

		queryDelete: .pointReferenceLines .condition?
					---
						withCondition: "
								.mySchema[myPoint];
								condition:
									.mySchema[myPoint]: @myVariable && 
										.mySchema[secondPoint]: @myOtherVar
								;
							
						";

						withoutCondition: "
								.mySchema[myPoint];
								.myOtherSchema[myOtherPoint];
							
						";
					;

		queryInsertUpdate: .queryAssignment .condition?
						---
							withCondition: "
									.mySchema[myPoint]:
										.mySubSchema[mySubPoint]: @myVariable;
										.mySubSchema[other]: @myVariable;
										.mySubSchema[third]: 
											.mySubSchema[mySubPoint]: @myVariable;
										;
									;

									condition:
										(
											.mySchema[myPoint]: @myVariable <> 
												.mySchema[secondPoint]: @myOtherVar
										) && (
												.mySchema[other]: @myVariable || (
													.mySchema[myPoint]: @myVariable && 
														.mySchema[secondPoint]: @myOtherVar
												)
											)
									;
								
							";

							withoutCondition: "
									.mySchema[myPoint]:
										.mySubSchema[mySubPoint]: @myVariable;
										.mySubSchema[other]: @myVariable;
										.mySubSchema[third]: 
											.mySubSchema[mySubPoint]: @myVariable;
										;
									;
								
							";
						;

		queryAssignment: .externalPointReference .COLON .queryAssignable+ .SEMI_COLON
						---
							variable: "
								.mySchema[myPoint]: @myVariable;
							";

							withSubAssignment: "
								.mySchema[myPoint]:
									.mySubSchema[mySubPoint]: @myVariable;
								;
							";

							withSubAssignments: "
								.mySchema[myPoint]:
									.mySubSchema[mySubPoint]: @myVariable;
									.mySubSchema[other]: @myVariable;
									.mySubSchema[third]: 
										.mySubSchema[mySubPoint]: @myVariable;
									;
								;
							";
						;

		queryAssignable: .queryVariableName
						| .queryAssignment
						;

		querySelect:.pointReferenceLines .condition?
				---
					withCondition: "
								.mySchema[myPoint];
								.mySchema[mySecondPoint];
								condition:
									(
										.mySchema[myPoint]: @myVariable <> 
											.mySchema[secondPoint]: @myOtherVar
									) && (
											.mySchema[other]: @myVariable || (
												.mySchema[myPoint]: @myVariable && 
													.mySchema[secondPoint]: @myOtherVar
											)
										)
								;
							
					";

					withoutCondition: "
								.mySchema[myPoint];
								.mySchema[mySecondPoint];
							
					";
				;

		condition: .CONDITION .COLON .conditionClause .SEMI_COLON
				---
					valid: "
						condition: 
							(
								.mySchema[myPoint]: @myVariable <> 
									.mySchema[secondPoint]: @myOtherVar
							) && (
									.mySchema[other]: @myVariable || (
										.mySchema[myPoint]: @myVariable && 
											.mySchema[secondPoint]: @myOtherVar
									)
								)
						;
					";
				;

		conditionClause: .conditionElement .logicalOperator .conditionElement
						---
							and: "
								.mySchema[myPoint]: @myVariable && .mySchema[secondPoint]: @myOtherVar
							";

							or: "
								.mySchema[myPoint]: @myVariable || .mySchema[secondPoint]: @myOtherVar
							";

							xor: "
								.mySchema[myPoint]: @myVariable <> .mySchema[secondPoint]: @myOtherVar
							";

							complex: "
								(
									.mySchema[myPoint]: @myVariable <> 
										.mySchema[secondPoint]: @myOtherVar
								) && (
										.mySchema[other]: @myVariable || (
											.mySchema[myPoint]: @myVariable && 
												.mySchema[secondPoint]: @myOtherVar
										)
									)
							";
						;

		conditionElement: .PARENTHESIS_OPEN .conditionClause .PARENTHESIS_CLOSE
						| .queryVariable
						---
							variable: "
								.mySchema[myPoint]: @myVariable
							";

							logical: "
								(.mySchema[myPoint]: @myVariable && .mySchema[secondPoint]: @myOtherVar)
							";

							complex: "
								(
									(
										.mySchema[myPoint]: @myVariable <> 
											.mySchema[secondPoint]: @myOtherVar
									) && (
											.mySchema[other]: @myVariable || (
												.mySchema[myPoint]: @myVariable && 
													.mySchema[secondPoint]: @myOtherVar
											)
										)
								)
							";
						;
		
		logicalOperator: .AND
						| .OR
						| .XOR
						---
							and: "&&";
							or: "||";
							xor: "<>";
						;

		queryBridges: .bridgeWeight+
					---
						valid: "
								22:
									.myOriginSchema[myPoint]: @myOtherValue;
									.myTargetSchea[myPoint]: @myValue;
								;

								43:
									.myOriginSchema[myPoint]: @myOtherValue;
									.myTargetSchea[myPoint]: @myValue;
								;
						";
					;

		bridgeWeight: .numbers .COLON .queryVariableLine[2] .SEMI_COLON
					---
						valid: "
							22:
								.myOriginSchema[myPoint]: @myOtherValue;
								.myTargetSchea[myPoint]: @myValue;
							;
						";
					;

		queryVariableLine: .queryVariable .SEMI_COLON
						---
							valid: ".mySchema[myPoint]: @myVariable;";
						;

		queryVariable: .externalPointReference .COLON .queryVariableName
							---
								valid: ".mySchema[myPoint]: @myVariable";
							;

		queryVariableName: .COMMERCIAL_A .variableName
						---
							valid: "@myVariable";
						;

		pointReferenceLines: .pointReferenceLine+
							---
								valid: "
									.mySchema[myPoint];
									.mySecondSchema[secondPoint];
								";
							;

		pointReferenceLine: .externalPointReference .SEMI_COLON
							---
								external: ".mySchema[myPoint];";
								internal: !".myPoint;";
							;

		externalPointReference: .reference .BRACKET_OPEN .variableName .BRACKET_CLOSE
							---
								valid: ".mySchema[myPoint]";
							;

		references: .reference+
					---
						valid: "
							.first .second .third
						";
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
		SMALLER_THAN: "<";
		INTERROGATION_POINT: "?";
		STAR: "*";
		PLUS: "+";
		COMMERCIAL_A: "@";
		AND: "&&";
		OR: "||";
		XOR: "<>";
		QUOTE: "\"";
		BACKSLASH: "\\";
		SLASH: "/";
		EQUAL: "=";
		PERCENT: "%";

		NAME: "name";
		SELECT: "select";
		CONDITION: "condition";
		INSERT: "insert";
		UPDATE: "update";
		DELETE: "delete";
		DROP: "drop";
		INT: "int";
		UINT: "uint";
		FLOAT: "float";
		BOOL: "bool";
		STRING: "string";
		MAP: "map";
		LIST: "list";
		SET: "set";
		SORTED_SET: "sortedSet";
		SELECTOR: "selector";
		ROUTE: "route";
		BRIDGES: "bridges";
		TRUE: "true";
		FALSE: "false";
	`)
}
