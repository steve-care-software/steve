package formats

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/contexts/contents/suites"
)

type format struct {
	hash    hash.Hash
	point   string
	grammar string
	suites  suites.Suites
}

func createFormat(
	hash hash.Hash,
	point string,
	grammar string,
) Format {
	return createFormatInternally(hash, point, grammar, nil)
}

func createFormatWithSuites(
	hash hash.Hash,
	point string,
	grammar string,
	suites suites.Suites,
) Format {
	return createFormatInternally(hash, point, grammar, suites)
}

func createFormatInternally(
	hash hash.Hash,
	point string,
	grammar string,
	suites suites.Suites,
) Format {
	out := format{
		hash:    hash,
		point:   point,
		grammar: grammar,
		suites:  suites,
	}

	return &out
}

// Hash returns the hash
func (obj *format) Hash() hash.Hash {
	return obj.hash
}

// Point returns the point
func (obj *format) Point() string {
	return obj.point
}

// Grammar returns the grammar
func (obj *format) Grammar() string {
	return obj.grammar
}

// HasSuites returns true if there is suites, false otherwise
func (obj *format) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *format) Suites() suites.Suites {
	return obj.suites
}
