package identities

// Identities represents identities
type Identities interface {
	List() []Identity
	Delete(index uint) (Identities, error)
}

// Identity represents an identity
type Identity interface {
	Name() string
	Container() string
}
