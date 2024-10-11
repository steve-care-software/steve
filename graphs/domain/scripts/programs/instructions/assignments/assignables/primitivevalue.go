package assignables

type primitiveValue struct {
	numeric   NumericValue
	boolValue *bool
	strValue  *string
}

func createPrimitiveValueWithNumeric(
	numeric NumericValue,
) PrimitiveValue {
	return createPrimitiveValueInternally(numeric, nil, nil)
}

func createPrimitiveValueWithBool(
	boolValue *bool,
) PrimitiveValue {
	return createPrimitiveValueInternally(nil, boolValue, nil)
}

func createPrimitiveValueWithString(
	strValue *string,
) PrimitiveValue {
	return createPrimitiveValueInternally(nil, nil, strValue)
}

func createPrimitiveValueInternally(
	numeric NumericValue,
	boolValue *bool,
	strValue *string,
) PrimitiveValue {
	return &primitiveValue{
		numeric:   numeric,
		boolValue: boolValue,
		strValue:  strValue,
	}
}

// IsNumeric returns true if the primitive value is numeric
func (obj *primitiveValue) IsNumeric() bool {
	return obj.numeric != nil
}

// Numeric returns the numeric value if present
func (obj *primitiveValue) Numeric() NumericValue {
	return obj.numeric
}

// IsBool returns true if the primitive value is boolean
func (obj *primitiveValue) IsBool() bool {
	return obj.boolValue != nil
}

// Bool returns the boolean value if present
func (obj *primitiveValue) Bool() *bool {
	return obj.boolValue
}

// IsString returns true if the primitive value is a string
func (obj *primitiveValue) IsString() bool {
	return obj.strValue != nil
}

// String returns the string value if present
func (obj *primitiveValue) String() *string {
	return obj.strValue
}
