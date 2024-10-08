package applications

import (
	"bytes"
	"encoding/gob"

	"github.com/steve-care-software/steve/graphs/domain/responses"
	"github.com/steve-care-software/steve/graphs/domain/scripts"
	"github.com/steve-care-software/steve/graphs/domain/scripts/commons/heads"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/headers"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/headers/names"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/headers/names/cardinalities"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links"
	"github.com/steve-care-software/steve/graphs/domain/scripts/schemas/connections/links/references"
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
	dbIdentifier string
}

func createApplication(
	storeListApp application_lists.Application,
	resourceApp application_resources.Application,
	hashAdapter hash.Adapter,
	dbIdentifier string,
) Application {
	out := application{
		storeListApp: storeListApp,
		resourceApp:  resourceApp,
		hashAdapter:  hashAdapter,
		dbIdentifier: dbIdentifier,
	}

	return &out
}

// Execute executes a script on the database
func (app *application) Execute(script scripts.Script) (responses.Response, error) {
	err := app.resourceApp.Init(app.dbIdentifier)
	if err != nil {
		return nil, err
	}

	if script.IsSchema() {
		schema := script.Schema()
		err = app.saveSchema(schema)
		if err != nil {
			return nil, err
		}
	}

	err = app.resourceApp.Commit()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (app *application) saveSchema(schema schemas.Schema) error {
	return nil
}

func (app *application) saveHead(head heads.Head) error {
	return nil
}

func (app *application) saveConnections(connectionsIns connections.Connections) ([]hash.Hash, error) {
	hashes := []hash.Hash{}
	list := connectionsIns.List()
	for _, oneConnection := range list {
		pHash, err := app.saveConnection(oneConnection)
		if err != nil {
			return nil, err
		}

		hashes = append(hashes, *pHash)
	}

	return hashes, nil
}

func (app *application) saveConnection(connectionIns connections.Connection) (*hash.Hash, error) {
	links := connectionIns.Links()
	retLinkHashes, err := app.saveLinks(links)
	if err != nil {
		return nil, err
	}

	linksBytes := [][]byte{}
	for _, oneHash := range retLinkHashes {
		linksBytes = append(linksBytes, oneHash.Bytes())
	}

	suites := connectionIns.Suites()
	retSuiteHashes, err := app.saveSuites(suites)
	if err != nil {
		return nil, err
	}

	suitesBytes := [][]byte{}
	for _, oneHash := range retSuiteHashes {
		suitesBytes = append(suitesBytes, oneHash.Bytes())
	}

	header := connectionIns.Header()
	ins := connection{
		header: app.saveConnectionHeader(header),
		links:  linksBytes,
		suites: suitesBytes,
	}

	return app.retrieveOrSave(ins)
}

func (app *application) saveConnectionHeader(header headers.Header) connectionHeader {
	name := header.Name()
	ins := connectionHeader{
		name: app.saveConnectionHeaderName(name),
	}

	if header.HasReverse() {
		reverseIns := header.Reverse()
		reverse := app.saveConnectionHeaderName(reverseIns)
		ins.pReverse = &reverse
	}

	return ins
}

func (app *application) saveConnectionHeaderName(name names.Name) connectionName {
	cardinality := name.Cardinality()
	return connectionName{
		name:        name.Name(),
		cardinality: app.saveConnectionCardinality(cardinality),
	}
}

func (app *application) saveConnectionCardinality(cardinality cardinalities.Cardinality) connectionCardinality {
	ins := connectionCardinality{
		min: cardinality.Min(),
	}

	if cardinality.HaxMax() {
		pMax := cardinality.Max()
		ins.pMax = pMax
	}

	return ins
}

func (app *application) saveLinks(linksIns links.Links) ([]hash.Hash, error) {
	hashes := []hash.Hash{}
	list := linksIns.List()
	for _, oneLink := range list {
		pHash, err := app.saveLink(oneLink)
		if err != nil {
			return nil, err
		}

		hashes = append(hashes, *pHash)
	}

	return hashes, nil
}

func (app *application) saveLink(linkIns links.Link) (*hash.Hash, error) {
	origin := linkIns.Origin()
	pOriginHash, err := app.saveReference(origin)
	if err != nil {
		return nil, err
	}

	target := linkIns.Target()
	pTargetHash, err := app.saveReference(target)
	if err != nil {
		return nil, err
	}

	ins := link{
		origin: pOriginHash.Bytes(),
		target: pTargetHash.Bytes(),
	}

	return app.retrieveOrSave(ins)
}

func (app *application) saveSuites(suitesIns suites.Suites) ([]hash.Hash, error) {
	hashes := []hash.Hash{}
	list := suitesIns.List()
	for _, oneSuite := range list {
		pHash, err := app.saveSuite(oneSuite)
		if err != nil {
			return nil, err
		}

		hashes = append(hashes, *pHash)
	}

	return hashes, nil
}

func (app *application) saveSuite(suiteIns suites.Suite) (*hash.Hash, error) {
	expectations := suiteIns.Expectations()
	retExpectations, err := app.saveSuiteExpectations(expectations)
	if err != nil {
		return nil, err
	}

	link := suiteIns.Link()
	pLinkHash, err := app.saveLink(link)
	if err != nil {
		return nil, err
	}

	ins := suite{
		name:         suiteIns.Name(),
		link:         pLinkHash.Bytes(),
		expectations: retExpectations,
	}

	return app.retrieveOrSave(ins)
}

func (app *application) saveSuiteExpectations(expectations expectations.Expectations) ([]suiteExpectation, error) {
	output := []suiteExpectation{}
	list := expectations.List()
	for _, oneExpectation := range list {
		pExpectation, err := app.saveSuiteExpectation(oneExpectation)
		if err != nil {
			return nil, err
		}

		output = append(output, *pExpectation)
	}

	return output, nil
}

func (app *application) saveSuiteExpectation(expectationIns expectations.Expectation) (*suiteExpectation, error) {
	references := expectationIns.References()
	retHashes, err := app.saveReferences(references)
	if err != nil {
		return nil, err
	}

	referencesBytes := [][]byte{}
	for _, oneHash := range retHashes {
		referencesBytes = append(referencesBytes, oneHash.Bytes())
	}

	return &suiteExpectation{
		references: referencesBytes,
		isFail:     expectationIns.IsFail(),
	}, nil
}

func (app *application) saveReferences(references references.References) ([]hash.Hash, error) {
	hashes := []hash.Hash{}
	list := references.List()
	for _, oneReference := range list {
		pHash, err := app.saveReference(oneReference)
		if err != nil {
			return nil, err
		}

		hashes = append(hashes, *pHash)
	}

	return hashes, nil
}

func (app *application) saveReference(referenceIns references.Reference) (*hash.Hash, error) {
	ins := reference{}
	if referenceIns.IsInternal() {
		ins.internal = referenceIns.Internal()
	}

	if referenceIns.IsExternal() {
		externalIns := referenceIns.External()
		ins.pExternal = &external{
			schema: externalIns.Schema(),
			point:  externalIns.Point(),
		}
	}

	return app.retrieveOrSave(ins)
}

func (app *application) retrieveOrSave(value any) (*hash.Hash, error) {
	data, pHash, err := app.toBytes(value)
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
