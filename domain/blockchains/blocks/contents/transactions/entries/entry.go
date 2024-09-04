package entries

import "github.com/steve-care-software/steve/domain/hash"

type entry struct {
	hash   hash.Hash
	flag   hash.Hash
	script hash.Hash
	fees   uint64
}

func createEntry(
	hash hash.Hash,
	flag hash.Hash,
	script hash.Hash,
	fees uint64,
) Entry {
	out := entry{
		hash:   hash,
		flag:   flag,
		script: script,
		fees:   fees,
	}

	return &out
}

// Hash returns the hash
func (obj *entry) Hash() hash.Hash {
	return obj.hash
}

// Flag returns the flag
func (obj *entry) Flag() hash.Hash {
	return obj.flag
}

// Script returns the script
func (obj *entry) Script() hash.Hash {
	return obj.script
}

// Fees returns the fees
func (obj *entry) Fees() uint64 {
	return obj.fees
}
