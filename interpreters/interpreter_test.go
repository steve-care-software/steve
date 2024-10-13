package interpreters

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestAssign_uint8_inline_withRemaining_Success(t *testing.T) {

	// set the amount of instructions:
	amountInstructions := uint64(1)
	amountInstructionsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountInstructionsBytes, amountInstructions)

	// set the index
	index := uint64(0)
	indexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBytes, index)

	// set the value
	value := uint8(54)

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindUint,
		Size8,
	}...)

	byteCode = append(byteCode, indexBytes...)
	byteCode = append(byteCode, []byte{
		OriginInline,
		value,
		EndInstruction,
	}...)

	byteCode = append(byteCode, remaining...)

	interpreter := NewInterpreter(byteCode)
	retStack, retBytes, err := interpreter.Execute()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err)
		return
	}

	if !bytes.Equal(retBytes, remaining) {
		t.Errorf("the returned bytes were invalid, expected: %v, returned; %v", retBytes, remaining)
		return
	}

	if stackValue, ok := retStack[KindUint][Size8][index]; ok {
		if stackValue != value {
			t.Errorf("the stack value (uint8[%d]) was expected to be %d, %d returned", index, value, stackValue)
			return
		}

		return
	}

	t.Errorf("the stack value (uint8[%d]) was expected to be defined", index)
}
