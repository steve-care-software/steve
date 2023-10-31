package resolutions

// Resolution represents a stencil resolution
type Resolution interface {
	Fees() uint16
	Affiliate() uint16
	Symbol() uint16
}
