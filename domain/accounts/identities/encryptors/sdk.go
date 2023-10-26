package encryptors

import "github.com/steve-care-software/steve/domain/accounts/identities/encryptors/publickeys"

// Encryptor represents an encryptor
type Encryptor interface {
	Public() publickeys.PublicKey
	Decrypt(cipher []byte) ([]byte, error)
	Bytes() []byte
}
