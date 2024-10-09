package params

import "github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds"

// Params represents params
type Params interface {
	Kind() kinds.Kind
	Internal() string
	External() string
	IsMandatory() bool
}
