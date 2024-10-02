package grammars

import (
	"bytes"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`
		v1;
		>.myRoot;
		#.first.second.third;

		myFirst: !.myFirst[1] .mySecond* .myThird+ .myFourth? .myFifth[1,]
					[
						.myFirst[0][1]->MY_RULE[0][1] :
							.myFirst[0][1]->MY_RULE[0][0]
						;

						.myFirst[0][1]->MY_RULE[0][1] :
							.myFirst[0][1]->mySecond[0][0]->myThird[0]
						;

						.myFirst[0][1]->MY_RULE[0][1] :
							.myFirst[0][1]->MY_RULE[0][0]:
								.myFirst[0][1]->mySecond[0][0]->myThird[0]
						;

						.myFirst[0][1]->MY_RULE[0][1];
					];
				 | .myFirst[1] .mySecond* .myThird+ .myFourth .myFifth[1,]
				 ---
				 	firstTest:!"this is some value";
					secondTest:!"this is some value";
				 ;

		mySecond: .myFirst[1] ._myConstant* .myThird+ .myFourth[2] .myFifth[1,]
					[
						.type[0][0]->FLOAT[0]:
							.value[0][0]->floatValue[0]
						;

						.type[0][0]->UINT[0]:
							.value[0][0]->uintValue[0]
						;
					];
				 ;

		_myConstant: .MY_RULE .MY_SECOND_RULE[2] ._mySubConstant ._otherConstant[2];

		FIRST: "this \" with escape";
		SECOND: "some value";
		`), remaining...)

	retAdapter := NewAdapter()
	retGrammar, retRemaining, err := retAdapter.ToGrammar(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retGrammar.Version() != 1 {
		t.Errorf("the version was expected to be %d, %d returned", 1, retGrammar.Version())
		return
	}

	if retGrammar.Root().Name() != "myRoot" {
		t.Errorf("the root was expected to be %s, %s returned", "myRoot", retGrammar.Root())
		return
	}

	retBlocks := retGrammar.Blocks().List()
	if len(retBlocks) != 2 {
		t.Errorf("the grammar was expected to contain %d block instances, %d returned", 2, len(retBlocks))
		return
	}

	retRules := retGrammar.Rules().List()
	if len(retRules) != 2 {
		t.Errorf("the grammar was expected to contain %d rule instances, %d returned", 2, len(retRules))
		return
	}

	if !retGrammar.HasOmissions() {
		t.Errorf("the grammar was expected to contain omissions")
		return
	}

	retOmissions := retGrammar.Omissions().List()
	if len(retOmissions) != 3 {
		t.Errorf("the grammar was expected to contain %d omission elements, %d returned", 3, len(retOmissions))
		return
	}

	if !retGrammar.HasConstants() {
		t.Errorf("the grammar was expected to contain constants")
		return
	}

	retConstants := retGrammar.Constants().List()
	if len(retConstants) != 1 {
		t.Errorf("the grammar was expected to contain %d constant elements, %d returned", 1, len(retConstants))
		return
	}
}

func TestAdapter_blocks_withoutBlocks_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(``), remaining...)

	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToBlocks(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestAdapter_suites_withoutSuites_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(``), remaining...)

	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToSuites(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestAdapter_suite_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`myTest:"somedata";`), remaining...)

	retAdapter := NewAdapter().(*adapter)
	retSuite, retRemaining, err := retAdapter.bytesToSuite(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retSuite.Name() != "myTest" {
		t.Errorf("the suite name was expected to be (%s), (%s) returned", "myTest", retSuite.Name())
		return
	}

	if retSuite.IsFail() {
		t.Errorf("the suite was expected to NOT fail")
		return
	}
}

func TestAdapter_suite_isFail_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`myTest:!"somedata";`), remaining...)

	retAdapter := NewAdapter().(*adapter)
	retSuite, retRemaining, err := retAdapter.bytesToSuite(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	if retSuite.Name() != "myTest" {
		t.Errorf("the suite name was expected to be (%s), (%s) returned", "myTest", retSuite.Name())
		return
	}

	if !retSuite.IsFail() {
		t.Errorf("the suite was expected to fail")
		return
	}
}

