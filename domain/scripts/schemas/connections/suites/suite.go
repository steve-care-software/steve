package suites

import "github.com/steve-care-software/steve/domain/hash"

type suite struct {
	hash   hash.Hash
	name   string
	path   []string
	isFail bool
}

func createSuite(
	hash hash.Hash,
	name string,
	path []string,
	isFail bool,
) Suite {
	out := suite{
		hash:   hash,
		name:   name,
		path:   path,
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

// Path returns the path
func (obj *suite) Path() []string {
	return obj.path
}

// IsFail returns true if the suite is expected to fail, false otherwise
func (obj *suite) IsFail() bool {
	return obj.isFail
}
