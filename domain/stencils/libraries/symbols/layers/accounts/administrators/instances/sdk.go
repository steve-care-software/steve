package instances

// Instance represents the instance
type Instance interface {
	Variable() string
	Method() Method
}

// Method represents the instance method
type Method interface {
	IsAdministrator() bool
	Administrator() Administrator
	IsIdentities() bool
	Identities() Identities
	IsIdentity() bool
	Identity() Identity
}

// Administrator represents the admin's method
type Administrator interface {
	IsHasIdentities() bool
	IsGetIdentities() bool
}

// Identities represents the identity's metod
type Identities interface {
	IsGetLength() bool
	IsFetchAtIndex() bool
	FetchAtIndex() uint
}

// Identity represents the identity's method
type Identity interface {
	IsGetName() bool
	IsGetContainer() bool
}
