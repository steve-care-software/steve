package publickeys

import "github.com/steve-care-software/steve/domain/commands/inputs/identities/encryptors/contents/publickeys/encrypts"

// Builder represents a public key builder
type Builder interface {
	Create() Builder
	WithEncrypt(encrypt encrypts.Encrypt) Builder
	IsBytes() Builder
	Now() (PublicKey, error)
}

// PublicKey represents a public key
type PublicKey interface {
	IsBytes() bool
	IsEncrypt() bool
	Encrypt() encrypts.Encrypt
}
