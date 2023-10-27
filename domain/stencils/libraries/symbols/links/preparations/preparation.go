package preparations

import "github.com/steve-care-software/steve/domain/pointers"

type preparation struct {
	isStop    bool
	load      pointers.Pointer
	exists    pointers.Pointer
	condition Condition
}

func createPreparationWithStop() Preparation {
	return createPreparationInternally(true, nil, nil, nil)
}

func createPreparationWithLoad(
	load pointers.Pointer,
) Preparation {
	return createPreparationInternally(false, load, nil, nil)
}

func createPreparationWithExists(
	exists pointers.Pointer,
) Preparation {
	return createPreparationInternally(false, nil, exists, nil)
}

func createPreparationWithCondition(
	condition Condition,
) Preparation {
	return createPreparationInternally(false, nil, nil, condition)
}

func createPreparationInternally(
	isStop bool,
	load pointers.Pointer,
	exists pointers.Pointer,
	condition Condition,
) Preparation {
	out := preparation{
		isStop:    isStop,
		load:      load,
		exists:    exists,
		condition: condition,
	}

	return &out
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
