package rules

// Rules represents the blockchain rules
type Rules interface {
	MiningValue() uint8
	BaseDifficulty() uint8
	IncreaseDifficultyPerTrx() float64
}
