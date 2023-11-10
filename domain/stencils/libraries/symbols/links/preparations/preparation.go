package preparations

import (
	"github.com/steve-care-software/steve/domain/blockchains/hash"
	"github.com/steve-care-software/steve/domain/pointers"
)

type preparation struct {
	hash      hash.Hash
	isStop    bool
	load      pointers.Pointer
	exists    pointers.Pointer
	condition Condition
}

func createPreparationWithStop(
	hash hash.Hash,
) Preparation {
	return createPreparationInternally(hash, true, nil, nil, nil)
}

func createPreparationWithLoad(
	hash hash.Hash,
	load pointers.Pointer,
) Preparation {
	return createPreparationInternally(hash, false, load, nil, nil)
}

func createPreparationWithExists(
	hash hash.Hash,
	exists pointers.Pointer,
) Preparation {
	return createPreparationInternally(hash, false, nil, exists, nil)
}

func createPreparationWithCondition(
	hash hash.Hash,
	condition Condition,
) Preparation {
	return createPreparationInternally(hash, false, nil, nil, condition)
}

func createPreparationInternally(
	hash hash.Hash,
	isStop bool,
	load pointers.Pointer,
	exists pointers.Pointer,
	condition Condition,
) Preparation {
	out := preparation{
		hash:      hash,
		isStop:    isStop,
		load:      load,
		exists:    exists,
		condition: condition,
	}

	return &out
}

// Hash returns the hash
func (obj *preparation) Hash() hash.Hash {
	return obj.hash
}

// IsStop returns true if stop, false otherwise
func (obj *preparation) IsStop() bool {
	return obj.isStop
}

// IsLoad returns true if load, false otherwise
func (obj *preparation) IsLoad() bool {
	return obj.load != nil
}

// Load returns the load, if any
func (obj *preparation) Load() pointers.Pointer {
	return obj.load
}

// IsExists returns true if exists, false otherwise
func (obj *preparation) IsExists() bool {
	return obj.exists != nil
}

// Exists returns the exists, if any
func (obj *preparation) Exists() pointers.Pointer {
	return obj.exists
}

// IsCondition returns true if condition, false otherwise
func (obj *preparation) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *preparation) Condition() Condition {
	return obj.condition
}
