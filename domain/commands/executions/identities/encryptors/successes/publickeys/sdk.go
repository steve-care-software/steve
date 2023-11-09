package publickeys

import "github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors/successes/publickeys/encrypts"

// Builder represents a public key builder
type Builder interface {
	Create() Builder
	WithBytes(bytes []byte) Builder
	WithEncrypt(encrypt encrypts.Encrypt) Builder
	Now() (PublicKey, error)
}

// PublicKey represents a public key
type PublicKey interface {
	IsBytes() bool
	Bytes() []byte
	IsEncrypt() bool
	Encrypt() encrypts.Encrypt
}
