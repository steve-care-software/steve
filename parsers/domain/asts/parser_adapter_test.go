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

		additionInParenthesis: .OPEN_PARENTHESIS .addition .CLOSE_PARENTHESIS;
		addition: (my_syscall .firstNumber:first .secondNumber:second) .firstNumber .PLUS_SIGN .secondNumber;
		secondNumber: (my_syscall) .N_THREE .N_FOUR .N_FIVE;
		firstNumber: .N_ONE .N_TWO;
		myReplacement: .N_ONE .N_THREE;
		replacedNumber: .N_TWO .N_FOUR;

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
		salut ( 12 + 345 )`), astRemaining...)

	grammarParserAdapter := grammars.NewParserAdapter()
	retGrammar, _, err := grammarParserAdapter.ToGrammar(grammarInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	parserAdapter := NewParserAdapter()
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

	grammarParserAdapter := grammars.NewParserAdapter()
	retGrammar, _, err := grammarParserAdapter.ToGrammar(grammarInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	parserAdapter := NewParserAdapter()
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

	grammarParserAdapter := grammars.NewParserAdapter()
	retGrammar, _, err := grammarParserAdapter.ToGrammar(grammarInput)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	parserAdapter := NewParserAdapter()
	_, _, err = parserAdapter.ToAST(retGrammar, astInput)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

}
