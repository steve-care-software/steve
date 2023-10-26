package preparations

import "github.com/steve-care-software/steve/domain/stencils/pointers/symbols"

// Builder represents preparations builder
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
	WithIndex(index uint) PreparationBuilder
	WithLoad(load symbols.Symbol) PreparationBuilder
	WithExists(exists symbols.Symbol) PreparationBuilder
	WithCondition(condition Preparations) PreparationBuilder
	IsStop() PreparationBuilder
	Now() (Preparation, error)
}

// Preparation represents preparation result
type Preparation interface {
	Index() uint
	Content() Content
}

// Content represents a preparation content
type Content interface {
	IsStop() bool
	IsLoad() bool
	Load() symbols.Symbol
	IsExists() bool
	Exists() symbols.Symbol
	IsCondition() bool
	Condition() Preparations
}
