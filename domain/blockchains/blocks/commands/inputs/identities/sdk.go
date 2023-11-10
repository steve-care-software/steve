package identities

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/connections"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/encryptors"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/identities"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/profiles"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/shares"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/identities/signers"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs/shares/dashboards"
)

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithIdentity(identity identities.Identity) Builder
	WithDashboard(dashboard dashboards.Dashboard) Builder
	WithEncryptor(encryptor encryptors.Encryptor) Builder
	WithSigner(signer signers.Signer) Builder
	WithProfile(profile profiles.Profile) Builder
	WithConnections(connections connections.Connections) Builder
	WithShares(shares shares.Shares) Builder
	Now() (Identity, error)
}

// Identity represents an identity input command
type Identity interface {
	IsIdentity() bool
	Identity() identities.Identity
	IsDashboard() bool
	Dashboard() dashboards.Dashboard
	IsEncryptor() bool
	Encryptor() encryptors.Encryptor
	IsSigner() bool
	Signer() signers.Signer
	IsProfile() bool
	Profile() profiles.Profile
	IsConnections() bool
	Connections() connections.Connections
	IsShares() bool
	Shares() shares.Shares
}
