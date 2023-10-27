package connections

import "github.com/steve-care-software/steve/domain/accounts/identities/signers/signatures"

// Connection represents a connection
type Connection interface {
	Input() []byte
	Signature() signatures.Signature
}
