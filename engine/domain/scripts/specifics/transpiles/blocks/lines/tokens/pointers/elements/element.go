package elements

import "github.com/steve-care-software/steve/hash"

type element struct {
	hash  hash.Hash
	token string
	rule  string
}

func createElementWithToken(
	hash hash.Hash,
	token string,
) Element {
	return createElementInternally(hash, token, "")
}

func createElementWithRule(
	hash hash.Hash,
	rule string,
) Element {
	return createElementInternally(hash, "", rule)
}

func createElementInternally(
	hash hash.Hash,
	token string,
	rule string,
) Element {
	out := element{
		hash:  hash,
		token: token,
		rule:  rule,
	}

	return &out
}

// Hash returns the hash
func (obj *element) Hash() hash.Hash {
	return obj.hash
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
