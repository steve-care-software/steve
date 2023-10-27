package identities

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/dashboards"
	"github.com/steve-care-software/steve/domain/accounts/identities/encryptors"
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles"
	"github.com/steve-care-software/steve/domain/accounts/identities/shares"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers"
)

// Identity represents an identity
type Identity interface {
	Dashboard() dashboards.Dashboard
	Encryptor() encryptors.Encryptor
	Signer() signers.Signer
	Profile() profiles.Profile
	HasConnections() bool
	Connections() profiles.Connections
	HasShares() bool
	Shares() shares.Shares
}
