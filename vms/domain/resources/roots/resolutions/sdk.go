package resolutions

// Builder represents a resolution builder
type Builder interface {
	Create() Builder
	WithFees(fees uint16) Builder
	WithAffiliate(affiliate uint16) Builder
	Now() (Resolution, error)
}

// Resolution represents a blockchain resolution
type Resolution interface {
	Fees() uint16
	Affiliate() uint16
}
