package layers

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/reduces"
)

type assignable struct {
	query   Query
	reduce  reduces.Reduce
	compare ConstantValues
	length  ConstantValue
	join    ConstantValues
	value   ConstantValue
}

func createAssignableWithQuery(
	query Query,
) Assignable {
	return createAssignableInternally(query, nil, nil, nil, nil, nil)
}

func createAssignableWithReduce(
	reduce reduces.Reduce,
) Assignable {
	return createAssignableInternally(nil, reduce, nil, nil, nil, nil)
}

func createAssignableWithCompare(
	compare ConstantValues,
) Assignable {
	return createAssignableInternally(nil, nil, compare, nil, nil, nil)
}

func createAssignableWithLength(
	length ConstantValue,
) Assignable {
	return createAssignableInternally(nil, nil, nil, length, nil, nil)
}

func createAssignableWithJoin(
	join ConstantValues,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, join, nil)
}

func createAssignableWithValue(
	value ConstantValue,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, value)
}

func createAssignableInternally(
	query Query,
	reduce reduces.Reduce,
	compare ConstantValues,
	length ConstantValue,
	join ConstantValues,
	value ConstantValue,
) Assignable {
	out := assignable{
		query:   query,
		reduce:  reduce,
		compare: compare,
		length:  length,
		join:    join,
		value:   value,
	}

	return &out
}

// IsQuery returns true if there is a query, false otherwise
func (obj *assignable) IsQuery() bool {
	return obj.query != nil
}

// Query returns the query, if any
func (obj *assignable) Query() Query {
	return obj.query
}

// IsReduce returns true if there is a reduce, false otherwise
func (obj *assignable) IsReduce() bool {
	return obj.reduce != nil
}

// Reduce returns the reduce, if any
func (obj *assignable) Reduce() reduces.Reduce {
	return obj.reduce
}

// IsCompare returns true if there is a compare, false otherwise
func (obj *assignable) IsCompare() bool {
	return obj.compare != nil
}

// Compare returns the compare, if any
func (obj *assignable) Compare() ConstantValues {
	return obj.compare
}

// IsLength returns true if there is a length, false otherwise
func (obj *assignable) IsLength() bool {
	return obj.length != nil
}

// Length returns the length, if any
func (obj *assignable) Length() ConstantValue {
	return obj.length
}

// IsJoin returns true if there is a join, false otherwise
func (obj *assignable) IsJoin() bool {
	return obj.join != nil
}

// Join returns the length, if any
func (obj *assignable) Join() ConstantValues {
	return obj.join
}

// IsValue returns true if there is a value, false otherwise
func (obj *assignable) IsValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *assignable) Value() ConstantValue {
	return obj.value
}
