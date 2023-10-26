package preparations

import "github.com/steve-care-software/steve/domain/stencils/pointers"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewPreparationBuilder creates a new preparation builder
func NewPreparationBuilder() PreparationBuilder {
	return createPreparationBuilder()
}

// NewConditionBuilder creates a new condition builder
func NewConditionBuilder() ConditionBuilder {
	return createConditionBuilder()
}

// Builder represents a preparations builder
type Builder interface {
	Create() Builder
	WithList(list []Preparation) Builder
	Now() (Preparations, error)
}

// Preparations represents preparations
type Preparations interface {
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
	Variable() string
	Preparations() Preparations
}
