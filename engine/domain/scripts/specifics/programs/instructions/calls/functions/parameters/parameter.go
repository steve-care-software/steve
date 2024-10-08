package parameters

import "github.com/steve-care-software/steve/hash"

type parameter struct {
	hash    hash.Hash
	current string
	local   string
}

func createParameter(
	hash hash.Hash,
	current string,
	local string,
) Parameter {
	out := parameter{
		hash:    hash,
		current: current,
		local:   local,
	}

	return &out
}

// Hash returns the hash
func (obj *parameter) Hash() hash.Hash {
	return obj.hash
}

// Current returns the current
func (obj *parameter) Current() string {
	return obj.current
}

// Local returns the local
func (obj *parameter) Local() string {
	return obj.local
}
