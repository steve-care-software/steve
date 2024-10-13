package interpreters

import "encoding/binary"

type testSuite struct {
	params        map[uint8]map[uint8]map[uint64]any
	byteCode      []byte
	remaining     []byte
	expectedStack map[uint8]map[uint8]map[uint64]any
}

func createAssign_uint8_inline_withRemaining() testSuite {
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

	expectedStack := map[uint8]map[uint8]map[uint64]any{
		KindUint: {
			Size8: {
				index: value,
			},
		},
	}

	return testSuite{
		params:        map[uint8]map[uint8]map[uint64]any{},
		byteCode:      byteCode,
		remaining:     remaining,
		expectedStack: expectedStack,
	}
}

func createAssign_uint8_stack_withRemaining() testSuite {
	// set the amount of instructions:
	amountInstructions := uint64(1)
	amountInstructionsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountInstructionsBytes, amountInstructions)

	// set the index
	index := uint64(1)
	indexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBytes, index)

	// set the param index
	paramIndex := uint64(0)
	paramIndexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(paramIndexBytes, paramIndex)

	// set the value:
	value := uint8(22)

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
		OriginStack,
	}...)

	byteCode = append(byteCode, paramIndexBytes...)
	byteCode = append(byteCode, EndInstruction)

	byteCode = append(byteCode, remaining...)

	expectedStack := map[uint8]map[uint8]map[uint64]any{
		KindUint: {
			Size8: {
				paramIndex: value,
				index:      value,
			},
		},
	}

	return testSuite{
		params: map[uint8]map[uint8]map[uint64]any{
			KindUint: {
				Size8: {
					paramIndex: value,
				},
			},
		},
		byteCode:      byteCode,
		remaining:     remaining,
		expectedStack: expectedStack,
	}

}

func createAssign_uint16_inline_withRemaining() testSuite {
	// set the amount of instructions:
	amountInstructions := uint64(1)
	amountInstructionsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountInstructionsBytes, amountInstructions)

	// set the index
	index := uint64(0)
	indexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBytes, index)

	// set the value
	value := uint16(32342)
	valueBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(valueBytes, value)

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindUint,
		Size16,
	}...)

	byteCode = append(byteCode, indexBytes...)
	byteCode = append(byteCode, []byte{
		OriginInline,
	}...)

	byteCode = append(byteCode, valueBytes...)
	byteCode = append(byteCode, EndInstruction)

	byteCode = append(byteCode, remaining...)

	expectedStack := map[uint8]map[uint8]map[uint64]any{
		KindUint: {
			Size16: {
				index: value,
			},
		},
	}

	return testSuite{
		params:        map[uint8]map[uint8]map[uint64]any{},
		byteCode:      byteCode,
		remaining:     remaining,
		expectedStack: expectedStack,
	}
}

func createAssign_uint16_stack_withRemaining() testSuite {
	// set the amount of instructions:
	amountInstructions := uint64(1)
	amountInstructionsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountInstructionsBytes, amountInstructions)

	// set the index
	index := uint64(0)
	indexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBytes, index)

	// set the param index
	paramIndex := uint64(0)
	paramIndexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(paramIndexBytes, paramIndex)

	// set the value
	value := uint16(32342)
	valueBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(valueBytes, value)

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindUint,
		Size16,
	}...)

	byteCode = append(byteCode, indexBytes...)
	byteCode = append(byteCode, []byte{
		OriginStack,
	}...)

	byteCode = append(byteCode, paramIndexBytes...)
	byteCode = append(byteCode, EndInstruction)

	byteCode = append(byteCode, remaining...)

	expectedStack := map[uint8]map[uint8]map[uint64]any{
		KindUint: {
			Size16: {
				index: value,
			},
		},
	}

	return testSuite{
		params: map[uint8]map[uint8]map[uint64]any{
			KindUint: {
				Size16: {
					paramIndex: value,
				},
			},
		},
		byteCode:      byteCode,
		remaining:     remaining,
		expectedStack: expectedStack,
	}

}
