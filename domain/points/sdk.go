package points

import (
	"github.com/steve-care-software/steve/domain/connections"
	"github.com/steve-care-software/steve/domain/points/programs"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewPointBuilder creates a new point builder
func NewPointBuilder() PointBuilder {
	return createPointBuilder()
}

// Builder represents the points builder
type Builder interface {
	Create() Builder
	WithList(list []Point) Builder
	Now() (Points, error)
}

// Points represents points
type Points interface {
	List() []Point
}

// PointBuilder represents the point builder
type PointBuilder interface {
	Create() PointBuilder
	WithConnection(connection connections.Connection) PointBuilder
	From(from []byte) PointBuilder
	Now() (Point, error)
}

// Point represents a point
type Point interface {
	Connection() connections.Connection
	From() []byte
	//HasNode() bool
	//Node() Node
}

// NodeBuilder represents a node builder
type NodeBuilder interface {
	Create() NodeBuilder
	WithParameter(parameter Parameter) NodeBuilder
	WithProgram(program programs.Program) NodeBuilder
	Now() (Node, error)
}

// Node represents a node
type Node interface {
	IsParameter() bool
	Parameter() Parameter
	IsProgram() bool
	Program() programs.Program
}

// ParameterBuilder represents the parameter builder
type ParameterBuilder interface {
	Create() ParameterBuilder
	WithContent(content []byte) ParameterBuilder
	WithProgram(program Point) ParameterBuilder
	Now() (Parameter, error)
}

// Parameter represents a parameter
type Parameter interface {
	Content() []byte
	Program() Point
}
