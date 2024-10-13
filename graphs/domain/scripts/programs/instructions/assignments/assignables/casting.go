package assignables

import (
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds"
)

type casting struct {
	assignable Assignable
	kind       kinds.Kind
}

func createCasting(
	assignable Assignable,
	kind kinds.Kind,
) Casting {
	return &casting{
		assignable: assignable,
		kind:       kind,
	}
}

// Assignable returns the assignable instance
func (obj *casting) Assignable() Assignable {
	return obj.assignable
}

// Kind returns the kind instance
func (obj *casting) Kind() kinds.Kind {
	return obj.kind
}
