package values

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/encryptors"
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles"
	"github.com/steve-care-software/steve/domain/accounts/identities/shares"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers"
	"github.com/steve-care-software/steve/domain/dashboards"
)

// Builder represents a value builder
type Builder interface {
	Create() Builder
	WithDashboard(dashboard dashboards.Dashboard) Builder
	WithEncryptor(encryptor encryptors.Encryptor) Builder
	WithSigner(signer signers.Signer) Builder
	WithProfile(profile profiles.Profile) Builder
	WithHasConnections(hasConnections bool) Builder
	WithConnections(connections profiles.Connections) Builder
	WithHasShares(hasShares bool) Builder
	WithShares(shares shares.Shares) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsDashboard() bool
	Dashboard() dashboards.Dashboard
	IsEncryptor() bool
	Encryptor() encryptors.Encryptor
	IsSigner() bool
	Signer() signers.Signer
	IsProfile() bool
	Profile() profiles.Profile
	IsHasConnections() bool
	HasConnections() *bool
	IsConnections() bool
	Connections() profiles.Connections
	IsHasShares() bool
	HasShares() *bool
	IsShares() bool
	Shares() shares.Shares
}
