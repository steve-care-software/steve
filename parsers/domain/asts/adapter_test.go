package asts

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/steve-care-software/steve/parsers/domain/grammars"
)

func TestParserAdapter_withBalance_Success(t *testing.T) {
	grammarInput := []byte(`
		v1;
		> .assignment;
		# .SPACE .TAB. EOL;
		
		assignment: .type .VARIABLE .EQUAL .value
					[
						.FLOAT[0]:
							.value[0][0]->floatValue[0]:
								!.value[0][0]->uintValue[0]
						;

						.type[0][0]->UINT[0]:
							.uintValue[0]:
								!.value[0][0]->floatValue[0]
						;
					];
				  ;

		type: .FLOAT
			| .UINT
			;
			
		value: .floatValue
			 | .uintValue
			 ;

		uintValue: .N_ZERO;
		floatValue: .N_ZERO .DOT .N_ONE;

		FLOAT: "float";
		EQUAL: "=";
		UINT: "uint";
		VARIABLE: "myVariable";
		N_ZERO: "0";
		N_ONE: "1";
		DOT: ".";
		SPACE: " ";
		TAB: "	";
		EOL: "
";
	`)

	grammarParserAdapter := grammars.NewAdapter()
	retGrammar, _, err := grammarParserAdapter.ToGrammar(grammarInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	parserAdapter := NewAdapter()
	astRemaining := []byte("|this is a remaining")

	validInputs := [][]byte{
		append([]byte(`float myVariable = 0.1`), astRemaining...),
		append([]byte(`uint myVariable = 0`), astRemaining...),
	}

	for _, oneValidInput := range validInputs {
		retAST, retRemaining, err := parserAdapter.ToAST(retGrammar, oneValidInput)
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		if !bytes.Equal(astRemaining, retRemaining) {
			t.Errorf("the returned remaining is invalid")
			return
		}

		if retAST == nil {
			t.Errorf("the returned AST was expected to NOT be nil")
			return
		}
	}

	invalidInputs := [][]byte{
		append([]byte(`uint myVariable = 0.1`), astRemaining...),
		append([]byte(`float myVariable = 0`), astRemaining...),
	}

	for _, oneInvalid := range invalidInputs {
		_, _, err := parserAdapter.ToAST(retGrammar, oneInvalid)
		if err == nil {
			t.Errorf("the error was expected to be valid, nil returned")
			return
		}
	}

}

func TestParserAdapter_Success(t *testing.T) {
	grammarInput := []byte(`
		v1;
		>.line;
		# .SPACE .TAB. EOL;

		line: !.additionInParenthesis .additionInParenthesis
			| .N_ZERO
			;

		additionInParenthesis: .OPEN_PARENTHESIS ._myConstant[2,3] .CLOSE_PARENTHESIS
							;



		_myConstant: .N_ZERO ._mySubConstant[2] .N_FOUR[4];
		_mySubConstant: .N_ONE .N_TWO[2] .N_THREE;

		N_ZERO: "0";
		N_ONE: "1";
		N_TWO: "2";
		N_THREE: "3";
		N_FOUR: "4";
		N_FIVE: "5";
		N_SIX: "6";
		OPEN_PARENTHESIS: "(";
		CLOSE_PARENTHESIS: ")";
		PLUS_SIGN: "+";
		SPACE: " ";
		TAB: "	";
		EOL: "
";
	`)

	astRemaining := []byte("this is a remaining")
	astInput := append([]byte(`
		salut ( 012231223444401223122344440122312234444 )`), astRemaining...)

	grammarParserAdapter := grammars.NewAdapter()
	retGrammar, _, err := grammarParserAdapter.ToGrammar(grammarInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	parserAdapter := NewAdapter()
	retAST, retRemaining, err := parserAdapter.ToAST(retGrammar, astInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(astRemaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	fmt.Printf("\n%v\n", retAST)

}

func TestParserAdapter_withFiniteRecursivity_Success(t *testing.T) {
	grammarInput := []byte(`
		v1;
		>.addition;
		# .SPACE .TAB. EOL;

		addition: .OPEN_PARENTHESIS .addition .CLOSE_PARENTHESIS
				| .firstNumber .PLUS_SIGN .secondNumber
				;

		secondNumber: .N_THREE .N_FOUR .N_FIVE;
		firstNumber: .N_ONE .N_TWO;

		N_ZERO: "0";
		N_ONE: "1";
		N_TWO: "2";
		N_THREE: "3";
		N_FOUR: "4";
		N_FIVE: "5";
		N_SIX: "6";
		OPEN_PARENTHESIS: "(";
		CLOSE_PARENTHESIS: ")";
		PLUS_SIGN: "+";
		SPACE: " ";
		TAB: "	";
		EOL: "
";
	`)

	astRemaining := []byte("this is a remaining")
	astInput := append([]byte(`
		( 12 + 345 )`), astRemaining...)

	grammarParserAdapter := grammars.NewAdapter()
	retGrammar, _, err := grammarParserAdapter.ToGrammar(grammarInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	parserAdapter := NewAdapter()
	retAST, retRemaining, err := parserAdapter.ToAST(retGrammar, astInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(astRemaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	fmt.Printf("\n%v\n", retAST)

}
