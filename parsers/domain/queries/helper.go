package queries

func fetchGrammarInput() []byte {
	return []byte(`
		v1;
		> .root;
		# .SPACE .TAB .EOL;

		root: .head .chainLine
			---
				valid: "
					v1;
					name: mySelector;
					myChain[0][0]->myChain[0][0]->myChain[0][0]->RULE;
				";
			;

		chainLine: .chain .SEMI_COLON
				---
					complex: "myChain[0][0]->myChain[0][0]->myChain[0][0]->RULE;";
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

		chain: .elementName .token?
			---
				index: "myChain";
				withToken: "myChain[0]";
				withTokenElement: "myChain[0][0]";
				withTokenElementChain: "myChain[0][0]->myToken";
				complex: "myChain[0][0]->myChain[0][0]->myChain[0][0]->RULE";
			;

		arrowChain: .ARROW .chain
				---
					valid: "->myChain[0][0]->myChain[0][0]->myChain[0][0]->RULE";
				;

		token: .index .element?
			---
				index: "[0]";
				withElement: "[0][1]->myToken";
			;

		element: .index .arrowChain?
				---
					index: "[0]";
					withChain: "[0]->myToken";
				;

		elementName: .variableName
					| .constantName
					| .ruleName
					---
						block: "myBlock";
						constant: "_myConstant";
						rule: "MY_RULE";
					;

		constantName: .UNDERSCORE .variableName
					---
						valid: "_myConstant";
					;

		index: .BRACKET_OPEN .numbers .BRACKET_CLOSE
					---
						zero: "[0]";
						anyNumber: "[23]";
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

		UNDERSCORE: "_";
		BRACKET_OPEN: "[";
		BRACKET_CLOSE: "]";
		ARROW: "->";
		SEMI_COLON: ";";
		COLON: ":";

		NAME: "name";

		SPACE: " ";
		TAB: "	";
		EOL: "
";
	`)
}
