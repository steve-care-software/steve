package applications

import (
	"bytes"
	"encoding/gob"

	"github.com/steve-care-software/steve/graphs/domain/responses"
	"github.com/steve-care-software/steve/graphs/domain/scripts"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references/externals"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/suites"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/suites/expectations"
	"github.com/steve-care-software/steve/hash"
	application_lists "github.com/steve-care-software/steve/lists/applications"
	application_resources "github.com/steve-care-software/steve/resources/applications"
)

type application struct {
	storeListApp application_lists.Application
	resourceApp  application_resources.Application
	hashAdapter  hash.Adapter
}

func createApplication(
	storeListApp application_lists.Application,
	resourceApp application_resources.Application,
	hashAdapter hash.Adapter,
) Application {
	out := application{
		storeListApp: storeListApp,
		resourceApp:  resourceApp,
		hashAdapter:  hashAdapter,
	}

	return &out
}

// Execute executes a script on the database
func (app *application) Execute(script scripts.Script) (responses.Response, error) {
	return nil, nil
}

func (app *application) saveSchema(schema schemas.Schema) error {
	return nil
}

func (app *application) saveHead(head heads.Head) error {
	return nil
}

func (app *application) saveConnections(connections connections.Connections) error {
	return nil
}

func (app *application) saveConnection(connection connections.Connection) error {
	return nil
}

func (app *application) saveLinks(links links.Links) error {
	return nil
}

func (app *application) saveLink(link links.Link) error {
	return nil
}

func (app *application) saveExternal(externalIns externals.External) (*hash.Hash, error) {
	ins := external{
		schema: externalIns.Schema(),
		point:  externalIns.Point(),
	}

	data, pHash, err := app.toBytes(ins)
	if err != nil {
		return nil, err
	}

	keyname := pHash.String()
	_, err = app.resourceApp.Retrieve(keyname)
	if err == nil {
		return pHash, nil
	}

	err = app.resourceApp.Insert(keyname, data)
	if err != nil {
		return nil, err
	}

	return pHash, nil
}

func (app *application) saveSuites(saves suites.Suites) error {
	return nil
}

func (app *application) saveSuite(save suites.Suite) error {
	return nil
}

func (app *application) saveExpectations(expectations expectations.Expectations) error {
	return nil
}

func (app *application) saveExpectation(expectation expectations.Expectation) error {
	return nil
}

func (app *application) saveReferences(references references.References) error {
	return nil
}

func (app *application) saveReference(reference references.Reference) error {
	return nil
}

func (app *application) toBytes(str any) ([]byte, *hash.Hash, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(str)
	if err != nil {
		return nil, nil, err
	}

	data := buf.Bytes()
	pHash, err := app.hashAdapter.FromBytes(data)
	if err != nil {
		return nil, nil, err
	}

	return data, pHash, nil

}
