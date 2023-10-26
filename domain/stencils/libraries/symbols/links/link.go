package links

import (
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/executions"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/origins"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/preparations"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/links/suites"
)

type link struct {
	origins      origins.Origins
	execution    executions.Execution
	preparations preparations.Preparations
	suites       suites.Suites
}

func createLink(
	origins origins.Origins,
	execution executions.Execution,
	preparations preparations.Preparations,
) Link {
	return createLinkInternally(origins, execution, preparations, nil)
}

func createLinkWithSuites(
	origins origins.Origins,
	execution executions.Execution,
	preparations preparations.Preparations,
	suites suites.Suites,
) Link {
	return createLinkInternally(origins, execution, preparations, suites)
}

func createLinkInternally(
	origins origins.Origins,
	execution executions.Execution,
	preparations preparations.Preparations,
	suites suites.Suites,
) Link {
	out := link{
		origins:      origins,
		execution:    execution,
		preparations: preparations,
		suites:       suites,
	}

	return &out
}

// Origins returns the origins
func (Obj *link) Origins() origins.Origins {
	return Obj.origins
}

// Execution returns the execution
func (Obj *link) Execution() executions.Execution {
	return Obj.execution
}

// Preparations returns the preparations
func (Obj *link) Preparations() preparations.Preparations {
	return Obj.preparations
}

// HasSuites returns true if there is suites, false otherwise
func (Obj *link) HasSuites() bool {
	return Obj.suites != nil
}

// Suites returns the suites, if any
func (Obj *link) Suites() suites.Suites {
	return Obj.suites
}
