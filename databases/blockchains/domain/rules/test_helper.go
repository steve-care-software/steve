package rules

// NewRulesForTests creates a new rules for tests
func NewRulesForTests(miningValue uint8, baseDiff uint8, incrDiffPerTrx float64) Rules {
	ins, err := NewBuilder().Create().
		WithMiningValue(miningValue).
		WithBaseDifficulty(baseDiff).
		WithIncreaseDifficultyPerTrx(incrDiffPerTrx).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
