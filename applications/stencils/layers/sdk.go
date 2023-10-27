package layers

import (
	applications_administrator "github.com/steve-care-software/steve/applications/accounts/administrators"
	"github.com/steve-care-software/steve/domain/accounts/administrators"
	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols/layers"
	"github.com/steve-care-software/steve/domain/stencils/queries"
	result_layers "github.com/steve-care-software/steve/domain/stencils/results/executions"
)

// Application represents the computer application
type Application interface {
	Execute(query queries.Query) (result_layers.Execution, error)
}

// ValueBuilder represents a value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithBytes(bytes []byte) ValueBuilder
	WithLayer(layer layers.Layer) ValueBuilder
	WithResult(result result_layers.Execution) ValueBuilder
	WithAccount(account Account) ValueBuilder
	Now() (Value, error)
}

// Value represents an assignable value
type Value interface {
	IsBytes() bool
	Bytes() []byte
	IsLayer() bool
	Layer() layers.Layer
	IsResult() bool
	Result() result_layers.Execution
	IsAccount() bool
	Account() Account
}

// AccountBuilder represents an account builder
type AccountBuilder interface {
	Create() AccountBuilder
	WithApplication(app applications_administrator.Application) AccountBuilder
	WithAdministrator(adminIns administrators.Administrator) AccountBuilder
	Now() (Account, error)
}

// Account represents an account
type Account interface {
	IsApplication() bool
	Application() applications_administrator.Application
	IsAdministrator() bool
	Administrator() administrators.Administrator
}
