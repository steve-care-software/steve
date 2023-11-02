package identities

// Identities represents identities
type Identities interface {
	List() []Identity
	Fetch(index uint) (Identity, error)
	Delete(index uint) (Identities, error)
	Amount() uint
	Exceeds(amount uint) bool
}

// Identity represents an identity
type Identity interface {
	Name() string
	Container() string
}
