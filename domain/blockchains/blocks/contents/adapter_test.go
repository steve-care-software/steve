package contents

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/domain/hash"
)

func TestAdapter_withRemaining_Success(t *testing.T) {
	hashAdapter := hash.NewAdapter()
	pParent, err := hashAdapter.FromBytes([]byte("parent"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

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

	pMiner, err := hashAdapter.FromBytes([]byte("miner hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pCommit, err := hashAdapter.FromBytes([]byte("commit hash"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	trx := NewContentForTests(
		transactions.NewTransactionsForTests([]transactions.Transaction{
			transactions.NewTransactionForTests(
				entries.NewEntryForTests(*pFlag, *pScript, 22),
				[]byte("lets say this is a signature"),
			),
			transactions.NewTransactionForTests(
				entries.NewEntryForTests(*pFlag, *pOtherScript, 34),
				[]byte("lets say this is a signature"),
			),
		}),
		*pParent,
		*pMiner,
		*pCommit,
	)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(trx)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining data")
	retTrx, retRemaining, err := adapter.ToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(trx, retTrx) {
		t.Errorf("the returned transaction is invalid")
		return
	}
}
