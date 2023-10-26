package identities

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/encryptors"
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers"
)

// Identity represents an identity
type Identity interface {
	Encryptor() encryptors.Encryptor
	Signer() signers.Signer
	Profile() profiles.Profile
	Connections() profiles.Connections
}
