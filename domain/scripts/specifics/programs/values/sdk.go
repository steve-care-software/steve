package values

import "github.com/steve-care-software/steve/domain/hash"

// Value represents a value
type Value interface {
	Hash() hash.Hash
	IsMultiple() bool
	Multiple() ValueMultiple
	IsSingle() bool
	Single() ValueSingle
}

// ValueMultiple represents a value multiple
type ValueMultiple interface {
	IsVector() bool
	Vector() VectorValue
	IsList() bool
	List() ValueList
}

// VectorValue represents a vector value
type VectorValue interface {
	Length() uint64
	List() ValueList
}

// ValueList represents a value list
type ValueList interface {
	IsString() bool
	String() *string
	IsUint() bool
	Uint() UintSingleValue
	IsInt() bool
	Int() IntSingleValue
	IsFloat() bool
	Float() FloatSingleValue
	IsEngine() bool
	Engine() EngineListValue
}

// ValueMap represents a value map
type ValueMap interface {
	IsString() bool
	String() map[string]string
	IsUint() bool
	Uint() UintMapValue
	IsInt() bool
	Int() IntMapValue
	IsFloat() bool
	Float() FloatMapValue
	IsEngine() bool
	Engine() EngineMapValue
}

// ValueSingle represents a value list
type ValueSingle interface {
	IsString() bool
	String() []string
	IsUint() bool
	Uint() UintListValue
	IsInt() bool
	Int() IntListValue
	IsFloat() bool
	Float() FloatListValue
	IsEngine() bool
	Engine() EngineSingleValue
}

// EngineListValue represents an engine list value
type EngineListValue interface {
	IsPath() bool
	Path() []string
	IsRoute() bool
	Route() []Route
}

// EngineMapValue represents an engine map value
type EngineMapValue interface {
	IsPath() bool
	Path() map[string]string
	IsRoute() bool
	Route() map[string]Route
}

// EngineSingleValue represents a single engine value
type EngineSingleValue interface {
	IsPath() bool
	Path() []string
	IsRoute() bool
	Route() Route
}

// Route represents a route
type Route interface {
	Scope() uint8 // schema, context
	Origin() Point
	Target() Point
}

// Point repreents a point
type Point interface {
	Bucket() string
	Point() string
}

// UintListValue represents an uint list value
type UintListValue interface {
	IsHeight() bool
	Height() []uint8
	IsSixteen() bool
	Sixteen() []uint16
	IsThirtyTwo() bool
	ThirtyTwo() []uint32
	IsSixtyFour() bool
	SixtyFour() []uint64
}

// UintMapValue represents an uint map value
type UintMapValue interface {
	IsHeight() bool
	Height() map[string]uint8
	IsSixteen() bool
	Sixteen() map[string]uint16
	IsThirtyTwo() bool
	ThirtyTwo() map[string]uint32
	IsSixtyFour() bool
	SixtyFour() map[string]uint64
}

// UintSingleValue represents an uint single value
type UintSingleValue interface {
	IsHeight() bool
	Height() *uint8
	IsSixteen() bool
	Sixteen() *uint16
	IsThirtyTwo() bool
	ThirtyTwo() *uint32
	IsSixtyFour() bool
	SixtyFour() *uint64
}

// IntListValue represents an int list value
type IntListValue interface {
	IsHeight() bool
	Height() []int8
	IsSixteen() bool
	Sixteen() []int16
	IsThirtyTwo() bool
	ThirtyTwo() []int32
	IsSixtyFour() bool
	SixtyFour() []int64
}

// IntMapValue represents an int map value
type IntMapValue interface {
	IsHeight() bool
	Height() map[string]int8
	IsSixteen() bool
	Sixteen() map[string]int16
	IsThirtyTwo() bool
	ThirtyTwo() map[string]int32
	IsSixtyFour() bool
	SixtyFour() map[string]int64
}

// IntSingleValue represents an int single value
type IntSingleValue interface {
	IsHeight() bool
	Height() *uint8
	IsSixteen() bool
	Sixteen() *uint16
	IsThirtyTwo() bool
	ThirtyTwo() *uint32
	IsSixtyFour() bool
	SixtyFour() *uint64
}

// FloatListValue represents a float list value
type FloatListValue interface {
	IsThirtyTwo() bool
	Thirtytwo() []float32
	IsSixtyFour() bool
	SixtyFour() []float64
}

// FloatMapValue represents a float map value
type FloatMapValue interface {
	IsThirtyTwo() bool
	Thirtytwo() map[string]float32
	IsSixtyFour() bool
	SixtyFour() map[string]float64
}

// FloatSingleValue represents a float single value
type FloatSingleValue interface {
	IsThirtyTwo() bool
	Thirtytwo() *float32
	IsSixtyFour() bool
	SixtyFour() *float64
}
