package transactions

import (
	"crypto/ed25519"
	"errors"

	"github.com/steve-care-software/steve/blockchains/domain/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/hash"
)

type transactionBuilder struct {
	hashAdapter hash.Adapter
	entry       entries.Entry
	signature   []byte
	pubKey      ed25519.PublicKey
}

func createTransactionBuilder(
	hashAdapter hash.Adapter,
) TransactionBuilder {
	out := transactionBuilder{
		hashAdapter: hashAdapter,
		entry:       nil,
		signature:   nil,
		pubKey:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *transactionBuilder) Create() TransactionBuilder {
	return createTransactionBuilder(
		app.hashAdapter,
	)
}

// WithEntry adds an entry to the builder
func (app *transactionBuilder) WithEntry(entry entries.Entry) TransactionBuilder {
	app.entry = entry
	return app
}

// WithSignature adds a signature to the builder
func (app *transactionBuilder) WithSignature(signature []byte) TransactionBuilder {
	app.signature = signature
	return app
}

// WithPublicKey adds a public key to the builder
func (app *transactionBuilder) WithPublicKey(pubKey ed25519.PublicKey) TransactionBuilder {
	app.pubKey = pubKey
	return app
}

// Now builds a new transaction instance
func (app *transactionBuilder) Now() (Transaction, error) {
	if app.entry == nil {
		return nil, errors.New("the entry is mandatory in ordder to build a Transaction instance")
	}

	if app.signature != nil && len(app.signature) <= 0 {
		app.signature = nil
	}

	if app.signature == nil {
		return nil, errors.New("the signature is mandatory in order to build a Transaction instance")
	}

	if app.pubKey == nil {
		return nil, errors.New("the public key is mandatory in order to build a Transaction instance")
	}

	if !ed25519.Verify(app.pubKey, app.entry.Hash().Bytes(), app.signature) {
		return nil, errors.New("the signature is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.entry.Hash().Bytes(),
		app.signature,
		app.pubKey,
	})

	if err != nil {
		return nil, err
	}

	return createTransaction(
		*pHash,
		app.entry,
		app.signature,
		app.pubKey,
	), nil
}
