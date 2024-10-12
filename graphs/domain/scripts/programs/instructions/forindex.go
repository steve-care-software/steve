package instructions

type forIndex struct {
	clause       ForUntilClause
	instructions ForInstructions
}

func createForIndex(
	clause ForUntilClause,
	instructions ForInstructions,
) ForIndex {
	out := forIndex{
		clause:       clause,
		instructions: instructions,
	}

	return &out
}

// Clause returns the clause
func (obj *forIndex) Clause() ForUntilClause {
	return obj.clause
}

// Instructions returns the instructions
func (obj *forIndex) Instructions() ForInstructions {
	return obj.instructions
}
