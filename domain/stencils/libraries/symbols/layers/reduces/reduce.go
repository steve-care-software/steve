package reduces

type reduce struct {
	variable string
	length   uint8
}

func createReduce(
	variable string,
	length uint8,
) Reduce {
	out := reduce{
		variable: variable,
		length:   length,
	}

	return &out
}

// Variable returns the variable
func (obj *reduce) Variable() string {
	return obj.variable
}

// Length returns the length
func (obj *reduce) Length() uint8 {
	return obj.length
}
