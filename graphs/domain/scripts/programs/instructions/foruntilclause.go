package instructions

type forUntilClause struct {
	name  string
	value uint
}

func createForUntilClause(name string, value uint) ForUntilClause {
	return &forUntilClause{
		name:  name,
		value: value,
	}
}

// Name returns the name of the clause
func (obj *forUntilClause) Name() string {
	return obj.name
}

// Value returns the value of the clause
func (obj *forUntilClause) Value() uint {
	return obj.value
}
