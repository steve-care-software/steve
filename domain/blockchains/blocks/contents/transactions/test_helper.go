package transactions

import "github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions/entries"

// NewTransactionsForTests creates a new transactions for tests
func NewTransactionsForTests(list []Transaction) Transactions {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewTransactionForTests creates a new transaction for tests
func NewTransactionForTests(entry entries.Entry, signature []byte) Transaction {
	ins, err := NewTransactionBuilder().Create().
		WithEntry(entry).
		WithSignature(signature).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
