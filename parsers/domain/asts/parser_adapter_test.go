package asts

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/steve-care-software/steve/parsers/domain/grammars"
)

func TestParserAdapter_Success(t *testing.T) {
	grammarInput := []byte(`
		v1;
		>.line;
		# .SPACE .TAB. EOL;
		
		line: !.additionInParenthesis .additionInParenthesis
			| .N_ZERO
			;

		additionInParenthesis: .OPEN_PARENTHESIS ._myConstant[2,3] .CLOSE_PARENTHESIS;

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

func TestParserAdapter_isInfiniteRecursive_returnsError(t *testing.T) {
	grammarInput := []byte(`
		v1;
		>.line;
		# .SPACE .TAB. EOL;

		line: .line
			| .N_ZERO
			;

		N_ZERO: "0";
		SPACE: " ";
		TAB: "	";
		EOL: "
";
	`)

	astRemaining := []byte("this is a remaining")
	astInput := append([]byte(`
		salut ( 12 + 345 )`), astRemaining...)

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
	_, _, err = parserAdapter.ToAST(retGrammar, astInput)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

}
