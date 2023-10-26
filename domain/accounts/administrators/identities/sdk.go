package identities

// Identities represents identities
type Identities interface {
	List() []Identity
}

// Identity represents an identity
type Identity interface {
	Name() string
	Container() string
}
