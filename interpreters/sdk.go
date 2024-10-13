package interpreters

const (
	// BeginInstruction represents a begin instruction
	BeginInstruction (uint8) = iota

	// EndInstruction represents an end of instruction
	EndInstruction
)

const (
	// InstructionAssignment represents an assignment
	InstructionAssignment (uint8) = iota

	// InstructionCast represents a cast instruction
	InstructionCast

	// InstructionPointer represents a pointer operation
	InstructionPointer

	// InstructionAddition represents an addition operation
	InstructionAddition

	// InstructionSubstraction represents a substraction
	InstructionSubstraction

	// InstructionMultiplication represents a multiplication
	InstructionMultiplication

	// InstructionDivision represents a division
	InstructionDivision

	// InstructionModulo represents a modulo
	InstructionModulo

	// InstructionEqual represents an equal operation
	InstructionEqual

	// InstructionNot represents a not operation
	InstructionNot

	// InstructionJumpNext represents a jump next
	InstructionJumpNext

	// InstructionJumpTo represents a jump to
	InstructionJumpTo

	// InstructionHash512 represents an hash512 operation
	InstructionHash512

	// InstructionPush represents a push operation
	InstructionPush

	// InstructionPop represents a pop operation
	InstructionPop

	// InstructionStop represents a stop operation
	InstructionStop

	// InstructionFileOpen represents an open file operation
	InstructionFileOpen

	// InstructionFileClose represents a close file operation
	InstructionFileClose

	// InstructionFileSeek represents a seek operation
	InstructionFileSeek

	// InstructionFilePositionAt represents a positionAt operation
	InstructionFilePositionAt

	// InstructionFileReadAt represents a reatAt operation
	InstructionFileReadAt

	// InstructionFileWriteAt represents a writeAt operation
	InstructionFileWriteAt

	// InstructionFileRename reprgesents a name file operation
	InstructionFileRename
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

	// KindPointer represents the pointer kind
	KindPointer
)

const (
	// Size8 represents the size 8
	Size8 (uint8) = iota

	// Size16 represents the size 16
	Size16

	// Size32 represents the size 32
	Size32

	// Size64 represents the size 64
	Size64
)

const (
	// OriginStack represents a variable origin
	OriginStack (uint8) = iota

	// OriginInline represents an inline origin
	OriginInline
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
	params map[uint8]map[uint8]map[uint64]any,
) Interpreter {
	return createInterpreter(
		instructions,
		params,
	)
}

// Interpreter represents the vm interpreter
type Interpreter interface {
	Execute() (map[uint8]map[uint8]map[uint64]any, []byte, error)
}
