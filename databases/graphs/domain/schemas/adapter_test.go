package schemas

import (
	"bytes"
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
		name: mySchema;

		son;
		father;
		grandFather;
		grandGrandFather;

		father[1](son[0,]): .mySchema[son] .grandFather
							| .father .grandGrandFather
							;

		father[0,3](son+): .son .father
							| .father .grandFather
							| .grandFather .grandGrandFather
							---
								mySuite[.son .grandGrandFather]:
									(.son .father .grandFather .grandGrandFather);
									!(.son .father .grandFather .grandGrandFather);
								;
							;

		grandFather?(grandSon[2,]): .son .grandFather
								| .father .grandGrandFather
								;

		grandFather(grandSon*): .son .grandFather
								| .father .grandGrandFather
								;
	`)

	remaining := []byte("this is the remaining")

	retSchema, retRemaining, err := adapter.ToSchema(append(input, remaining...))
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(remaining, retRemaining) {
		t.Errorf("the returned remaining is invalid")
		return
	}

	header := retSchema.Header()
	if header.Version() != 1 {
		t.Errorf("the version was expected to be %d, %d returned", 1, header.Version())
		return
	}

	if header.Name() != "mySchema" {
		t.Errorf("the name was expected to be '%s', '%s' returned", "mySchema", header.Name())
		return
	}
}
