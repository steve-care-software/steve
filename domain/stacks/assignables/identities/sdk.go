package identities

import (
	"github.com/steve-care-software/steve/domain/accounts/identities"
	"github.com/steve-care-software/steve/domain/accounts/identities/encryptors"
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles"
	"github.com/steve-care-software/steve/domain/accounts/identities/shares"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers"
)

// Identity represents an identity
type Identity interface {
	IsInstance() bool
	Instance() identities.Identity
	IsEncryptor() bool
	Encryptor() encryptors.Encryptor
	IsSigner() bool
	Signer() signers.Signer
	IsProfile() bool
	Profile() profiles.Profile
	IsConnections() bool
	Connections() profiles.Connections
	IsShares() bool
	Shares() shares.Shares
}
