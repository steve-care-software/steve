package identities

import (
	"github.com/steve-care-software/steve/domain/commands/executions/identities/connections"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/encryptors"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/identities"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/profiles"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/shares"
	"github.com/steve-care-software/steve/domain/commands/executions/identities/signers"
	"github.com/steve-care-software/steve/domain/commands/executions/shares/dashboards"
)

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithIdentity(identity identities.Identity) Builder
	WithDashboard(dashboard dashboards.Dashboard) Builder
	WithConnections(connections connections.Connections) Builder
	WithProfile(profile profiles.Profile) Builder
	WithShares(shares shares.Shares) Builder
	WithEncryptor(encryptor encryptors.Encryptor) Builder
	WithSigner(signer signers.Signer) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	IsIdentity() bool
	Identity() identities.Identity
	IsDashboard() bool
	Dashboard() dashboards.Dashboard
	IsConnections() bool
	Connections() connections.Connections
	IsProfile() bool
	Profile() profiles.Profile
	IsShares() bool
	Shares() shares.Shares
	IsEncryptor() bool
	Encryptor() encryptors.Encryptor
	IsSigner() bool
	Signer() signers.Signer
}
