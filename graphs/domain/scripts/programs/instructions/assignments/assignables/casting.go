package assignables

import (
	"github.com/steve-care-software/steve/engine/domain/scripts/specifics/programs/instructions/assignments/assignables"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds"
)

type casting struct {
	assignable assignables.Assignable
	kind       kinds.Kind
}

func createCasting(
	assignable assignables.Assignable,
	kind kinds.Kind,
) Casting {
	return &casting{
		assignable: assignable,
		kind:       kind,
	}
}

// Assignable returns the assignable instance
func (obj *casting) Assignable() assignables.Assignable {
	return obj.assignable
}

// Kind returns the kind instance
func (obj *casting) Kind() kinds.Kind {
	return obj.kind
}
