package responses

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles/authorizations/requests"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers/signatures"
)

// Response represents an authorization response
type Response interface {
	Request() requests.Request
	Signature() signatures.Signature
}
