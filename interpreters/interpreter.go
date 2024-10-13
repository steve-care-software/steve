package interpreters

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type interpreter struct {
	instructions []byte
	stackUint8   map[uint64]uint8
}

func createInterpreter(
	instructions []byte,
) Interpreter {
	out := interpreter{
		instructions: instructions,
	}

	return out.init()
}

// Execute executes the interpreter
func (app *interpreter) Execute() ([]byte, error) {
	app.init()
	return app.execInstructions(app.instructions)
}

func (app *interpreter) init() Interpreter {
	app.stackUint8 = map[uint64]uint8{}
	return app
}

func (app *interpreter) execInstructions(input []byte) ([]byte, error) {
	remaining := input
	for {

		if len(remaining) <= 0 {
			break
		}

		retRemaining, isEnd, err := app.execInstruction(remaining)
		if err != nil {
			return nil, err
		}

		if isEnd {
			return remaining, nil
		}

		remaining = retRemaining
	}

	return remaining, nil
}

func (app *interpreter) execInstruction(input []byte) ([]byte, bool, error) {
	if len(input) <= 0 {
		return nil, true, nil
	}

	remaining := input[1:]
	switch input[0] {
	case InstructionAssignment:
		return app.execAssignment(remaining)
	default:
		str := fmt.Sprintf("the byte (%d) is not a valid instruction definer", input[0])
		return nil, false, errors.New(str)
	}
}

func (app *interpreter) execAssignment(input []byte) ([]byte, bool, error) {
	remaining := input[1:]
	switch input[0] {
	case KindUint:
		return app.execAssignmentUint(remaining)
	case KindInt:
		return app.execAssignmentInt(remaining)
	case KindFloat:
		return app.execAssignmentFloat(remaining)
	case KindBool:
		return app.execAssignmentBool(remaining)
	case KindPointer:
		return app.execAssignmentPointer(remaining)
	default:
		str := fmt.Sprintf("the byte (%d) is not a valid assignment definer", input[0])
		return nil, false, errors.New(str)
	}
}

func (app *interpreter) execAssignmentUint(input []byte) ([]byte, bool, error) {
	remaining := input[1:]
	switch input[0] {
	case Size8:
		return app.execAssignmentUint8(remaining)
	case Size16:
		return app.execAssignmentUint16(remaining)
	case Size32:
		return app.execAssignmentUint32(remaining)
	case Size64:
		return app.execAssignmentUint64(remaining)
	default:
		str := fmt.Sprintf("the byte (%d) is not a valid uint assignment definer", input[0])
		return nil, false, errors.New(str)
	}
}

func (app *interpreter) execAssignmentUint8(input []byte) ([]byte, bool, error) {
	// find the variable index:
	pIndex, isEnd, retRemaining, err := app.fetchValueUint64Inline(input)
	if err != nil {
		return nil, false, err
	}

	if isEnd {
		return nil, true, nil
	}

	// find the value:
	pValue, isEnd, retRemaining, err := app.fetchValueUint8(retRemaining)
	if err != nil {
		return nil, false, err
	}

	if isEnd {
		return nil, true, nil
	}

	// execute the assignment:
	app.stackUint8[*pIndex] = *pValue
	return retRemaining, false, nil
}

func (app *interpreter) execAssignmentUint16(input []byte) ([]byte, bool, error) {
	return nil, false, nil
}

func (app *interpreter) execAssignmentUint32(input []byte) ([]byte, bool, error) {
	return nil, false, nil
}

func (app *interpreter) execAssignmentUint64(input []byte) ([]byte, bool, error) {
	return nil, false, nil
}

func (app *interpreter) execAssignmentInt(input []byte) ([]byte, bool, error) {
	return nil, false, nil
}

func (app *interpreter) execAssignmentFloat(input []byte) ([]byte, bool, error) {
	return nil, false, nil
}

func (app *interpreter) execAssignmentBool(input []byte) ([]byte, bool, error) {
	return nil, false, nil
}

func (app *interpreter) execAssignmentPointer(input []byte) ([]byte, bool, error) {
	return nil, false, nil
}

func (app *interpreter) fetchValueUint8(input []byte) (*uint8, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, true, nil, nil
	}

	remaining := input[1:]
	switch input[0] {
	case OriginStack:
		return app.fetchValueUint8Stack(remaining)
	case OriginInline:
		return app.fetchValueUint8Inline(remaining)
	default:
		str := fmt.Sprintf("the byte (%d) is not a valid uint8 origin", input[0])
		return nil, false, nil, errors.New(str)
	}
}

func (app *interpreter) fetchValueUint8Stack(input []byte) (*uint8, bool, []byte, error) {
	pIndex, isEnd, retRemaining, err := app.fetchValueUint64Inline(input)
	if err != nil {
		return nil, false, nil, err
	}

	if isEnd {
		return nil, true, nil, nil
	}

	if value, ok := app.stackUint8[*pIndex]; ok {
		return &value, false, retRemaining, nil
	}

	str := fmt.Sprintf("the the value (index: %d) is not valid on the uint8 stack", *pIndex)
	return nil, false, nil, errors.New(str)
}

func (app *interpreter) fetchValueUint8Inline(input []byte) (*uint8, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, true, nil, nil
	}

	value := input[0]
	return &value, false, input[1:], nil
}

func (app *interpreter) fetchValueUint64Inline(input []byte) (*uint64, bool, []byte, error) {
	if len(input) <= 8 {
		return nil, true, nil, nil
	}

	// Convert byte slice to uint64 (little-endian)
	value := binary.LittleEndian.Uint64(input[:8])
	return &value, false, input[8:], nil
}
