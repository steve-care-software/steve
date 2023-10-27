package resolutions

// Resolution represents a stencil resolution
type Resolution interface {
	Fees() uint16
	Engine() uint16
	Affiliates() uint16
	Symbol() uint16
}