func TestAdapter_suite_withInvalidElement_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`myTest:myElement`), remaining...)

	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToSuite(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestAdapter_suite_withInvalidBlockNameDefinition_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`#myTest:.myElement`), remaining...)

	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToSuite(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestAdapter_suite_withoutSuiteLineSuffix_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`myTest:.myElement`), remaining...)

	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToSuite(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestAdapter_suite_withoutSuiteLineSuffix_withoutRemainingBytes_returnsError(t *testing.T) {
	input := []byte(`myTest:.myElement`)
	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToSuite(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
	}
}

func TestAdapter_lines_withoutLine_returnsError(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(`not a line`), remaining...)

	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToLines(input)
	if err == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
		return
	}
}

func TestAdapter_tokens_Success(t *testing.T) {
	remaining := []byte("this is some remaining")
	input := append([]byte(`.myFirst[1] .mySecond* .myThird+ .myFourth .myFifth[1,]`), remaining...)

	retAdapter := NewAdapter().(*adapter)
	retToken, retRemaining, err := retAdapter.bytesToTokens(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	list := retToken.List()
	if len(list) != 5 {
		t.Errorf("the tokens list was expected to contain %d tokens, %d returned", 5, len(list))
		return
	}
}

func TestAdapter_token_withBlockName_withCardinality_Success(t *testing.T) {
	remaining := []byte("this is some remaining")
	input := append([]byte(`.myToken[1]`), remaining...)

	retAdapter := NewAdapter().(*adapter)
	retToken, retRemaining, err := retAdapter.bytesToToken(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retToken.Name() != "myToken" {
		t.Errorf("the token name is invalid, expected (%s), returned (%s)", "myToken", retToken.Name())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	cardinality := retToken.Cardinality()
	if cardinality.Min() != 1 {
		t.Errorf("the cardinality min was expected to be (%d), (%d) returned", 1, cardinality.Min())
		return
	}

	if !cardinality.HasMax() {
		t.Errorf("the cardinality was expected to contain a max")
		return
	}

	pMax := cardinality.Max()
	max := *pMax
	if max != 1 {
		t.Errorf("the cardinality max was expected to be (%d), (%d) returned", 1, max)
		return
	}
}

func TestAdapter_token_withBlockName_withoutCardinality_Success(t *testing.T) {
	remaining := []byte("%!this is some remaining")
	input := append([]byte(` . myToken`), remaining...)

	retAdapter := NewAdapter().(*adapter)
	retToken, retRemaining, err := retAdapter.bytesToToken(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retToken.Name() != "myToken" {
		t.Errorf("the token name is invalid, expected (%s), returned (%s)", "myToken", retToken.Name())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	cardinality := retToken.Cardinality()
	if cardinality.Min() != 1 {
		t.Errorf("the cardinality min was expected to be (%d), (%d) returned", 1, cardinality.Min())
		return
	}

	if !cardinality.HasMax() {
		t.Errorf("the cardinality was expected to contain a max")
		return
	}

	pMax := cardinality.Max()
	max := *pMax
	if max != 1 {
		t.Errorf("the cardinality max was expected to be (%d), (%d) returned", 1, max)
		return
	}
}

func TestAdapter_token_withRuleName_withCardinality_Success(t *testing.T) {
	remaining := []byte("this is some remaining")
	input := append([]byte(`. MY_RULE [1]`), remaining...)

	retAdapter := NewAdapter().(*adapter)
	retToken, retRemaining, err := retAdapter.bytesToToken(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if retToken.Name() != "MY_RULE" {
		t.Errorf("the token name is invalid, expected (%s), returned (%s)", "MY_RULE", retToken.Name())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the remaining bytes are invalid, expected (%s), returned (%s)", string(remaining), string(retRemaining))
		return
	}

	cardinality := retToken.Cardinality()
	if cardinality.Min() != 1 {
		t.Errorf("the cardinality min was expected to be (%d), (%d) returned", 1, cardinality.Min())
		return
	}

	if !cardinality.HasMax() {
		t.Errorf("the cardinality was expected to contain a max")
		return
	}

	pMax := cardinality.Max()
	max := *pMax
	if max != 1 {
		t.Errorf("the cardinality max was expected to be (%d), (%d) returned", 1, max)
		return
	}
}

func TestAdapter_token_withoutBlockName_withoutRuleName_returnsError(t *testing.T) {
	remaining := []byte("this is some remaining")
	input := append([]byte(`.+++`), remaining...)

	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToToken(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}

}

func TestAdapter_token_withoutTokenReferenceByte_returnsError(t *testing.T) {
	remaining := []byte("this is some remaining")
	input := append([]byte(`myToken [ 1 ]`), remaining...)

	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToToken(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestAdapter_token_withoutInput_returnsError(t *testing.T) {
	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToToken([]byte{})
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestAdapter_rule_Success(t *testing.T) {
	expectedName := "MY_RULE"
	expectedValue := []byte(`this " with escape`)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`MY_RULE: "this \" with escape";this is some remaining`)

	retAdapter := NewAdapter().(*adapter)
	retRule, retRemaining, err := retAdapter.bytesToRule(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retRule.Name() != expectedName {
		t.Errorf("the name was expected to be %s, %s returned", expectedName, retRule.Name())
		return
	}

	if !bytes.Equal(expectedValue, retRule.Bytes()) {
		t.Errorf("the expected value was (%s), returned (%s)", expectedValue, retRule.Bytes())
		return
	}
}

func TestAdapter_rule_withInvalidName_returnsError(t *testing.T) {
	input := []byte(`_MY_RULE: "this \" with escape";this is some remaining`)
	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToRule(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestAdapter_rule_withoutValue_returnsError(t *testing.T) {
	input := []byte(`MY_RULE: "";this is some remaining`)
	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToRule(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestAdapter_cardinality_withoutMax_Success(t *testing.T) {
	expectedMin := uint(1)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`[1, ]this is some remaining`)

	retAdapter := NewAdapter().(*adapter)
	retCardinality, retRemaining, err := retAdapter.bytesToCardinality(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retCardinality.Min() != expectedMin {
		t.Errorf("the min was expected to be %d, %d returned", expectedMin, retCardinality.Min())
		return
	}

	if retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to NOT contain a max")
		return
	}
}

func TestAdapter_cardinality_withMax_Success(t *testing.T) {
	expectedMin := uint(1)
	expectedMax := uint(1)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`[ 1, 1 ] this is some remaining`)

	retAdapter := NewAdapter().(*adapter)
	retCardinality, retRemaining, err := retAdapter.bytesToCardinality(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retCardinality.Min() != expectedMin {
		t.Errorf("the min was expected to be %d, %d returned", expectedMin, retCardinality.Min())
		return
	}

	if !retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to contain a max")
		return
	}

	pRetMax := retCardinality.Max()
	if *pRetMax != expectedMax {
		t.Errorf("the max was expected to be %d, %d returned", expectedMax, *pRetMax)
		return
	}
}

func TestAdapter_cardinality_withZeroPlus_Success(t *testing.T) {
	expectedMin := uint(0)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`*this is some remaining`)

	retAdapter := NewAdapter().(*adapter)
	retCardinality, retRemaining, err := retAdapter.bytesToCardinality(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retCardinality.Min() != expectedMin {
		t.Errorf("the min was expected to be %d, %d returned", expectedMin, retCardinality.Min())
		return
	}

	if retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to NOT contain a max")
		return
	}
}

func TestAdapter_cardinality_withOnePlus_Success(t *testing.T) {
	expectedMin := uint(1)
	expectedRemaining := []byte("this is some remaining")
	input := []byte(`+this is some remaining`)

	retAdapter := NewAdapter().(*adapter)
	retCardinality, retRemaining, err := retAdapter.bytesToCardinality(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(expectedRemaining, retRemaining) {
		t.Errorf("the expected renaining was (%s), returned (%s)", expectedRemaining, retRemaining)
		return
	}

	if retCardinality.Min() != expectedMin {
		t.Errorf("the min was expected to be %d, %d returned", expectedMin, retCardinality.Min())
		return
	}

	if retCardinality.HasMax() {
		t.Errorf("the cardinality was expected to NOT contain a max")
		return
	}
}

func TestAdapter_cardinality_withInvalidInput_returnsError(t *testing.T) {
	input := []byte(`this is some invalid input`)
	retAdapter := NewAdapter().(*adapter)
	_, _, err := retAdapter.bytesToCardinality(input)
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
