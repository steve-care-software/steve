package roots

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/scripts/components/heads"
	"github.com/steve-care-software/steve/domain/scripts/specifics/pipelines/executions/suites"
)

type root struct {
	hash      hash.Hash
	head      heads.Head
	pipelines []string
	suites    suites.Suites
}

func createRoot(
	hash hash.Hash,
	head heads.Head,
	pipelines []string,
) Root {
	return createRootInternally(
		hash,
		head,
		pipelines,
		nil,
	)
}

func createRootWithSuites(
	hash hash.Hash,
	head heads.Head,
	pipelines []string,
	suites suites.Suites,
) Root {
	return createRootInternally(
		hash,
		head,
		pipelines,
		suites,
	)
}

func createRootInternally(
	hash hash.Hash,
	head heads.Head,
	pipelines []string,
	suites suites.Suites,
) Root {
	out := root{
		hash:      hash,
		head:      head,
		pipelines: pipelines,
		suites:    suites,
	}

	return &out
}

// Hash returns the hash
func (obj *root) Hash() hash.Hash {
	return obj.hash
}

// Head returns the head
func (obj *root) Head() heads.Head {
	return obj.head
}

// Pipelines returns the pipelines
func (obj *root) Pipelines() []string {
	return obj.pipelines
}

// HasSuites returns true if there is suites, false otherwise
func (obj *root) HasSuites() bool {
	return obj.suites != nil
}

// Suites returns the suites, if any
func (obj *root) Suites() suites.Suites {
	return obj.suites
}
