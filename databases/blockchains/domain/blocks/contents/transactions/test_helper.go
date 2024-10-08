package transactions

import (
	"crypto/ed25519"

	"github.com/steve-care-software/steve/databases/blockchains/domain/blocks/contents/transactions/entries"
)

// NewTransactionsForTests creates a new transactions for tests
func NewTransactionsForTests(list []Transaction) Transactions {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewTransactionForTests creates a new transaction for tests
func NewTransactionForTests(entry entries.Entry, signature []byte, pubKey ed25519.PublicKey) Transaction {
	ins, err := NewTransactionBuilder().Create().
		WithEntry(entry).
		WithSignature(signature).
		WithPublicKey(pubKey).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
