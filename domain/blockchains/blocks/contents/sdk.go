package contents

import (
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/hash"
)

// Content represents a block content
type Content interface {
	Hash() hash.Hash
	Database() hash.Hash
	Transactions() transactions.Transactions
	Parent() hash.Hash
}
