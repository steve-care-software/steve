package applications

import (
	"os"
	"testing"

	"github.com/steve-care-software/steve/graphs/domain/scripts"
	lists "github.com/steve-care-software/steve/lists/applications"
	resources "github.com/steve-care-software/steve/resources/applications"
)

func TestApplication_withSchema_Success(t *testing.T) {
	scriptCode := []byte(`
		head:
			engine: v1;
			name: mySchema;
			access: 
				read: .first .second (0.2);
				write: 
					.first .again;
					review: .first .second .third (0.1);
				;
			;
		;

		son;
		father;
		grandFather;
		grandGrandFather;

		father[0,3](son[1,]): .son .father
						| .father .grandFather
						| .grandFather .grandGrandFather
						---
							mySuite[ .mySchema[son] .grandGrandFather]:
								!(.son .father .grandFather .grandGrandFather);
								(.son .father .grandFather .grandGrandFather);
							;
						;

		grandFather(grandSon[2,]): .son .grandFather
								| .father .grandGrandFather
								;
	`)

	scriptAdapter, err := scripts.NewAdapterFactory().Create()
	if err != nil {
		t.Errorf("there was an error while running the grammar test suites: %s", err.Error())
		return
	}

	script, _, err := scriptAdapter.ToScript(scriptCode)
	if err != nil {
		t.Errorf("there was an error while running the grammar test suites: %s", err.Error())
		return
	}

	baseDir := "./test_files"
	defer func() {
		os.RemoveAll(baseDir)
	}()

	resourceApp, err := resources.NewBuilder().Create().
		WithBasePath(baseDir).
		WithReadChunkSize(1024).
		WithTargetIdentifier("target.tmp").
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	err = resourceApp.Init("my_database.db")
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	listApp, err := lists.NewBuilder().Create().
		WithResource(resourceApp).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	dbIdentifier := "my_database.db"
	application, err := NewBuilder(
		listApp,
		resourceApp,
	).Create().WithIdentifier(dbIdentifier).Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	_, err = application.Execute(script)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}
}
