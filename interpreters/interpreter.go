package interpreters

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type interpreter struct {
	instructions []byte
	stack        map[uint8]map[uint8]map[uint64]any
}

func createInterpreter(
	instructions []byte,
	params map[uint8]map[uint8]map[uint64]any,
) Interpreter {
	out := interpreter{
		instructions: instructions,
		stack:        params,
	}

	return out.init()
}

// Execute executes the interpreter
func (app *interpreter) Execute() (map[uint8]map[uint8]map[uint64]any, []byte, error) {
	app.init()
	retRemaining, err := app.execInstructions(app.instructions)
	if err != nil {
		return nil, nil, err
	}

	// cleanup:
	output := map[uint8]map[uint8]map[uint64]any{}
	for idxFirst, oneMap := range app.stack {
		firstOutput := map[uint8]map[uint64]any{}
		for idxSecond, oneSubMap := range oneMap {
			if len(oneSubMap) <= 0 {
				continue
			}

			firstOutput[idxSecond] = oneSubMap
		}

		if len(firstOutput) <= 0 {
			continue
		}

		output[idxFirst] = firstOutput
	}

	// return:
	return output, retRemaining, nil
}

func (app *interpreter) init() Interpreter {
	/*
		uint8
	*/
	if _, ok := app.stack[KindUint]; !ok {
		app.stack[KindUint] = map[uint8]map[uint64]any{}
	}

	if _, ok := app.stack[KindUint][Size8]; !ok {
		app.stack[KindUint][Size8] = map[uint64]any{}
	}

	if _, ok := app.stack[KindUint][Size16]; !ok {
		app.stack[KindUint][Size16] = map[uint64]any{}
	}

	if _, ok := app.stack[KindUint][Size32]; !ok {
		app.stack[KindUint][Size32] = map[uint64]any{}
	}

	if _, ok := app.stack[KindUint][Size64]; !ok {
		app.stack[KindUint][Size64] = map[uint64]any{}
	}

	/*
		int8
	*/

	if _, ok := app.stack[KindInt]; !ok {
		app.stack[KindInt] = map[uint8]map[uint64]any{}
	}

	if _, ok := app.stack[KindInt][Size8]; !ok {
		app.stack[KindInt][Size8] = map[uint64]any{}
	}

	return app
}

func (app *interpreter) execInstructions(input []byte) ([]byte, error) {
	pAmount, isEnd, retRemaining, err := app.fetchValueUint64Inline(input)
	if err != nil {
		return nil, err
	}

	if isEnd {
		return nil, errors.New("the byteCode was expected to contain the amount of instructions contained in the execution, but the byteCode was empty")
	}

	amount := int(*pAmount)
	remaining := retRemaining
	for i := 0; i < amount; i++ {

		if len(remaining) <= 0 {
			break
		}

		retRemaining, isEnd, err := app.execInstructionLine(remaining)
		if err != nil {
			return nil, err
		}

		if isEnd {
			str := fmt.Sprintf("the byteCode was expected to contain %d instructions, the end of instructions was reached at index: %d", *pAmount, i)
			return nil, errors.New(str)
		}

		remaining = retRemaining
	}

	return remaining, nil
}

