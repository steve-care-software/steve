package requests

import "github.com/steve-care-software/steve/domain/accounts/identities/signers/signatures"

// Request represents an authorization request
type Request interface {
	Name() string
	Token() []byte
	Signature() signatures.Signature
}
