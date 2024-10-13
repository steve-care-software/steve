package primitives

type primitive struct {
	flag uint8
}

func createPrimitive(
	flag uint8,
) Primitive {
	out := primitive{
		flag: flag,
	}

	return &out
}

// Flag returns the flag
func (obj *primitive) Flag() uint8 {
	return obj.flag
}

// IsNumeric returns true if it is numeric, false otherwise
func (obj *primitive) IsNumeric() bool {
	return obj.flag <= FlagFloat
}
