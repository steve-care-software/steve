package interpreters

import (
	"encoding/binary"
)

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

func createAssign_uint32_inline_withRemaining() testSuite {
	// set the amount of instructions:
	amountInstructions := uint64(1)
	amountInstructionsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountInstructionsBytes, amountInstructions)

	// set the index
	index := uint64(0)
	indexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBytes, index)

	// set the value
	value := uint32(3234233343)
	valueBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(valueBytes, value)

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindUint,
		Size32,
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
			Size32: {
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

func createAssign_uint32_stack_withRemaining() testSuite {
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
	value := uint32(3234343432)
	valueBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(valueBytes, value)

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindUint,
		Size32,
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
			Size32: {
				index: value,
			},
		},
	}

	return testSuite{
		params: map[uint8]map[uint8]map[uint64]any{
			KindUint: {
				Size32: {
					paramIndex: value,
				},
			},
		},
		byteCode:      byteCode,
		remaining:     remaining,
		expectedStack: expectedStack,
	}

}

func createAssign_uint64_inline_withRemaining() testSuite {
	// set the amount of instructions:
	amountInstructions := uint64(1)
	amountInstructionsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountInstructionsBytes, amountInstructions)

	// set the index
	index := uint64(0)
	indexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBytes, index)

	// set the value
	value := uint64(3234223412341233343)
	valueBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valueBytes, value)

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindUint,
		Size64,
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
			Size64: {
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

func createAssign_uint64_stack_withRemaining() testSuite {
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
	value := uint64(3234342232323233432)
	valueBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valueBytes, value)

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindUint,
		Size64,
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
			Size64: {
				index: value,
			},
		},
	}

	return testSuite{
		params: map[uint8]map[uint8]map[uint64]any{
			KindUint: {
				Size64: {
					paramIndex: value,
				},
			},
		},
		byteCode:      byteCode,
		remaining:     remaining,
		expectedStack: expectedStack,
	}

}

func createAssign_int8_inline_withRemaining() testSuite {
	// set the amount of instructions:
	amountInstructions := uint64(1)
	amountInstructionsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountInstructionsBytes, amountInstructions)

	// set the index
	index := uint64(0)
	indexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBytes, index)

	// set the value
	value := int8(-54)

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindInt,
		Size8,
	}...)

	byteCode = append(byteCode, indexBytes...)
	byteCode = append(byteCode, []byte{
		OriginInline,
		uint8(value),
		EndInstruction,
	}...)

	byteCode = append(byteCode, remaining...)

	expectedStack := map[uint8]map[uint8]map[uint64]any{
		KindInt: {
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

func createAssign_int8_stack_withRemaining() testSuite {
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
	value := int8(22)

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindInt,
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
		KindInt: {
			Size8: {
				paramIndex: value,
				index:      value,
			},
		},
	}

	return testSuite{
		params: map[uint8]map[uint8]map[uint64]any{
			KindInt: {
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

func createAssign_int16_inline_withRemaining() testSuite {
	// set the amount of instructions:
	amountInstructions := uint64(1)
	amountInstructionsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountInstructionsBytes, amountInstructions)

	// set the index
	index := uint64(0)
	indexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBytes, index)

	// set the value
	value := int16(-23232)
	valueBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(valueBytes, uint16(value))

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindInt,
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
		KindInt: {
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

func createAssign_int16_stack_withRemaining() testSuite {
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

	// set the value
	value := int16(23232)
	valueBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(valueBytes, uint16(value))

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindInt,
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
		KindInt: {
			Size16: {
				paramIndex: value,
				index:      value,
			},
		},
	}

	return testSuite{
		params: map[uint8]map[uint8]map[uint64]any{
			KindInt: {
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

func createAssign_int32_inline_withRemaining() testSuite {
	// set the amount of instructions:
	amountInstructions := uint64(1)
	amountInstructionsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountInstructionsBytes, amountInstructions)

	// set the index
	index := uint64(0)
	indexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBytes, index)

	// set the value
	value := int32(-232233232)
	valueBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(valueBytes, uint32(value))

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindInt,
		Size32,
	}...)

	byteCode = append(byteCode, indexBytes...)
	byteCode = append(byteCode, []byte{
		OriginInline,
	}...)

	byteCode = append(byteCode, valueBytes...)
	byteCode = append(byteCode, EndInstruction)
	byteCode = append(byteCode, remaining...)

	expectedStack := map[uint8]map[uint8]map[uint64]any{
		KindInt: {
			Size32: {
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

func createAssign_int32_stack_withRemaining() testSuite {
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

	// set the value
	value := int32(232356462)
	valueBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(valueBytes, uint32(value))

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindInt,
		Size32,
	}...)

	byteCode = append(byteCode, indexBytes...)
	byteCode = append(byteCode, []byte{
		OriginStack,
	}...)

	byteCode = append(byteCode, paramIndexBytes...)
	byteCode = append(byteCode, EndInstruction)

	byteCode = append(byteCode, remaining...)

	expectedStack := map[uint8]map[uint8]map[uint64]any{
		KindInt: {
			Size32: {
				paramIndex: value,
				index:      value,
			},
		},
	}

	return testSuite{
		params: map[uint8]map[uint8]map[uint64]any{
			KindInt: {
				Size32: {
					paramIndex: value,
				},
			},
		},
		byteCode:      byteCode,
		remaining:     remaining,
		expectedStack: expectedStack,
	}

}

func createAssign_int64_inline_withRemaining() testSuite {
	// set the amount of instructions:
	amountInstructions := uint64(1)
	amountInstructionsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(amountInstructionsBytes, amountInstructions)

	// set the index
	index := uint64(0)
	indexBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBytes, index)

	// set the value
	value := int64(-2322334564366543232)
	valueBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valueBytes, uint64(value))

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindInt,
		Size64,
	}...)

	byteCode = append(byteCode, indexBytes...)
	byteCode = append(byteCode, []byte{
		OriginInline,
	}...)

	byteCode = append(byteCode, valueBytes...)
	byteCode = append(byteCode, EndInstruction)
	byteCode = append(byteCode, remaining...)

	expectedStack := map[uint8]map[uint8]map[uint64]any{
		KindInt: {
			Size64: {
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

func createAssign_int64_stack_withRemaining() testSuite {
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

	// set the value
	value := int64(2323564242342344162)
	valueBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(valueBytes, uint64(value))

	// set the remaining:
	remaining := []byte("this is some remaining data")

	byteCode := amountInstructionsBytes
	byteCode = append(byteCode, []byte{
		BeginInstruction,
		InstructionAssignment,
		KindInt,
		Size64,
	}...)

	byteCode = append(byteCode, indexBytes...)
	byteCode = append(byteCode, []byte{
		OriginStack,
	}...)

	byteCode = append(byteCode, paramIndexBytes...)
	byteCode = append(byteCode, EndInstruction)

	byteCode = append(byteCode, remaining...)

	expectedStack := map[uint8]map[uint8]map[uint64]any{
		KindInt: {
			Size64: {
				paramIndex: value,
				index:      value,
			},
		},
	}

	return testSuite{
		params: map[uint8]map[uint8]map[uint64]any{
			KindInt: {
				Size64: {
					paramIndex: value,
				},
			},
		},
		byteCode:      byteCode,
		remaining:     remaining,
		expectedStack: expectedStack,
	}

}