func (app *interpreter) execInstructionLine(input []byte) ([]byte, bool, error) {
	if len(input) <= 0 {
		return nil, false, nil
	}

	if input[0] != BeginInstruction {
		str := fmt.Sprintf("the byte (%d) was expected to be the begin instruction byte (%d)", input[0], BeginInstruction)
		return nil, false, errors.New(str)
	}

	retBytes, isEnd, err := app.execInstruction(input[1:])
	if err != nil {
		return nil, false, err
	}

	if isEnd {
		return nil, true, nil
	}

	if len(retBytes) <= 0 {
		str := fmt.Sprintf("the bytes were NOT expected to be empty, the end instruction byte (%d) was expected", EndInstruction)
		return nil, false, errors.New(str)
	}

	if retBytes[0] != EndInstruction {
		str := fmt.Sprintf("the byte (%d) was expected to be the end instruction byte (%d)", retBytes[0], EndInstruction)
		return nil, false, errors.New(str)
	}

	return retBytes[1:], false, nil
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

func (app *interpreter) execAssignmentInt(input []byte) ([]byte, bool, error) {
	remaining := input[1:]
	switch input[0] {
	case Size8:
		return app.execAssignmentInt8(remaining)
	/*case Size16:
		return app.execAssignmentInt16(remaining)
	case Size32:
		return app.execAssignmentInt32(remaining)
	case Size64:
		return app.execAssignmentInt64(remaining)*/
	default:
		str := fmt.Sprintf("the byte (%d) is not a valid int assignment definer", input[0])
		return nil, false, errors.New(str)
	}
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

/*
	int8
*/

func (app *interpreter) execAssignmentInt8(input []byte) ([]byte, bool, error) {
	// find the variable index:
	pIndex, isEnd, retRemaining, err := app.fetchValueUint64Inline(input)
	if err != nil {
		return nil, false, err
	}

	if isEnd {
		return nil, true, nil
	}

	// find the value:
	pValue, isEnd, retRemaining, err := app.fetchValueInt8(retRemaining)
	if err != nil {
		return nil, false, err
	}

	if isEnd {
		return nil, true, nil
	}

	// execute the assignment:
	app.stack[KindInt][Size8][*pIndex] = *pValue
	return retRemaining, false, nil
}

func (app *interpreter) fetchValueInt8(input []byte) (*int8, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, true, nil, nil
	}

	remaining := input[1:]
	switch input[0] {
	case OriginStack:
		return app.fetchValueInt8Stack(remaining)
	case OriginInline:
		return app.fetchValueInt8Inline(remaining)
	default:
		str := fmt.Sprintf("the byte (%d) is not a valid int8 origin", input[0])
		return nil, false, nil, errors.New(str)
	}
}

func (app *interpreter) fetchValueInt8Stack(input []byte) (*int8, bool, []byte, error) {
	pIndex, isEnd, retRemaining, err := app.fetchValueUint64Inline(input)
	if err != nil {
		return nil, false, nil, err
	}

	if isEnd {
		return nil, true, nil, nil
	}

	if value, ok := app.stack[KindInt][Size8][*pIndex]; ok {
		if casted, ok := value.(int8); ok {
			return &casted, false, retRemaining, nil
		}

		str := fmt.Sprintf("casting error: the stack value (index: %d) was expected to contain a int8 value", *pIndex)
		return nil, false, nil, errors.New(str)
	}

	str := fmt.Sprintf("the value (index: %d) is not valid on the int8 stack", *pIndex)
	return nil, false, nil, errors.New(str)
}

func (app *interpreter) fetchValueInt8Inline(input []byte) (*int8, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, true, nil, nil
	}

	value := int8(input[0])
	return &value, false, input[1:], nil
}

/*
	uint64
*/

func (app *interpreter) execAssignmentUint64(input []byte) ([]byte, bool, error) {
	// find the variable index:
	pIndex, isEnd, retRemaining, err := app.fetchValueUint64Inline(input)
	if err != nil {
		return nil, false, err
	}

	if isEnd {
		return nil, true, nil
	}

	// find the value:
	pValue, isEnd, retRemaining, err := app.fetchValueUint64(retRemaining)
	if err != nil {
		return nil, false, err
	}

	if isEnd {
		return nil, true, nil
	}

	// execute the assignment:
	app.stack[KindUint][Size64][*pIndex] = *pValue
	return retRemaining, false, nil
}

func (app *interpreter) fetchValueUint64(input []byte) (*uint64, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, true, nil, nil
	}

	remaining := input[1:]
	switch input[0] {
	case OriginStack:
		return app.fetchValueUint64Stack(remaining)
	case OriginInline:
		return app.fetchValueUint64Inline(remaining)
	default:
		str := fmt.Sprintf("the byte (%d) is not a valid uint64 origin", input[0])
		return nil, false, nil, errors.New(str)
	}
}

func (app *interpreter) fetchValueUint64Stack(input []byte) (*uint64, bool, []byte, error) {
	pIndex, isEnd, retRemaining, err := app.fetchValueUint64Inline(input)
	if err != nil {
		return nil, false, nil, err
	}

	if isEnd {
		return nil, true, nil, nil
	}

	if value, ok := app.stack[KindUint][Size64][*pIndex]; ok {
		if casted, ok := value.(uint64); ok {
			return &casted, false, retRemaining, nil
		}

		str := fmt.Sprintf("casting error: the stack value (index: %d) was expected to contain a uint64 value", *pIndex)
		return nil, false, nil, errors.New(str)
	}

	str := fmt.Sprintf("the value (index: %d) is not valid on the uint64 stack", *pIndex)
	return nil, false, nil, errors.New(str)
}

func (app *interpreter) fetchValueUint64Inline(input []byte) (*uint64, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, true, nil, nil
	}

	valueBytes := input[:8]
	value := binary.LittleEndian.Uint64(valueBytes)
	return &value, false, input[8:], nil
}

/*
	uint32
*/

func (app *interpreter) execAssignmentUint32(input []byte) ([]byte, bool, error) {
	// find the variable index:
	pIndex, isEnd, retRemaining, err := app.fetchValueUint64Inline(input)
	if err != nil {
		return nil, false, err
	}

	if isEnd {
		return nil, true, nil
	}

	// find the value:
	pValue, isEnd, retRemaining, err := app.fetchValueUint32(retRemaining)
	if err != nil {
		return nil, false, err
	}

	if isEnd {
		return nil, true, nil
	}

	// execute the assignment:
	app.stack[KindUint][Size32][*pIndex] = *pValue
	return retRemaining, false, nil
}

