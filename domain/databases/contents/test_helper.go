package contents

import "github.com/steve-care-software/steve/domain/hash"

// NewContentsForTests creates a new contents for tests
func NewContentsForTests(kind uint, data [][]byte) Contents {
	list := []Content{}
	for _, oneData := range data {
		ins := NewContentForTests(kind, oneData)
		list = append(list, ins)
	}

	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentForTests creates a new content for tests
func NewContentForTests(kind uint, data []byte) Content {
	pHash, err := hash.NewAdapter().FromBytes(data)
	if err != nil {
		panic(err)
	}

	ins, err := NewContentBuilder().Create().WithHash(*pHash).WithData(data).WithKind(kind).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
