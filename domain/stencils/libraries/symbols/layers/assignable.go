package layers

import (
	"errors"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/accounts"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/constantvalues"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers/reduces"
)

type assignable struct {
	query   Query
	reduce  reduces.Reduce
	compare constantvalues.ConstantValues
	length  constantvalues.ConstantValue
	join    constantvalues.ConstantValues
	value   constantvalues.ConstantValue
	account accounts.Account
}

func createAssignableWithQuery(
	query Query,
) Assignable {
	return createAssignableInternally(query, nil, nil, nil, nil, nil, nil)
}

func createAssignableWithReduce(
	reduce reduces.Reduce,
) Assignable {
	return createAssignableInternally(nil, reduce, nil, nil, nil, nil, nil)
}

func createAssignableWithCompare(
	compare constantvalues.ConstantValues,
) Assignable {
	return createAssignableInternally(nil, nil, compare, nil, nil, nil, nil)
}

func createAssignableWithLength(
	length constantvalues.ConstantValue,
) Assignable {
	return createAssignableInternally(nil, nil, nil, length, nil, nil, nil)
}

func createAssignableWithJoin(
	join constantvalues.ConstantValues,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, join, nil, nil)
}

func createAssignableWithValue(
	value constantvalues.ConstantValue,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, value, nil)
}

func createAssignableWithAccount(
	account accounts.Account,
) Assignable {
	return createAssignableInternally(nil, nil, nil, nil, nil, nil, account)
}

func createAssignableInternally(
	query Query,
	reduce reduces.Reduce,
	compare constantvalues.ConstantValues,
	length constantvalues.ConstantValue,
	join constantvalues.ConstantValues,
	value constantvalues.ConstantValue,
	account accounts.Account,
) Assignable {
	out := assignable{
		query:   query,
		reduce:  reduce,
		compare: compare,
		length:  length,
		join:    join,
		value:   value,
		account: account,
	}

	return &out
}

// Hash returns the hash
func (obj *assignable) Hash() hash.Hash {
	if obj.IsQuery() {
		return obj.Query().Hash()
	}

	if obj.IsReduce() {
		return obj.Reduce().Hash()
	}

	if obj.IsCompare() {
		return obj.Compare().Hash()
	}

	if obj.IsLength() {
		return obj.length.Hash()
	}

	if obj.IsJoin() {
		return obj.join.Hash()
	}

	if obj.IsValue() {
		return obj.value.Hash()
	}

	panic(errors.New("finish account hash in assignable"))
	return nil
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
func (obj *assignable) Compare() constantvalues.ConstantValues {
	return obj.compare
}

// IsLength returns true if there is a length, false otherwise
func (obj *assignable) IsLength() bool {
	return obj.length != nil
}

// Length returns the length, if any
func (obj *assignable) Length() constantvalues.ConstantValue {
	return obj.length
}

// IsJoin returns true if there is a join, false otherwise
func (obj *assignable) IsJoin() bool {
	return obj.join != nil
}

// Join returns the length, if any
func (obj *assignable) Join() constantvalues.ConstantValues {
	return obj.join
}

// IsValue returns true if there is a value, false otherwise
func (obj *assignable) IsValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *assignable) Value() constantvalues.ConstantValue {
	return obj.value
}

// IsAccount returns true if there is an account, false otherwise
func (obj *assignable) IsAccount() bool {
	return obj.account != nil
}

// Account returns the account, if any
func (obj *assignable) Account() accounts.Account {
	return obj.account
}
