package scripts

// PROPERTY_HEAD

func grammarInput() []byte {
	return []byte(`
		v1;
		> .head;
		# .SPACE .TAB .EOL;

		head: .PROPERTY_HEAD .COLON .headOptions .SEMI_COLON
			---
				all: "
						head:
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

		headOptions: .propertyName .headOptionalOption*
				  ---
				  		propertyName: "
							name: myName;
						";

						access: "
							name: myName;
							access: 
								read: .first;
								write: .first .second;
								review: .first;
							;
						";

						compensation: "
							name: myName;
							compensation: 0.1, 0.23, 0.45;
						";

						all: "
							name: myName;
							compensation: 0.1, 0.23, 0.45;
							access: 
								read: .first;
								write: .first .second;
								review: .first;
							;
						";

						nameAtTheEnd: !"
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

		compensation: .PROPERTY_COMPENSATION .COLON .threeFloatNumbersBetweenZeroAndOne .SEMI_COLON
					---
						valid: "compensation: 0.1, 0.23, 0.45;";
					;

		access: .PROPERTY_ACCESS .COLON .permissionOptions .SEMI_COLON
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

		propertyReview: .PROPERTY_REVIEW .COLON .references .SEMI_COLON
					---
						oneRole: "review: .myRole;";
						multipleRoles: "review: .first .second .third;";
					;

		propertyWrite: .PROPERTY_WRITE .COLON .references  .SEMI_COLON
					---
						oneRole: "write: .myRole;";
						multipleRoles: "write: .first .second .third;";
					;

		propertyRead: .PROPERTY_READ .COLON .references .SEMI_COLON
					---
						oneRole: "read: .myRole;";
						multipleRoles: "read: .first .second .third;";
					;

		propertyName: .PROPERTY_NAME .COLON .variableName .SEMI_COLON
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

		PROPERTY_HEAD: "head";
		PROPERTY_NAME: "name";
		PROPERTY_READ: "read";
		PROPERTY_WRITE: "write";
		PROPERTY_REVIEW: "review";
		PROPERTY_ACCESS: "access";
		PROPERTY_COMPENSATION: "compensation";

		
	`)
}
