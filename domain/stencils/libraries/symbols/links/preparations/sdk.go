package preparations

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/pointers"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewPreparationBuilder creates a new preparation builder
func NewPreparationBuilder() PreparationBuilder {
	hashAdapter := hash.NewAdapter()
	return createPreparationBuilder(
		hashAdapter,
	)
}

// NewConditionBuilder creates a new condition builder
func NewConditionBuilder() ConditionBuilder {
	hashAdapter := hash.NewAdapter()
	return createConditionBuilder(
		hashAdapter,
	)
}

// Builder represents a preparations builder
type Builder interface {
	Create() Builder
	WithList(list []Preparation) Builder
	Now() (Preparations, error)
}

// Preparations represents preparations
type Preparations interface {
	Hash() hash.Hash
	List() []Preparation
}

// PreparationBuilder represents a preparation builder
type PreparationBuilder interface {
	Create() PreparationBuilder
	WithLoad(load pointers.Pointer) PreparationBuilder
	WithExists(exists pointers.Pointer) PreparationBuilder
	WithCondition(condition Condition) PreparationBuilder
	IsStop() PreparationBuilder
	Now() (Preparation, error)
}

// Preparation represents a preparation
type Preparation interface {
	Hash() hash.Hash
	IsStop() bool
	IsLoad() bool
	Load() pointers.Pointer
	IsExists() bool
	Exists() pointers.Pointer
	IsCondition() bool
	Condition() Condition
}

// ConditionBuilder represents a condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithVariable(variable string) ConditionBuilder
	WithPreparations(preparations Preparations) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Variable() string
	Preparations() Preparations
}
