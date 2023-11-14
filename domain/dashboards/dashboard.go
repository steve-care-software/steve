package dashboards

import "github.com/steve-care-software/steve/domain/stencils"

type dashboard struct {
	root    stencils.Stencil
	visitor stencils.Stencil
	library stencils.Stencils
}

func createDashboard(
	root stencils.Stencil,
	visitor stencils.Stencil,
	library stencils.Stencils,
) Dashboard {
	out := dashboard{
		root:    root,
		visitor: visitor,
		library: library,
	}

	return &out
}

// Root returns the root stencil
func (obj *dashboard) Root() stencils.Stencil {
	return obj.root
}

// Visitor returns the visitor stencil
func (obj *dashboard) Visitor() stencils.Stencil {
	return obj.visitor
}

// Library returns the library
func (obj *dashboard) Library() stencils.Stencils {
	return obj.library
}
