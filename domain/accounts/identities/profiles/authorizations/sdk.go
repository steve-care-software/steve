package authorizations

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles/authorizations/requests"
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles/authorizations/responses"
)

// Authorization represents an authorization
type Authorization interface {
	IsRequest() bool
	Request() requests.Request
	IsResponse() bool
	Response() responses.Response
}