func (app *interpreter) fetchValueUint32(input []byte) (*uint32, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, true, nil, nil
	}

	remaining := input[1:]
	switch input[0] {
	case OriginStack:
		return app.fetchValueUint32Stack(remaining)
	case OriginInline:
		return app.fetchValueUint32Inline(remaining)
	default:
		str := fmt.Sprintf("the byte (%d) is not a valid uint32 origin", input[0])
		return nil, false, nil, errors.New(str)
	}
}

func (app *interpreter) fetchValueUint32Stack(input []byte) (*uint32, bool, []byte, error) {
	pIndex, isEnd, retRemaining, err := app.fetchValueUint64Inline(input)
	if err != nil {
		return nil, false, nil, err
	}

	if isEnd {
		return nil, true, nil, nil
	}

	if value, ok := app.stack[KindUint][Size32][*pIndex]; ok {
		if casted, ok := value.(uint32); ok {
			return &casted, false, retRemaining, nil
		}

		str := fmt.Sprintf("casting error: the stack value (index: %d) was expected to contain a uint32 value", *pIndex)
		return nil, false, nil, errors.New(str)
	}

	str := fmt.Sprintf("the value (index: %d) is not valid on the uint32 stack", *pIndex)
	return nil, false, nil, errors.New(str)
}

func (app *interpreter) fetchValueUint32Inline(input []byte) (*uint32, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, true, nil, nil
	}

	valueBytes := input[:4]
	value := binary.LittleEndian.Uint32(valueBytes)
	return &value, false, input[4:], nil
}

/*
	uint16
*/

func (app *interpreter) execAssignmentUint16(input []byte) ([]byte, bool, error) {
	// find the variable index:
	pIndex, isEnd, retRemaining, err := app.fetchValueUint64Inline(input)
	if err != nil {
		return nil, false, err
	}

	if isEnd {
		return nil, true, nil
	}

	// find the value:
	pValue, isEnd, retRemaining, err := app.fetchValueUint16(retRemaining)
	if err != nil {
		return nil, false, err
	}

	if isEnd {
		return nil, true, nil
	}

	// execute the assignment:
	app.stack[KindUint][Size16][*pIndex] = *pValue
	return retRemaining, false, nil
}

func (app *interpreter) fetchValueUint16(input []byte) (*uint16, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, true, nil, nil
	}

	remaining := input[1:]
	switch input[0] {
	case OriginStack:
		return app.fetchValueUint16Stack(remaining)
	case OriginInline:
		return app.fetchValueUint16Inline(remaining)
	default:
		str := fmt.Sprintf("the byte (%d) is not a valid uint16 origin", input[0])
		return nil, false, nil, errors.New(str)
	}
}

func (app *interpreter) fetchValueUint16Stack(input []byte) (*uint16, bool, []byte, error) {
	pIndex, isEnd, retRemaining, err := app.fetchValueUint64Inline(input)
	if err != nil {
		return nil, false, nil, err
	}

	if isEnd {
		return nil, true, nil, nil
	}

	if value, ok := app.stack[KindUint][Size16][*pIndex]; ok {
		if casted, ok := value.(uint16); ok {
			return &casted, false, retRemaining, nil
		}

		str := fmt.Sprintf("casting error: the stack value (index: %d) was expected to contain a uint16 value", *pIndex)
		return nil, false, nil, errors.New(str)
	}

	str := fmt.Sprintf("the value (index: %d) is not valid on the uint16 stack", *pIndex)
	return nil, false, nil, errors.New(str)
}

func (app *interpreter) fetchValueUint16Inline(input []byte) (*uint16, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, true, nil, nil
	}

	valueBytes := input[:2]
	value := binary.LittleEndian.Uint16(valueBytes)
	return &value, false, input[2:], nil
}

/*
	uint8
*/

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
	app.stack[KindUint][Size8][*pIndex] = *pValue
	return retRemaining, false, nil
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

	if value, ok := app.stack[KindUint][Size8][*pIndex]; ok {
		if casted, ok := value.(uint8); ok {
			return &casted, false, retRemaining, nil
		}

		str := fmt.Sprintf("casting error: the stack value (index: %d) was expected to contain a uint8 value", *pIndex)
		return nil, false, nil, errors.New(str)
	}

	str := fmt.Sprintf("the value (index: %d) is not valid on the uint8 stack", *pIndex)
	return nil, false, nil, errors.New(str)
}

func (app *interpreter) fetchValueUint8Inline(input []byte) (*uint8, bool, []byte, error) {
	if len(input) <= 0 {
		return nil, true, nil, nil
	}

	value := input[0]
	return &value, false, input[1:], nil
}
