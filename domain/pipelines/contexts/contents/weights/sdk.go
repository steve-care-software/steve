package weights

// Weights represents weights
type Weights interface {
	List() []Weight
}

// Weight represents a weight
type Weight interface {
	Name() string
	Value() uint
	HasReverse() bool
	Reverse() string
}
