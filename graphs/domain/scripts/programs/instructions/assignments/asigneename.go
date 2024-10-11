package assignments

import "github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"

type assigneeName struct {
	references references.References
	index      *uint
}

func createAssigneeName(references references.References) AssigneeName {
	return &assigneeName{
		references: references,
	}
}

func createAssigneeNameWithIndex(references references.References, index uint) AssigneeName {
	return &assigneeName{
		references: references,
		index:      &index,
	}
}

// References returns the references
func (obj *assigneeName) References() references.References {
	return obj.references
}

// HasIndex returns true if the index is present
func (obj *assigneeName) HasIndex() bool {
	return obj.index != nil
}

// Index returns the index if present
func (obj *assigneeName) Index() *uint {
	return obj.index
}
