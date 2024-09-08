package activities

import (
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources/pointers"
)

type activity struct {
	hash    hash.Hash
	commits commits.Commits
	head    hash.Hash
}

func createActivity(
	hash hash.Hash,
	commits commits.Commits,
	head hash.Hash,
) Activity {
	out := activity{
		hash:    hash,
		commits: commits,
		head:    head,
	}

	return &out
}

// Hash returns the hash
func (obj *activity) Hash() hash.Hash {
	return obj.hash
}

// Commits returns the commits
func (obj *activity) Commits() commits.Commits {
	return obj.commits
}

// Head returns the head
func (obj *activity) Head() hash.Hash {
	return obj.head
}

// Map returns the map
func (obj *activity) Map() (map[string]pointers.Pointer, []string, error) {
	return obj.fetchRecursively(obj.head, obj.commits, []string{})
}

func (obj *activity) fetchRecursively(head hash.Hash, commits commits.Commits, deleted []string) (map[string]pointers.Pointer, []string, error) {
	retCommit, err := commits.Fetch(head)
	if err != nil {
		return nil, nil, err
	}

	commitMp, commitDeleted := retCommit.Modifications().Map()
	deleted = append(deleted, commitDeleted...)
	if retCommit.HasParent() {
		parent := retCommit.Parent()
		retMap, retDeleted, err := obj.fetchRecursively(parent, commits, deleted)
		if err != nil {
			return nil, nil, err
		}

		for keyname, onePointer := range retMap {
			if _, ok := commitMp[keyname]; ok {
				continue
			}

			isDeleted := false
			for _, oneDeleted := range deleted {
				if keyname == oneDeleted {
					isDeleted = true
					break
				}
			}

			if isDeleted {
				continue
			}

			commitMp[keyname] = onePointer
		}

		deleted = append(deleted, retDeleted...)
	}

	return commitMp, deleted, nil
}
