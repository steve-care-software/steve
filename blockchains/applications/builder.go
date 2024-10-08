package applications

import (
	"errors"

	blockchains "github.com/steve-care-software/steve/blockchains/domain"
	"github.com/steve-care-software/steve/blockchains/domain/blocks"
	"github.com/steve-care-software/steve/blockchains/domain/blocks/contents"
	"github.com/steve-care-software/steve/blockchains/domain/blocks/contents/transactions"
	"github.com/steve-care-software/steve/blockchains/domain/blocks/contents/transactions/entries"
	"github.com/steve-care-software/steve/blockchains/domain/identities"
	"github.com/steve-care-software/steve/blockchains/domain/roots"
	"github.com/steve-care-software/steve/blockchains/domain/rules"
	"github.com/steve-care-software/steve/engine/applications/cryptography"
	"github.com/steve-care-software/steve/hash"
	lists "github.com/steve-care-software/steve/lists/applications"
	resources "github.com/steve-care-software/steve/resources/applications"
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
	identityNamesList          string
	blockchainListKeyname      string
	identityKeynamePrefix      string
	identityUnitsKeynamePrefix string
	blockchainKeynamePrefix    string
	scriptKeynamePrefix        string
	blockKeynamePrefix         string
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
	identityNamesList string,
	blockchainListKeyname string,
	identityKeynamePrefix string,
	identityUnitsKeynamePrefix string,
	blockchainKeynamePrefix string,
	scriptKeynamePrefix string,
	blockKeynamePrefix string,
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
		identityNamesList:          identityNamesList,
		blockchainListKeyname:      blockchainListKeyname,
		identityKeynamePrefix:      identityKeynamePrefix,
		identityUnitsKeynamePrefix: identityUnitsKeynamePrefix,
		blockchainKeynamePrefix:    blockchainKeynamePrefix,
		scriptKeynamePrefix:        scriptKeynamePrefix,
		blockKeynamePrefix:         blockKeynamePrefix,
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
		app.identityNamesList,
		app.blockchainListKeyname,
		app.identityKeynamePrefix,
		app.identityUnitsKeynamePrefix,
		app.blockchainKeynamePrefix,
		app.scriptKeynamePrefix,
		app.blockKeynamePrefix,
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
		app.identityNamesList,
		app.blockchainListKeyname,
		app.identityKeynamePrefix,
		app.identityUnitsKeynamePrefix,
		app.blockchainKeynamePrefix,
		app.scriptKeynamePrefix,
		app.blockKeynamePrefix,
	), nil
}
