package blockchains

import (
	"errors"

	"github.com/steve-care-software/steve/applications/cryptography"
	resources "github.com/steve-care-software/steve/applications/resources"
	"github.com/steve-care-software/steve/applications/resources/lists"
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/domain/blockchains/identities"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
	"github.com/steve-care-software/steve/domain/blockchains/rules"
	"github.com/steve-care-software/steve/domain/hash"
	"github.com/steve-care-software/steve/domain/uuids"
)

type builder struct {
	listApp                    lists.Application
	resourceApp                resources.Application
	cryptographyApp            cryptography.Application
	identityAdapter            identities.Adapter
	identityBuilder            identities.Builder
	blockchainAdapter          blockchains.Adapter
	blockchainBuilder          blockchains.Builder
	rootBuilder                roots.Builder
	rulesBuilder               rules.Builder
	blocksAdapter              blocks.Adapter
	blocksBuilder              blocks.Builder
	blockBuilder               blocks.BlockBuilder
	contentBuilder             contents.Builder
	transactionsBuilder        transactions.Builder
	transactionBuilder         transactions.TransactionBuilder
	entryBuilder               entries.Builder
	hashAdapter                hash.Adapter
	uuidAdapter                uuids.Adapter
	identityNamesList          string
	blockchainListKeyname      string
	identityKeynamePrefix      string
	identityUnitsKeynamePrefix string
	blockchainKeynamePrefix    string
	scriptKeynamePrefix        string
	blockKeynamePrefix         string
	blockQueueKeyname          string
}

func createBuilder(
	cryptographyApp cryptography.Application,
	identityAdapter identities.Adapter,
	identityBuilder identities.Builder,
	blockchainAdapter blockchains.Adapter,
	blockchainBuilder blockchains.Builder,
	rootBuilder roots.Builder,
	rulesBuilder rules.Builder,
	blocksAdapter blocks.Adapter,
	blocksBuilder blocks.Builder,
	blockBuilder blocks.BlockBuilder,
	contentBuilder contents.Builder,
	transactionsBuilder transactions.Builder,
	transactionBuilder transactions.TransactionBuilder,
	entryBuilder entries.Builder,
	hashAdapter hash.Adapter,
	uuidAdapter uuids.Adapter,
	identityNamesList string,
	blockchainListKeyname string,
	identityKeynamePrefix string,
	identityUnitsKeynamePrefix string,
	blockchainKeynamePrefix string,
	scriptKeynamePrefix string,
	blockKeynamePrefix string,
	blockQueueKeyname string,
) Builder {
	out := builder{
		listApp:                    nil,
		resourceApp:                nil,
		cryptographyApp:            cryptographyApp,
		identityAdapter:            identityAdapter,
		identityBuilder:            identityBuilder,
		blockchainAdapter:          blockchainAdapter,
		blockchainBuilder:          blockchainBuilder,
		rootBuilder:                rootBuilder,
		rulesBuilder:               rulesBuilder,
		blocksAdapter:              blocksAdapter,
		blocksBuilder:              blocksBuilder,
		blockBuilder:               blockBuilder,
		contentBuilder:             contentBuilder,
		transactionsBuilder:        transactionsBuilder,
		transactionBuilder:         transactionBuilder,
		entryBuilder:               entryBuilder,
		hashAdapter:                hashAdapter,
		uuidAdapter:                uuidAdapter,
		identityNamesList:          identityNamesList,
		blockchainListKeyname:      blockchainListKeyname,
		identityKeynamePrefix:      identityKeynamePrefix,
		identityUnitsKeynamePrefix: identityUnitsKeynamePrefix,
		blockchainKeynamePrefix:    blockchainKeynamePrefix,
		scriptKeynamePrefix:        scriptKeynamePrefix,
		blockKeynamePrefix:         blockKeynamePrefix,
		blockQueueKeyname:          blockQueueKeyname,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.cryptographyApp,
		app.identityAdapter,
		app.identityBuilder,
		app.blockchainAdapter,
		app.blockchainBuilder,
		app.rootBuilder,
		app.rulesBuilder,
		app.blocksAdapter,
		app.blocksBuilder,
		app.blockBuilder,
		app.contentBuilder,
		app.transactionsBuilder,
		app.transactionBuilder,
		app.entryBuilder,
		app.hashAdapter,
		app.uuidAdapter,
		app.identityNamesList,
		app.blockchainListKeyname,
		app.identityKeynamePrefix,
		app.identityUnitsKeynamePrefix,
		app.blockchainKeynamePrefix,
		app.scriptKeynamePrefix,
		app.blockKeynamePrefix,
		app.blockQueueKeyname,
	)
}

// WithResource adds a resource app to the builder
func (app *builder) WithResource(resourceApp resources.Application) Builder {
	app.resourceApp = resourceApp
	return app
}

// WithList adds a list app to the builder
func (app *builder) WithList(listApp lists.Application) Builder {
	app.listApp = listApp
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.resourceApp == nil {
		return nil, errors.New("the resource application is mandatory in order to build an Application instance")
	}

	if app.listApp == nil {
		return nil, errors.New("the list application is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.listApp,
		app.resourceApp,
		app.cryptographyApp,
		app.identityAdapter,
		app.identityBuilder,
		app.blockchainAdapter,
		app.blockchainBuilder,
		app.rootBuilder,
		app.rulesBuilder,
		app.blocksAdapter,
		app.blocksBuilder,
		app.blockBuilder,
		app.contentBuilder,
		app.transactionsBuilder,
		app.transactionBuilder,
		app.entryBuilder,
		app.hashAdapter,
		app.uuidAdapter,
		app.identityNamesList,
		app.blockchainListKeyname,
		app.identityKeynamePrefix,
		app.identityUnitsKeynamePrefix,
		app.blockchainKeynamePrefix,
		app.scriptKeynamePrefix,
		app.blockKeynamePrefix,
		app.blockQueueKeyname,
	), nil
}