package elements

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/programs/grammars/rules"
	"github.com/steve-care-software/steve/domain/scripts/specifics/grammars/blocks/lines/tokens/elements/references"
)

// Element represents an element
type Element interface {
	Hash() hash.Hash
	IsReference() bool
	Reference() references.Reference
	IsRule() bool
	Rule() rules.Rule
	IsConstant() bool
	Constant() string
}
