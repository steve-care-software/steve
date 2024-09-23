package targets

type target struct {
	constant string
	rule     string
}

func createTargetWithConstant(
	constant string,
) Target {
	return createTargetInternally(constant, "")
}

func createTargetWithRule(
	rule string,
) Target {
	return createTargetInternally("", rule)
}

func createTargetInternally(
	constant string,
	rule string,
) Target {
	out := target{
		constant: constant,
		rule:     rule,
	}

	return &out
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *target) IsConstant() bool {
	return obj.constant != ""
}

// Constant returns the constant, if any
func (obj *target) Constant() string {
	return obj.constant
}

// IsRule returns true if there is a rule, false otherwise
func (obj *target) IsRule() bool {
	return obj.rule != ""
}

// Rule returns the rule, if any
func (obj *target) Rule() string {
	return obj.rule
}
