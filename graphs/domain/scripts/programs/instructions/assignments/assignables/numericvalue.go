package assignables

type numericValue struct {
	flValue  *float64
	uiValue  *uint64
	intValue *int64
}

func createNumericValueWithFloat(
	flValue *float64,
) NumericValue {
	return createNumericValueInternally(flValue, nil, nil)
}

func createNumericValueWithUint(
	uiValue *uint64,
) NumericValue {
	return createNumericValueInternally(nil, uiValue, nil)
}

func createNumericValueWithInt(
	intValue *int64,
) NumericValue {
	return createNumericValueInternally(nil, nil, intValue)
}

func createNumericValueInternally(
	flValue *float64,
	uiValue *uint64,
	intValue *int64,
) NumericValue {
	return &numericValue{
		flValue:  flValue,
		uiValue:  uiValue,
		intValue: intValue,
	}
}

// IsFloat returns true if the value is a float
func (obj *numericValue) IsFloat() bool {
	return obj.flValue != nil
}

// Float returns the float value
func (obj *numericValue) Float() float64 {
	if obj.flValue != nil {
		return *obj.flValue
	}
	return 0
}

// IsUint returns true if the value is an unsigned integer
func (obj *numericValue) IsUint() bool {
	return obj.uiValue != nil
}

// Uint returns the unsigned integer value
func (obj *numericValue) Uint() *uint64 {
	return obj.uiValue
}

// IsInt returns true if the value is an integer
func (obj *numericValue) IsInt() bool {
	return obj.intValue != nil
}

// Int returns the integer value
func (obj *numericValue) Int() *int64 {
	return obj.intValue
}
