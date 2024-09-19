package elements

type element struct {
	token string
	rule  string
}

func createElementWithToken(
	token string,
) Element {
	return createElementInternally(token, "")
}

func createElementWithRule(
	rule string,
) Element {
	return createElementInternally("", rule)
}

func createElementInternally(
	token string,
	rule string,
) Element {
	out := element{
		token: token,
		rule:  rule,
	}

	return &out
}

// IsToken returns true if there is a token, false otherwise
func (obj *element) IsToken() bool {
	return obj.token != ""
}

// Token returns the token, if any
func (obj *element) Token() string {
	return obj.token
}

// IsRule returns true if there is a rule, false otherwise
func (obj *element) IsRule() bool {
	return obj.rule != ""
}

// Rule returns the rule, if any
func (obj *element) Rule() string {
	return obj.rule
}
