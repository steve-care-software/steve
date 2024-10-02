package references

import "github.com/steve-care-software/steve/databases/graphs/domain/languages/connections/links/references/externals"

// References represents references
type References interface {
	List() []Reference
}

type Reference interface {
	IsInternal() bool
	Internal() string
	IsExternal() bool
	External() externals.External
}
