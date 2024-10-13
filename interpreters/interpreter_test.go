package interpreters

import (
	"bytes"
	"reflect"
	"testing"
)

func TestInterpreter_Success(t *testing.T) {

	testSuitesList := []testSuite{
		createAssign_uint8_inline_withRemaining(),
		createAssign_uint8_stack_withRemaining(),
		createAssign_uint16_inline_withRemaining(),
		createAssign_uint16_stack_withRemaining(),
		createAssign_uint32_inline_withRemaining(),
		createAssign_uint32_stack_withRemaining(),
		createAssign_uint64_inline_withRemaining(),
		createAssign_uint64_stack_withRemaining(),
		createAssign_int8_inline_withRemaining(),
		createAssign_int8_stack_withRemaining(),
		createAssign_int16_inline_withRemaining(),
		createAssign_int16_stack_withRemaining(),
		createAssign_int32_inline_withRemaining(),
		createAssign_int32_stack_withRemaining(),
	}

	for idx, oneTestSuite := range testSuitesList {
		interpreter := NewInterpreter(
			oneTestSuite.byteCode,
			oneTestSuite.params,
		)

		retStack, retBytes, err := interpreter.Execute()
		if err != nil {
			t.Errorf("suite (index: %d): the error was expected to be nil, error returned: %s", idx, err)
			return
		}

		if !bytes.Equal(retBytes, oneTestSuite.remaining) {
			t.Errorf("suite (index: %d): the returned bytes were invalid, expected: %v, returned; %v", idx, oneTestSuite.remaining, retBytes)
			return
		}

		if !reflect.DeepEqual(retStack, oneTestSuite.expectedStack) {
			t.Errorf("suite (index: %d): the returned stack is invalid, expected: %v, returned: %v", idx, oneTestSuite.expectedStack, retStack)
			return
		}
	}
}
