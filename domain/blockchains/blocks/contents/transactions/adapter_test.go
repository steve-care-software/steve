package transactions

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/domain/hash"
)

func TestAdapter_single_Success(t *testing.T) {
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

	trx := NewTransactionForTests(
		entries.NewEntryForTests(*pFlag, *pScript, 34),
		[]byte("lets say this is a signature"),
	)

	adapter := NewAdapter()
	retBytes, err := adapter.InstanceToBytes(trx)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retTrx, retRemaining, err := adapter.BytesToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(retRemaining) > 0 {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(trx, retTrx) {
		t.Errorf("the returned transaction is invalid")
		return
	}
}

func TestAdapter_multiple_withRemaining_Success(t *testing.T) {
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

	trx := NewTransactionsForTests([]Transaction{
		NewTransactionForTests(
			entries.NewEntryForTests(*pFlag, *pScript, 22),
			[]byte("lets say this is a signature"),
		),
		NewTransactionForTests(
			entries.NewEntryForTests(*pFlag, *pOtherScript, 34),
			[]byte("lets say this is a signature"),
		),
	})

	adapter := NewAdapter()
	retBytes, err := adapter.InstancesToBytes(trx)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining")
	retTrx, retRemaining, err := adapter.BytesToInstances(append(retBytes, remaining...))
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
