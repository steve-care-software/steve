package rules

type rules struct {
	miningValue    uint8
	baseDifficulty uint8
	incrDiffPerTrx float64
}

func createRules(
	miningValue uint8,
	baseDifficulty uint8,
	incrDiffPerTrx float64,
) Rules {
	out := rules{
		miningValue:    miningValue,
		baseDifficulty: baseDifficulty,
		incrDiffPerTrx: incrDiffPerTrx,
	}

	return &out
}

// MiningValue return the miningValue
func (obj *rules) MiningValue() uint8 {
	return obj.miningValue
}

// BaseDifficulty return the baseDifficulty
func (obj *rules) BaseDifficulty() uint8 {
	return obj.baseDifficulty
}

// IncreaseDifficultyPerTrx return the increase the difficulty per trx
func (obj *rules) IncreaseDifficultyPerTrx() float64 {
	return obj.incrDiffPerTrx
}
