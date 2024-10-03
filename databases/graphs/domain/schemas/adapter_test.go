package schemas

import (
	"errors"
	"testing"
)

func TestAdapter_Success(t *testing.T) {
	adapter, err := NewAdapterFactory().Create()
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	input := []byte(`
		v1;
		name: myName;

		son;
		father;
		grandFather;
		grandGrandFather;

		father[0,3](son+): .son .father
							| .father .grandFather
							| .grandFather .grandGrandFather
							---
								mySuite[.son .grandGrandFather]:
									(.son .father .grandFather .grandGrandFather);
									!(.son .father .grandFather .grandGrandFather);
								;
							;

		grandFather(grandSon[2,]): .son .grandFather
								| .father .grandGrandFather
								;
	`)

	_, _, err = adapter.ToSchema(input)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	panic(errors.New("stop"))
}
