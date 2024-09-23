package transpiles

func grammarInput() []byte {
	return []byte(`
		v1;
		> .all;
		# .SPACE .TAB .EOL;

		all: .header .block[1,]
			---
				valid: "
						v1;
						origin: b1680239c9fb16ab9f24bf6cd0091aaabe83c6cc2e5e0df24243435e9875b08dee46c12a2fd9ec9030179cd8c952b27ef32f54678386892add0bc3ab9b08f053;
						target: 2f078ff4f34a3c1d59a3dd1eac85e629c92800814bb442e81aa366849e59f0088b880af3ab2e1ff1dfc2e5c15217044bec27b61a15d259a7fcb413e5067aa5e8;
						>.myBlock;

						myBlock: .myToken[2] .MY_RULE[3]
						   	   | (.myOrigin[2] -> .myTarget) .secondToken[1] (.MY_SECOND_ORIGIN[2] -> .MY_TARGET)
						   	   ---
									firstTest:\"this is some value\";
									secondTest:!\"this is some value\";
						   	   ;

						secondBlock: .myToken[2] .MY_RULE[3]
						   		   | (.myOrigin[2] -> .myTarget) .secondToken[1] (.MY_SECOND_ORIGIN[2] -> .MY_TARGET)
						   		   ;
				";
			;

		header: .version .origin .target .root
			  ---
					valid: "
						v1;
						origin: b1680239c9fb16ab9f24bf6cd0091aaabe83c6cc2e5e0df24243435e9875b08dee46c12a2fd9ec9030179cd8c952b27ef32f54678386892add0bc3ab9b08f053;
						target: 2f078ff4f34a3c1d59a3dd1eac85e629c92800814bb442e81aa366849e59f0088b880af3ab2e1ff1dfc2e5c15217044bec27b61a15d259a7fcb413e5067aa5e8;
						>.myBlock;
					";
			  ;

		root: .BIGGER_THAN .DOT .variableName .SEMI_COLON
			---
				valid: ">.myRoot;";
			;

		target: .W_TARGET .COLON .hex .SEMI_COLON
			  ---
				valid: "target: b1680239c9fb16ab9f24bf6cd0091aaabe83c6cc2e5e0df24243435e9875b08dee46c12a2fd9ec9030179cd8c952b27ef32f54678386892add0bc3ab9b08f053;";
			  ;
			 

		origin: .W_ORIGIN .COLON .hex .SEMI_COLON
			  ---
				valid: "origin: b1680239c9fb16ab9f24bf6cd0091aaabe83c6cc2e5e0df24243435e9875b08dee46c12a2fd9ec9030179cd8c952b27ef32f54678386892add0bc3ab9b08f053;";
			  ;

		version: .LL_V .numbers .SEMI_COLON
				---
					oneNumber: "v1;";
					multipleNumbers: "v123;";
				;

		block: .variableName .COLON .lines .testSuites[0,] .SEMI_COLON
			---
				withSuites: "
					myBlock: .myToken[2] .MY_RULE[3]
						   | (.myOrigin[2] -> .myTarget) .secondToken[1] (.MY_SECOND_ORIGIN[2] -> .MY_TARGET)
						   ---
								firstTest:\"this is some value\";
								secondTest:!\"this is some value\";
						   ;
				";

				withoutSuites: "
					myBlock: .myToken[2] .MY_RULE[3]
						   | (.myOrigin[2] -> .myTarget) .secondToken[1] (.MY_SECOND_ORIGIN[2] -> .MY_TARGET)
						   ;
				";
			;

		testSuites: .MINUS[3] .testSuite+
				  ---
				  		valid: "
							---
							firstTest:\"this is some value\";
							secondTest:!\"this is some value\";
						";
				  ;

		testSuite: .variableName .COLON .EXCLAMATION_POINT[0,1] .stringValue .SEMI_COLON
				---
					validSuite: "firstTest:\"this is some value\";";
					failSuite: "firstTest:!\"this is some value\";";
				;

		stringValue: .QUOTE ![.BACKSLASH].QUOTE .QUOTE
					---
						stringInQuotes: "\"this is 13 \\\" string values!\"";
					;

		lines: .token[1,] .pipeLine[0,]
			---
				oneLine: "(.myOrigin[2] -> .myTarget) .myToken[2]";
				twoLines: "
					.myToken[2]
					| (.myOrigin[2] -> .myTarget) .myToken[2]
				";

				multipleines: "
					.myToken[2]
					| (.myOrigin[2] -> .myTarget)
					| (.myOrigin[2] -> .myTarget) .myToken[2]
				";
			;
		
		pipeLine: .PIPE .token[1,]
				---
					valid: "| (.myOrigin[2] -> .myTarget) .myToken[2]";
				;

		token: .update
			 | .pointer
			 ---
			 	update: "(.myOrigin[2] -> .myTarget)";
				insert: ".myToken[2]";
			 ;

		update: .OPEN_PARENTHESIS .pointer .arrow .elementReference .CLOSE_PARENTHESIS
			  ---
			  		valid: "(.myOrigin[2] -> .myTarget)";
					validAgain: "(.MY_RULE[2] -> .TARGET_RULE)";
					validMixed: "(.MY_RULE[2] -> .myTarget)";
					validMixedReversed: "(.myOrigin[2,] -> .TARGET_RULE)";
			  ;

		arrow: .MINUS .BIGGER_THAN
			 ---
			 	arrow: "->";
			 ;

		pointer: .elementReference .cardinality
			   ---
					token: ".myToken[2]";
					rule: ".MY_RULE[2,]";
			   ;

		elementReference: .DOT .element
						---
							token: ".myToken";
							rule: ".MY_RULE";
						;

		element: .variableName
				 | .ruleName
				 ---
					ruleValid: "MY_RULE";
					ruleInvalidUnderscoreFirst: !"_RULE";
					ruleInvalidNoSecondLetter: !"R";
					tokenValid: "myVariable";
					tokenFirstUpperCaseLetter: !"MyVariable";
				;

		cardinality: .OPEN_BRACKET .cardinalityValue .CLOSE_BRACKET
				   ---
						number: "[23]";
						numberWithComma: "[23,]";
						twoNumbers: "[0, 34]";
				   ;

		cardinalityValue: .numbers .COMMA .numbers[0,]
				   		| .numbers
						---
							number: "23";
							numberWithComma: "23,";
							twoNumbers: "0";
						;

		ruleName: .oneUpperCaseLetter .ruleCharAfterFirstChar+
				---
					valid: "MY_RULE";
					invalidUnderscoreFirst: !"_RULE";
					invalidNoSecondLetter: !"R";
				;

		ruleCharAfterFirstChar: .oneUpperCaseLetter
							  | .UNDERSCORE
							  ---
							  	letterA: "A";
								letterB: "B";
								letterC: "C";
								letterD: "D";
								letterE: "E";
								letterF: "F";
								letterG: "G";
								letterH: "H";
								letterI: "I";
								letterJ: "J";
								letterK: "K";
								letterL: "L";
								letterM: "M";
								letterN: "N";
								letterO: "O";
								letterP: "P";
								letterQ: "Q";
								letterR: "R";
								letterS: "S";
								letterT: "T";
								letterU: "U";
								letterV: "V";
								letterW: "W";
								letterX: "X";
								letterY: "Y";
								letterZ: "Z";
								underscore: "_";
							  ;

		variableName: .oneLowerCaseLetter .letters+
					---
						good: "myVariable";
						firstUpperCaseLetter: !"MyVariable";
					;

		hex: .oneHexCharacter[128]
			---
				valid: "b1680239c9fb16ab9f24bf6cd0091aaabe83c6cc2e5e0df24243435e9875b08dee46c12a2fd9ec9030179cd8c952b27ef32f54678386892add0bc3ab9b08f053";
				tooManyCharacters: !"wb1680239c9fb16ab9f24bf6cd0091aaabe83c6cc2e5e0df24243435e9875b08dee46c12a2fd9ec9030179cd8c952b27ef32f54678386892add0bc3ab9b08f053";
			;


		oneHexCharacter: .letterLowercaseAtoF
					   | .oneNumber
					   ---
							letterA: "a";
							letterB: "b";
							letterC: "c";
							letterD: "d";
							letterE: "e";
							letterF: "f";
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

		letters: .uppercaseLetters
				 | .lowerCaseLetters
				---
					oneLowerCaseLetter: "a";
					lowerCaseLetters: "abcdefghijklmnopqrstuvwxyz";
					oneUpperCaseLetter: "A";
					upperCaseLetter: "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
					oneNumber: !"0";
				;

		uppercaseLetters: .oneUpperCaseLetter+;
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
							;

		lowerCaseLetters: .oneLowerCaseLetter+
							---
						  		oneLowerCaseLetter: "a";
						  		lowerCaseLetters: "abcdefghijklmnopqrstuvwxyz";
								oneUpperCaseLetter: !"A";
								upperCaseLetter: !"ABCDEFGHIJKLMNOPQRSTUVWXYZ";
								oneNumber: !"0";
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
							;

		letterLowercaseAtoF: .LL_A
						   | .LL_B
						   | .LL_C
						   | .LL_D
						   | .LL_E
						   | .LL_F
						   ---
								letterA: "a";
								letterB: "b";
								letterC: "c";
								letterD: "d";
								letterE: "e";
								letterF: "f";
						   ;

		numbers: .oneNumber+
				---
					oneNumber: "1";
					numberWithAllNumbers: "1234567890";
					negativeNumberWithAllNumbers: !"-1234567890";
					oneLettter: !"a";
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

		W_ORIGIN: "origin";
		W_TARGET: "target";

		SEMI_COLON: ";";
		COLON: ":";
		BIGGER_THAN: ">";
		DOT: ".";
		UNDERSCORE: "_";
		OPEN_BRACKET: "[";
		CLOSE_BRACKET: "]";
		COMMA: ",";
		MINUS: "-";
		OPEN_PARENTHESIS: "(";
		CLOSE_PARENTHESIS: ")";
		QUOTE: "\"";
		BACKSLASH: "\\";
		EXCLAMATION_POINT: "!";
		PIPE: "|";

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

		SPACE: " ";
		TAB: "	";
		EOL: "
";
	`)
}
