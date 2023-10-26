package inputs

// Builder represents an input builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithBytes(bytes []byte) Builder
	Now() (Input, error)
}

// Input represents an input
type Input interface {
	Variable() string
	Bytes() []byte
}
