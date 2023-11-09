package successes

import (
	"github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors/successes/decrypts"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors/successes/publickeys"
)

// Builder represents the encryptor builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithDecrypt(decrypt decrypts.Decrypt) Builder
	WithPublicKey(publicKey publickeys.PublicKey) Builder
	WithBytes(bytes []byte) Builder
	Now() (Success, error)
}

// Success represents a success
type Success interface {
	Variable() string
	Content() Content
}

// Content represents a success content
type Content interface {
	IsDecrypt() bool
	Decrypt() decrypts.Decrypt
	IsPublicKey() bool
	PublicKey() publickeys.PublicKey
	IsBytes() bool
	Bytes() []byte
}
