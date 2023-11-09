package identities

import (
	"github.com/steve-care-software/steve/domain/accounts/identities/encryptors"
	"github.com/steve-care-software/steve/domain/accounts/identities/profiles"
	"github.com/steve-care-software/steve/domain/accounts/identities/shares"
	"github.com/steve-care-software/steve/domain/accounts/identities/signers"
	"github.com/steve-care-software/steve/domain/credentials"
	"github.com/steve-care-software/steve/domain/dashboards"
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

// Repository represents the identity repository
type Repository interface {
	List() ([]string, error)
	Retrieve(credentials credentials.Credentials) (Identity, error)
}

// Service represents the identity service
type Service interface {
	Insert(identity Identity, password []byte) error
	Save(identity Identity, password []byte, newPassword []byte) error
	Delete(credentials credentials.Credentials) error
}
