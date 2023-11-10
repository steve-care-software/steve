package resolutions

// Resolution represents a blockchain resolution
type Resolution interface {
	Fees() uint16
	Affiliate() uint16
	Symbol() uint16
}
