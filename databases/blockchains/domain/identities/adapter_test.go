package identities

import (
	"bytes"
	"crypto/ed25519"
	"reflect"
	"testing"

	"github.com/steve-care-software/steve/commons/hash"
)

func TestAdapter_withRemaining_Success(t *testing.T) {
	_, pk, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	identity := NewIdentityForTests("myName", pk)

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(identity)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining bytes")
	retIdentity, retRemaining, err := adapter.ToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(identity, retIdentity) {
		t.Errorf("the returned identity is invalid")
		return
	}
}

func TestAdapter_withFlags_withRemaining_Success(t *testing.T) {
	_, pk, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	hashAdapter := hash.NewAdapter()
	pFirstFlag, err := hashAdapter.FromBytes([]byte("this is first"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	pSecondFlag, err := hashAdapter.FromBytes([]byte("this is second"))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	identity := NewIdentityWithFlagsForTests("myName", pk, []hash.Hash{
		*pFirstFlag,
		*pSecondFlag,
	})

	adapter := NewAdapter()
	retBytes, err := adapter.ToBytes(identity)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	remaining := []byte("this is some remaining bytes")
	retIdentity, retRemaining, err := adapter.ToInstance(append(retBytes, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	if !reflect.DeepEqual(identity, retIdentity) {
		t.Errorf("the returned identity is invalid")
		return
	}
}
