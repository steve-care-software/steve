package applications

import (
	"testing"

	"github.com/steve-care-software/steve/parsers/domain/grammars"
	"github.com/steve-care-software/steve/parsers/domain/walkers/elements"
)

type testValue struct {
	First  string
	Second string
}

func TestApplication_withSuites_execute_Success(t *testing.T) {

	scriptInput := []byte(`
		first:second
	`)

	sequence := elements.Element{
		ElementFn: func(input any) (any, error) {
			return input, nil
		},
		TokenList: &elements.TokenList{
			MapFn: func(elementName string, mp map[string][]any) (any, error) {
				return &testValue{
					First:  string(mp["variableName"][0].(string)),
					Second: string(mp["variableComplex"][0].(string)),
				}, nil
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
				"variableComplex": {
					SelectorScript: []byte(`
						v1;
						name: mySelector;
						variableComplex[0][0];
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

	grammarInput := []byte(`
		v1;
		> .pointer;
		# .SPACE .TAB .EOL;

		pointer: .variableName .COLON .variableComplex
				---
					valid: "first:second";
				;

		variableComplex: .letters+;

		variableName: .oneLowerCaseLetter .letters+
					---
						good: "myVariable";
						firstUpperCaseLetter: !"MyVariable";
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

		COLON: ":";

		SPACE: " ";
		TAB: "	";
		EOL: "
";
	`)

	grammarAdapter := grammars.NewAdapter()
	retGrammar, _, err := grammarAdapter.ToGrammar(grammarInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	application, err := NewBuilder().Create().WithElement(sequence).Now()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = application.Suites(retGrammar)
	if err != nil {
		t.Errorf("there was an error while running the grammar test suites: %s", err.Error())
		return
	}

	retOutput, _, err := application.Execute(scriptInput, retGrammar)
	if err != nil {
		t.Errorf("there was an error while running the grammar test suites: %s", err.Error())
		return
	}

	if casted, ok := retOutput.(*testValue); ok {
		if casted.First != "first" {
			t.Errorf("the first property was expected to be '%s', '%s' returned", casted.First, "first")
			return
		}

		if casted.Second != "second" {
			t.Errorf("the second property was expected to be '%s', '%s' returned", casted.Second, "second")
			return
		}

		return
	}

	t.Error("the returned parsed instance could not be casted properly")
}

func TestApplication_grammar_withSuites_Success(t *testing.T) {
	grammarInput := []byte(`
		v1;
		> .instructions;
		# .SPACE .TAB .EOL;

		instructions: .instruction*;
		instruction: .instructionPossibilities .SEMI_COLON;
		instructionPossibilities: .assignment;

		assignment: .boolAssignment
				  	| .uintAssignment
					| .intAssignment
					| .floatAssignment
					---
						bool: "bool myVariable = true";
						uint: "uint8 myVariable = 32";
						int: "int64 myVariable = -345";
						float: "float32 myVariable = -12.45";
					;

		stringValue: .QUOTE ![.BACKSLASH].QUOTE .QUOTE
					---
						stringInQuotes: "\"this is 13 \\\" string values!\"";
					;

		boolAssignment: .typeBoolDefinition .EQUAL .boolValue
						---
							boolAssignment: "bool myVariable = true";
						;

		typeBoolDefinition: .typeBool .variableName
							---
								definition: "bool myVariable";
							;
		boolValue: .TRUE
				   | .FALSE
				   ---
						true: "true";
						false: "false";
					;


		uintAssignment: .typeUintDefinition .EQUAL .numbers
						---
							uint8: "uint8 myVariable = 45";
							uint16: "uint16 myVariable = 45";
							uint32: "uint32 myVariable = 45";
							uint64: "uint64 myVariable = 45";
						;

		typeUintDefinition: .typeUintAll .variableName
							---
								uint8: "uint8 myVariable";
								uint16: "uint16 myVariable";
								uint32: "uint32 myVariable";
								uint64: "uint64 myVariable";
							;

		intAssignment: .typeIntDefinition .EQUAL .intNumbers;
		typeIntDefinition: .typeIntAll .variableName
							---
								int8: "int8 myVariable";
								int16: "int16 myVariable";
								int32: "int32 myVariable";
								int64: "int64 myVariable";
							;

		floatAssignment: .typeFloatDefinition .EQUAL .floatNumbers
						---
							float32: "float32 myVariable = 12.32";
							float64: "float32 myVariable = 12.32";
						;

		typeFloatDefinition: .typeFloatAll .variableName
							---
								float32: "float32 myVariable";
								float64: "float64 myVariable";
							;

		variableName: .oneLowerCaseLetter .letters+
					---
						good: "myVariable";
						firstUpperCaseLetter: !"MyVariable";
					;

		typePrimitive: .typeBool
					   | .typeString
					   | .typeFloatAll
					   | .typeUintAll
					   | .typeIntAll
					   ---
							bool: "bool";
							string: "string";
							float32: "float32";
							float64: "float64";
							int8: "int8";
							int16: "int16";
							int32: "int32";
							int64: "int64";
							uint8: "uint8";
							uint16: "uint16";
							uint32: "uint32";
							uint64: "uint64";
					   ;

		typeBool: .LL_B .LL_O[2] .LL_L;
		typeString: .LL_S .LL_T .LL_R .LL_I .LL_N .LL_G;

		typeFloatAll: .typeFloat64
				 	| .typeFloat32
					---
						float32: "float32";
						float64: "float64";
					;

		typeFloat64: .typeFloat .sixtyFour;
		typeFloat32: .typeFloat .thirtyTwo;
		typeFloat: .LL_F .LL_L .LL_O .LL_A .LL_T;

		typeUintAll: .typeUint64
				  	| .typeUint32
					| .typeUint16
					| .typeUint8
					---
						int8: !"int8";
						int16: !"int16";
						int32: !"int32";
						int64: !"int64";
						uint8: "uint8";
						uint16: "uint16";
						uint32: "uint32";
						uint64: "uint64";
				  	;

		typeUint64: .typeUint .sixtyFour;
		typeUint32: .typeUint .thirtyTwo;
		typeUint16: .typeUint .sixteen;
		typeUint8: .typeUint .N_HEIGHT;
		typeUint: .LL_U .typeInt;

		typeIntAll: .typeInt64
				  | .typeInt32
				  | .typeInt16
				  | .typeInt8
				  ---
					int8: "int8";
					int16: "int16";
					int32: "int32";
					int64: "int64";
					uint8: !"uint8";
					uint16: !"uint16";
					uint32: !"uint32";
					uint64: !"uint64";
				;

		typeInt64: .typeInt .sixtyFour;
		typeInt32: .typeInt .thirtyTwo;
		typeInt16: .typeInt .sixteen;
		typeInt8: .typeInt .N_HEIGHT;
		typeInt: .LL_I .LL_N .LL_T;

		sixtyFour: .N_SIX .N_FOUR;
		thirtyTwo: .N_THREE .N_TWO;
		sixteen: .N_ONE .N_SIX;

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

		floatNumbers: .negativeFloatNumber
					  | .floatNumber
					---
						floatValue: "13.09";
						negativeFloatValue: "-13.09";
						negativeNumberWithAllNumbers: !"-1234567890";
						oneLettter: !"a";
					;

		negativeFloatNumber: .MINUS .floatNumber;
		floatNumber: .numbers .DOT .numbers
					---
						floatValue: "13.09";
						negativeFloatValue: !"-13.09";
						negativeNumberWithAllNumbers: !"-1234567890";
						oneLettter: !"a";
					;

		intNumbers: .negativeNumber
					| .numbers
					---
						negativeNumberWithAllNumbers: "-1234567890";
						numberWithAllNumbers: "1234567890";
						oneLettter: !"a";
					;

		negativeNumber: .MINUS .numbers
				---
					oneNegativeZero: "-0";
					negativeNumberWithAllNumbers: "-1234567890";
					numberWithAllNumbers: !"1234567890";
					oneLettter: !"a";
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

		BACKSLASH: "\\";
		QUOTE: "\"";

		OPEN_PARENTHESIS: "(";
		CLOSE_PARENTHESIS: ")";

		PLUS: "+";
		MINUS: "-";
		DOT: ".";
		EQUAL: "=";
		SEMI_COLON: ";";

		TRUE: "true";
		FALSE: "false";

		SPACE: " ";
		TAB: "	";
		EOL: "
";
	`)

	grammarAdapter := grammars.NewAdapter()
	retGrammar, _, err := grammarAdapter.ToGrammar(grammarInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	application, _ := NewBuilder().Create().Now()
	err = application.Suites(retGrammar)
	if err != nil {
		t.Errorf("there was an error while running the grammar test suites: %s", err.Error())
		return
	}
}
