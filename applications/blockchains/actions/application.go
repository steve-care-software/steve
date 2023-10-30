package actions

import (
	"errors"
	"time"

	"github.com/steve-care-software/steve/domain/accounts/identities/signers"
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/headers"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/headers/identifiers"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/headers/links"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/resources/pointers"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/transactions"
	"github.com/steve-care-software/steve/domain/hash"
)

type application struct {
	blockchainService   blockchains.Service
	blockchainBuilder   blockchains.Builder
	blockBuilder        blocks.Builder
	contentBuilder      blocks.ContentBuilder
	parentBuilder       blocks.ParentBuilder
	commitsBuilder      commits.Builder
	commitBuilder       commits.CommitBuilder
	resourcesBuilder    resources.Builder
	resourceBuilder     resources.ResourceBuilder
	pointerBuilder      pointers.Builder
	headerBuilder       headers.Builder
	linksBuilder        links.Builder
	identifiersBuilder  identifiers.Builder
	transactionsBuilder transactions.Builder
	hashAdapter         hash.Adapter
	signer              signers.Signer
	current             blockchains.Blockchain
	trxQueue            transactions.Transactions
	commitsQueue        commits.Commits
}

func createApplication(
	blockchainService blockchains.Service,
	blockchainBuilder blockchains.Builder,
	blockBuilder blocks.Builder,
	contentBuilder blocks.ContentBuilder,
	parentBuilder blocks.ParentBuilder,
	commitsBuilder commits.Builder,
	commitBuilder commits.CommitBuilder,
	resourcesBuilder resources.Builder,
	resourceBuilder resources.ResourceBuilder,
	pointerBuilder pointers.Builder,
	headerBuilder headers.Builder,
	linksBuilder links.Builder,
	identifiersBuilder identifiers.Builder,
	transactionsBuilder transactions.Builder,
	hashAdapter hash.Adapter,
	signer signers.Signer,
	current blockchains.Blockchain,
) Application {
	out := application{
		blockchainService:   blockchainService,
		blockchainBuilder:   blockchainBuilder,
		blockBuilder:        blockBuilder,
		contentBuilder:      contentBuilder,
		parentBuilder:       parentBuilder,
		commitsBuilder:      commitsBuilder,
		commitBuilder:       commitBuilder,
		resourcesBuilder:    resourcesBuilder,
		resourceBuilder:     resourceBuilder,
		pointerBuilder:      pointerBuilder,
		headerBuilder:       headerBuilder,
		linksBuilder:        linksBuilder,
		identifiersBuilder:  identifiersBuilder,
		transactionsBuilder: transactionsBuilder,
		hashAdapter:         hashAdapter,
		signer:              signer,
		current:             current,
		trxQueue:            nil,
		commitsQueue:        nil,
	}

	return &out
}

// Insert inserts a transaction
func (app *application) Insert(trx transactions.Transaction) error {
	list := []transactions.Transaction{}
	if app.trxQueue != nil {
		list = app.trxQueue.List()
	}

	list = append(list, trx)
	queue, err := app.transactionsBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		return err
	}

	app.trxQueue = queue
	return nil
}

// Commit commit actions
func (app *application) Commit(message string) error {
	if app.trxQueue != nil {
		return errors.New("the transaction queue is empty: nothing to commit")
	}

	list := app.trxQueue.List()
	resourcesList := []resources.Resource{}
	nextIndex := app.current.Pointer().Next()
	for _, oneTrx := range list {
		kind := oneTrx.Kind()
		bytes := oneTrx.Bytes()
		pHash, err := app.hashAdapter.FromBytes(bytes)
		if err != nil {
			return err
		}

		identifier, err := app.identifiersBuilder.Create().
			WithHash(*pHash).
			WithKind(kind).
			Now()

		if err != nil {
			return err
		}

		link := oneTrx.Link()
		header, err := app.headerBuilder.Create().
			WithIdentifier(identifier).
			WithLink(link).
			Now()

		if err != nil {
			return err
		}

		pointer, err := app.pointerBuilder.Create().
			WithIndex(nextIndex).
			WithLength(uint(len(bytes))).
			Now()

		if err != nil {
			return err
		}

		createdOn := time.Now().UTC()
		resource, err := app.resourceBuilder.Create().
			WithHeader(header).
			WithPointer(pointer).
			CreatedOn(createdOn).
			Now()

		if err != nil {
			return err
		}

		resourcesList = append(resourcesList, resource)
	}

	resources, err := app.resourcesBuilder.Create().
		WithList(resourcesList).
		Now()

	if err != nil {
		return err
	}

	commit, err := app.commitBuilder.Create().
		WithResources(resources).
		WithMessage(message).
		Now()

	if err != nil {
		return err
	}

	commitsList := []commits.Commit{}
	if app.commitsQueue != nil {
		commitsList = app.commitsQueue.List()
	}

	commitsList = append(commitsList, commit)
	updatedCommits, err := app.commitsBuilder.Create().
		WithList(commitsList).
		Now()

	if err != nil {
		return err
	}

	app.commitsQueue = updatedCommits
	app.trxQueue = nil
	return nil
}

// Push pushes commits
func (app *application) Push() error {
	if app.commitsQueue != nil {
		return errors.New("the commit queue is empty: nothing to push")
	}

	head := app.current.Head().Content()
	parent := head.Parent()
	content, err := app.contentBuilder.Create().
		WithCommits(app.commitsQueue).
		WithParent(parent).
		Now()

	if err != nil {
		return err
	}

	msg := content.Hash().Bytes()
	vote, err := app.signer.Vote(msg)
	if err != nil {
		return err
	}

	newHead, err := app.blockBuilder.Create().
		WithContent(content).
		WithVote(vote).
		Now()

	if err != nil {
		return err
	}

	identifier := app.current.Identifier()
	updatedBlockchain, err := app.blockchainBuilder.Create().
		WithIdentifier(identifier).
		WithHead(newHead).
		Now()

	if err != nil {
		return err
	}

	err = app.blockchainService.Save(updatedBlockchain)
	if err != nil {
		return err
	}

	app.commitsQueue = nil
	return nil
}

// Rollback rollbacks 1 commit
func (app *application) Rollback() error {
	if app.commitsQueue != nil {
		return errors.New("the commit queue is empty: nothing to rollback")
	}

	list := app.commitsQueue.List()
	if len(list) <= 1 {
		app.commitsQueue = nil
		return nil
	}

	updated, err := app.commitsBuilder.Create().WithList(list[1:]).Now()
	if err != nil {
		return err
	}

	app.commitsQueue = updated
	return nil
}

// Cancel cancels all commits and transactions
func (app *application) Cancel() {
	app.commitsQueue = nil
	app.trxQueue = nil
}
