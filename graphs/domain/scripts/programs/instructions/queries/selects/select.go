package selects

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/programs/instructions/queries/conditions"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references/externals"
)

type selectData struct {
	externals externals.Externals
	isDelete  bool
	condition conditions.Condition
}

func createSelect(
	externals externals.Externals,
	isDelete bool,
) Select {
	return createSelectInternally(
		externals,
		isDelete,
		nil,
	)
}

func createSelectWithCondition(
	externals externals.Externals,
	isDelete bool,
	condition conditions.Condition,
) Select {
	return createSelectInternally(
		externals,
		isDelete,
		condition,
	)
}

func createSelectInternally(
	externals externals.Externals,
	isDelete bool,
	condition conditions.Condition,
) Select {
	out := selectData{
		externals: externals,
		isDelete:  isDelete,
		condition: condition,
	}

	return &out
}

// IsDelete returns true if the select is for deletion
func (obj *selectData) IsDelete() bool {
	return obj.isDelete
}

// Externals returns the externals of the select
func (obj *selectData) Externals() externals.Externals {
	return obj.externals
}

// HasCondition returns true if the select has a condition
func (obj *selectData) HasCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition of the select
func (obj *selectData) Condition() conditions.Condition {
	return obj.condition
}
