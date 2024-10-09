package assignments

import "github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references/externals"

type assignment struct {
	external externals.External
	variable string
}

func createAssignment(external externals.External, variable string) Assignment {
	return &assignment{
		external: external,
		variable: variable,
	}
}

// External returns the external of the assignment
func (obj *assignment) External() externals.External {
	return obj.external
}

// Variable returns the variable of the assignment
func (obj *assignment) Variable() string {
	return obj.variable
}
