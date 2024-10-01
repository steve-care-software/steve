package suites

import "github.com/steve-care-software/steve/engine/domain/hash"

type suite struct {
	hash   hash.Hash
	name   string
	value  []byte
	isFail bool
}

func createSuite(
	hash hash.Hash,
	name string,
	value []byte,
	isFail bool,
) Suite {
	out := suite{
		hash:   hash,
		name:   name,
		value:  value,
		isFail: isFail,
	}

	return &out
}

// Hash returns the hash
func (obj *suite) Hash() hash.Hash {
	return obj.hash
}

// Name returns the name
func (obj *suite) Name() string {
	return obj.name
}

// Value returns the value
func (obj *suite) Value() []byte {
	return obj.value
}

// IsFail returns true if the suite is expected to fail, false otherwise
func (obj *suite) IsFail() bool {
	return obj.isFail
}
