package assignments

import "github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/containers/kinds"

type assignee struct {
	name AssigneeName
	kind kinds.Kind
}

func createAssignee(name AssigneeName) Assignee {
	return createAssigneeInternally(name, nil)
}

func createAssigneeWithKind(name AssigneeName, kind kinds.Kind) Assignee {
	return createAssigneeInternally(name, kind)
}

func createAssigneeInternally(name AssigneeName, kind kinds.Kind) Assignee {
	return &assignee{
		name: name,
		kind: kind,
	}
}

// Name returns the name of the assignee
func (obj *assignee) Name() AssigneeName {
	return obj.name
}

// HasKind checks if the kind field is set
func (obj *assignee) HasKind() bool {
	return obj.kind != nil
}

// Kind returns the kind if it exists
func (obj *assignee) Kind() kinds.Kind {
	if obj.HasKind() {
		return obj.kind
	}
	return nil
}
