package interpreters

const (
	// BeginInstruction represents a begin instruction
	BeginInstruction (uint8) = iota

	// EndInstruction represents an end of instruction
	EndInstruction
)

const (
	// OperationPointer represents a pointer operation
	OperationPointer (uint8) = iota

	// OperationAddition represents an addition operation
	OperationAddition

	// OperationSubstraction represents a substraction
	OperationSubstraction

	// OperationMultiplication represents a multiplication
	OperationMultiplication

	// OperationDivision represents a division
	OperationDivision

	// OperationModulo represents a modulo
	OperationModulo

	// OperationEqual represents an equal operation
	OperationEqual

	// OperationNot represents a not operation
	OperationNot

	// OperationJumpNext represents a jump next
	OperationJumpNext

	// OperationJumpTo represents a jump to
	OperationJumpTo

	// OperationAssignment represents an assignment
	OperationAssignment

	// OperationHash512 represents an hash512 operation
	OperationHash512

	// OperationPush represents a push operation
	OperationPush

	// OperationPop represents a pop operation
	OperationPop

	// OperationStop represents a stop operation
	OperationStop

	// OperationOpenFile represents an open file operation
	OperationOpenFile

	// OperationCloseFile represents a close file operation
	OperationCloseFile

	// OperationSeek represents a seek operation
	OperationSeek

	// OperationPositionAt represents a positionAt operation
	OperationPositionAt

	// OperationReadAt represents a reatAt operation
	OperationReadAt

	// OperationWriteAt represents a writeAt operation
	OperationWriteAt
)

const (
	// ContainerList represents the list container
	ContainerList (uint8) = iota

	// ContainerSet represents the set container
	ContainerSet

	// ContainerSortedSet represents the sorted set container
	ContainerSortedSet

	// ContainerMap represents the map container
	ContainerMap
)

const (
	// KindUint represents the uint kind
	KindUint (uint8) = iota

	// KindInt represents the int kind
	KindInt

	// KindFloat represents the float kind
	KindFloat

	// KindBool represents the bool kind
	KindBool
)

const (
	// OriginStack represents a variable origin
	OriginStack (uint8) = iota

	// OriginStatic represents a static origin
	OriginStatic
)

const (
	// PositionBegin represents the begin position
	PositionBegin (uint8) = iota

	// PositionCurrent represents the current position
	PositionCurrent

	// PositionEnd represents the end position
	PositionEnd
)

// NewInterpreter creates a new interpreter
func NewInterpreter(
	instructions []byte,
	params map[string][]byte,
) Interpreter {
	return createInterpreter(
		instructions,
		params,
	)
}

// Interpreter represents the vm interpreter
type Interpreter interface {
	Execute() ([]byte, error)
}
