package blockchains

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/domain/hash"
)

func TestMine_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pFlag, err := hashAdapter.FromBytes([]byte("flag"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pScript, err := hashAdapter.FromBytes([]byte("script"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pOtherScript, err := hashAdapter.FromBytes([]byte("other script"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	trx := transactions.NewTransactionsForTests([]transactions.Transaction{
		transactions.NewTransactionForTests(
			entries.NewEntryForTests(*pFlag, *pScript, 22),
			[]byte("lets say this is a signature"),
		),
		transactions.NewTransactionForTests(
			entries.NewEntryForTests(*pFlag, *pOtherScript, 34),
			[]byte("lets say this is a signature"),
		),
	})

	requestedDifficulty := uint8(3)
	miningValue := uint8(0)
	result, err := mine(hashAdapter, trx, requestedDifficulty, miningValue)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pHash, err := hashAdapter.FromMultiBytes([][]byte{
		trx.Hash().Bytes(),
		result,
	})

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	prefix := []byte{
		miningValue,
		miningValue,
		miningValue,
	}

	if !bytes.HasPrefix(pHash.Bytes(), prefix) {
		t.Errorf("the returned result is invalid")
		return
	}

}
