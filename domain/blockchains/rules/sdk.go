package rules

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the rules adapter
type Adapter interface {
	ToBytes(ins Rules) ([]byte, error)
	ToInstance(data []byte) (Rules, error)
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithMiningValue(miningValue uint8) Builder
	WithBaseDifficulty(baseDifficulty uint8) Builder
	WithIncreaseDifficultyPerTrx(incrDiffPerTrx float64) Builder
	Now() (Rules, error)
}

// Rules represents the blockchain rules
type Rules interface {
	MiningValue() uint8
	BaseDifficulty() uint8
	IncreaseDifficultyPerTrx() float64
}
