package assignables

type operator struct {
	arithmetic *uint8
	relational *uint8
	equal      *uint8
	logical    *uint8
}

func createOperatorWithArithmetic(
	arithmetic *uint8,
) Operator {
	return createOperatorInternally(arithmetic, nil, nil, nil)
}

func createOperatorWithRelational(
	relational *uint8,
) Operator {
	return createOperatorInternally(nil, relational, nil, nil)
}

func createOperatorWithEqual(
	equal *uint8,
) Operator {
	return createOperatorInternally(nil, nil, equal, nil)
}

func createOperatorWithLogical(
	logical *uint8,
) Operator {
	return createOperatorInternally(nil, nil, nil, logical)
}

func createOperatorInternally(
	arithmetic *uint8,
	relational *uint8,
	equal *uint8,
	logical *uint8,
) Operator {
	return &operator{
		arithmetic: arithmetic,
		relational: relational,
		equal:      equal,
		logical:    logical,
	}
}

// IsArithmetic returns true if an arithmetic operator is set
func (obj *operator) IsArithmetic() bool {
	return obj.arithmetic != nil
}

// Arithmetic returns the arithmetic operator if present
func (obj *operator) Arithmetic() *uint8 {
	return obj.arithmetic
}

// IsRelational returns true if a relational operator is set
func (obj *operator) IsRelational() bool {
	return obj.relational != nil
}

// Relational returns the relational operator if present
func (obj *operator) Relational() *uint8 {
	return obj.relational
}

// IsEqual returns true if an equal operator is set
func (obj *operator) IsEqual() bool {
	return obj.equal != nil
}

// Equal returns the equal operator if present
func (obj *operator) Equal() *uint8 {
	return obj.equal
}

// IsLogical returns true if a logical operator is set
func (obj *operator) IsLogical() bool {
	return obj.logical != nil
}

// Logical returns the logical operator if present
func (obj *operator) Logical() *uint8 {
	return obj.logical
}
