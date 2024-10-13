package params

import "github.com/steve-care-software/steve/graphs/domain/scripts/commons/kinds"

type params struct {
	kind      kinds.Kind
	internal  string
	external  string
	mandatory bool
}

func createParams(
	kind kinds.Kind,
	internal string,
	external string,
	mandatory bool,
) Params {
	out := params{
		kind:      kind,
		internal:  internal,
		external:  external,
		mandatory: mandatory,
	}

	return &out
}

func (obj *params) Kind() kinds.Kind {
	return obj.kind
}

func (obj *params) Internal() string {
	return obj.internal
}

func (obj *params) External() string {
	return obj.external
}

func (obj *params) IsMandatory() bool {
	return obj.mandatory
}
