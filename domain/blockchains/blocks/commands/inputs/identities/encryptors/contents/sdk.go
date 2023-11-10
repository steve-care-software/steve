package contents

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/encryptors/contents/decrypts"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/encryptors/contents/publickeys"
)

// Builder represents the encryptor builder
type Builder interface {
	Create() Builder
	WithDecrypt(decrypt decrypts.Decrypt) Builder
	WithPublicKey(pubKey publickeys.PublicKey) Builder
	IsBytes() Builder
	Now() (Content, error)
}

// Content represents an content
type Content interface {
	IsBytes() bool
	IsDecrypt() bool
	Decrypt() decrypts.Decrypt
	IsPublicKey() bool
	PublicKey() publickeys.PublicKey
}
