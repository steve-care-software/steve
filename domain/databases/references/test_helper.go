package references

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/trees"
)

// NewReferenceForTests creates a new reference with peers for tests
func NewReferenceForTests(maxPeerAmount uint) Reference {
	commits := NewCommitsForTests(32)
	ins, err := NewBuilder().Create().WithCommits(commits).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewReferenceWithContentKeysForTests creates a new reference with contentKeys for tests
func NewReferenceWithContentKeysForTests(maxPeerAmount uint) Reference {
	contentKeys, err := NewContentKeysBuilder().Create().WithList([]ContentKey{
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
		NewContentKeyForTests(),
	}).Now()
	if err != nil {
		panic(err)
	}

	commits := NewCommitsForTests(32)
	ins, err := NewBuilder().Create().WithContentKeys(contentKeys).WithCommits(commits).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommitsForTests creates a new commits for tests
func NewCommitsForTests(amount uint) Commits {
	list := []Commit{}
	for i := 0; i < int(amount); i++ {
		list = append(list, NewCommitForTests())
	}

	ins, err := NewCommitsBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommitForTests creates a new commit for tests
func NewCommitForTests() Commit {
	createdOn := time.Now().UTC()
	action := NewActionWithInsert()
	ins, err := NewCommitBuilder().Create().WithAction(action).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommitWithParentForTests creates a new commit with parent for tests
func NewCommitWithParentForTests() Commit {
	pParentHash, err := hash.NewAdapter().FromBytes([]byte("this is a parent hash"))
	if err != nil {
		panic(err)
	}

	createdOn := time.Now().UTC()
	action := NewActionWithDelete()
	ins, err := NewCommitBuilder().Create().WithAction(action).WithParent(*pParentHash).CreatedOn(createdOn).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewActionWithInsertAndDelete creates a new action with insert and delete
func NewActionWithInsertAndDelete() Action {
	insert, err := trees.NewBuilder().Create().WithBlocks([][]byte{
		[]byte("first"),
		[]byte("second"),
		[]byte("third"),
	}).Now()
	if err != nil {
		panic(err)
	}

	del, err := trees.NewBuilder().Create().WithBlocks([][]byte{
		[]byte("first del"),
		[]byte("second del"),
		[]byte("third del"),
	}).Now()
	if err != nil {
		panic(err)
	}

	ins, err := NewActionBuilder().Create().WithInsert(insert).WithDelete(del).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewActionWithInsert creates a new action with insert
func NewActionWithInsert() Action {
	values, err := trees.NewBuilder().Create().WithBlocks([][]byte{
		[]byte("first"),
		[]byte("second"),
		[]byte("third"),
	}).Now()
	if err != nil {
		panic(err)
	}

	ins, err := NewActionBuilder().Create().WithInsert(values).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewActionWithDelete creates a new action with delete
func NewActionWithDelete() Action {
	values, err := trees.NewBuilder().Create().WithBlocks([][]byte{
		[]byte("first"),
		[]byte("second"),
		[]byte("third"),
	}).Now()
	if err != nil {
		panic(err)
	}

	ins, err := NewActionBuilder().Create().WithDelete(values).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewPointerForTests creates a new pointer for tests
func NewPointerForTests() Pointer {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	from := uint(r1.Int() + 1)

	s1 = rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)

	length := uint(r1.Int() + 1)
	pointer, err := NewPointerBuilder().Create().From(from).WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	return pointer
}

// NewContentKeyForTests creates a new content key for tests
func NewContentKeyForTests() ContentKey {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	pHash, err := hash.NewAdapter().FromBytes([]byte(fmt.Sprintf("this is some data %d", r1.Int())))
	if err != nil {
		panic(err)
	}

	pCommitHash, err := hash.NewAdapter().FromBytes([]byte(fmt.Sprintf("this is some commit data %d", r1.Int())))
	if err != nil {
		panic(err)
	}

	from := uint(r1.Intn(233456))
	length := uint(r1.Intn(22323)) + 1
	pointer, err := NewPointerBuilder().Create().From(from).WithLength(length).Now()
	if err != nil {
		panic(err)
	}

	ins, err := NewContentKeyBuilder().Create().
		WithHash(*pHash).
		WithKind(43).
		WithContent(pointer).
		WithCommit(*pCommitHash).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
