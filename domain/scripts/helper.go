package scripts

func grammarInput() []byte {
	return []byte(`
		v1;
		> .head;
		# .SPACE .TAB .EOL;

		script: .schema
			---
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
							myList: list;
							mySet: set;
							mySortedSet: sorted_set;
							myVector: vector[float];

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
					myList: list;
					mySet: set;
					mySortedSet: sorted_set;
					myVector: vector[float];

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

		pointSuite: .variableName .BRACKET_OPEN .references .BRACKET_CLOSE .COLON .pointSuiteOptionLine+ .SEMI_COLON
				---
					valid: "
						mySuite[.son .grandGrandFather]:
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
						| .referencesInParenthesis
						---
							expectedInvalid: "!(.son .father .grandFather .grandGrandFather)";
							expectedValid: "(.son .father .grandFather .grandGrandFather)";
						;

		pointInvalidLink: .EXCLAMATION_POINT .referencesInParenthesis
						---
							valid: "!(.son .father .grandFather .grandGrandFather)";
						;

		referencesInParenthesis: .PARENTHESIS_OPEN .references .PARENTHESIS_CLOSE
							---
								valid: "(.son .father .grandFather .grandGrandFather)";
							;

		suitePrefix: .HYPHEN[3]
				---
					valid: "---";
				;

		links: .references .pipeReferences*
			---
				oneLine: ".son .father";
				multipleLine: "
					.son .father
					| .father .grandFather
					| .grandFather .grandGrandFather
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

		pipeReferences: .PIPE .references
					---
						valid: "| .son .father";
					;

		pointWithTypes: .pointWithType+
					---
						valid: "
							son;
							father;
							grandFather;
							grandGrandFather;
							myList: list;
							mySet: set;
							mySortedSet: sorted_set;
							myVector: vector[float];
						";
					;

		pointWithType: .variableName .colonPointType? .SEMI_COLON
					---
						point: "myPoint;";
						pointWithType: "mySecondPoint: vector[uint];";
					;

		colonPointType: .COLON .pointType
					---
						list: ": list";
						set: ": set";
						sortedSet: ": sorted_set";
				  		vectorFloat: ": vector[float]";
						vectorInt: ": vector[int]";
						vectorUint: ": vector[uint]";
					;
					
		pointType: .pointVector
				  | .LIST
				  | .SET
				  | .SORTED_SET 
				  ---
				  		list: "list";
						set: "set";
						sortedSet: "sorted_set";
				  		vectorFloat: "vector[float]";
						vectorInt: "vector[int]";
						vectorUint: "vector[uint]";
				  ;

		pointVector: .VECTOR .BRACKET_OPEN .vectorTypes .BRACKET_CLOSE
					---
						float: "vector[float]";
						int: "vector[int]";
						uint: "vector[uint]";
					;

		vectorTypes: .FLOAT
				   | .INT
				   | .UINT
				   ---
				   		float: "float";
						int: "int";
						uint: "uint";
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
		LIST: "list";
		SET: "set";
		SORTED_SET: "sorted_set";
		VECTOR: "vector";
		FLOAT: "float";
		INT: "int";
		UINT: "uint";

		BRACKET_OPEN: "[";
		BRACKET_CLOSE: "]";
		PARENTHESIS_OPEN: "(";
		PARENTHESIS_CLOSE: ")";
		PIPE: "|";
		HYPHEN: "-";
		EXCLAMATION_POINT: "!";

		
	`)
}
