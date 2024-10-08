package resources

import (
	"errors"

	"github.com/steve-care-software/steve/resources/domain/contents"
	"github.com/steve-care-software/steve/resources/domain/headers"
	"github.com/steve-care-software/steve/resources/domain/headers/activities"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications/resources"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications/resources/pointers"
)

type builder struct {
	contentBuilder       contents.Builder
	headerBuilder        headers.Builder
	activityBuilder      activities.Builder
	commitsBuilder       commits.Builder
	commitBuilder        commits.CommitBuilder
	modificationsBuilder modifications.Builder
	modificationBuilder  modifications.ModificationBuilder
	resourcesBuilder     resources.Builder
	resourceBuilder      resources.ResourceBuilder
	pointerBuilder       pointers.Builder
	headerAdapter        headers.Adapter
	readChunkSize        uint64
	basePath             string
	targetIdentifier     string
}

func createBuilder(
	contentBuilder contents.Builder,
	headerBuilder headers.Builder,
	activityBuilder activities.Builder,
	commitsBuilder commits.Builder,
	commitBuilder commits.CommitBuilder,
	modificationsBuilder modifications.Builder,
	modificationBuilder modifications.ModificationBuilder,
	resourcesBuilder resources.Builder,
	resourceBuilder resources.ResourceBuilder,
	pointerBuilder pointers.Builder,
	headerAdapter headers.Adapter,
) Builder {
	out := builder{
		contentBuilder:       contentBuilder,
		headerBuilder:        headerBuilder,
		activityBuilder:      activityBuilder,
		commitsBuilder:       commitsBuilder,
		commitBuilder:        commitBuilder,
		modificationsBuilder: modificationsBuilder,
		modificationBuilder:  modificationBuilder,
		resourcesBuilder:     resourcesBuilder,
		resourceBuilder:      resourceBuilder,
		pointerBuilder:       pointerBuilder,
		headerAdapter:        headerAdapter,
		readChunkSize:        0,
		basePath:             "",
		targetIdentifier:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.contentBuilder,
		app.headerBuilder,
		app.activityBuilder,
		app.commitsBuilder,
		app.commitBuilder,
		app.modificationsBuilder,
		app.modificationBuilder,
		app.resourcesBuilder,
		app.resourceBuilder,
		app.pointerBuilder,
		app.headerAdapter,
	)
}

// WithReadChunkSize adds readChunkSize to the builder
func (app *builder) WithReadChunkSize(readChkSize uint64) Builder {
	app.readChunkSize = readChkSize
	return app
}

// WithBasePath adds a basePath to the builder
func (app *builder) WithBasePath(basePath string) Builder {
	app.basePath = basePath
	return app
}

// WithTargetIdentifier adds a targetIdentifier to the builder
func (app *builder) WithTargetIdentifier(targetIdentifier string) Builder {
	app.targetIdentifier = targetIdentifier
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.readChunkSize <= 0 {
		return nil, errors.New("the readChunkSize must be greater than zero (0) in order to build an Application instance")
	}

	if app.basePath == "" {
		return nil, errors.New("the basePath is mandatory in order to build an Application instance")
	}

	if app.targetIdentifier == "" {
		return nil, errors.New("the targetIdentifier is mandatory in order to build an Application instance")
	}

	return createApplication(
		app.contentBuilder,
		app.headerBuilder,
		app.activityBuilder,
		app.commitsBuilder,
		app.commitBuilder,
		app.modificationsBuilder,
		app.modificationBuilder,
		app.resourcesBuilder,
		app.resourceBuilder,
		app.pointerBuilder,
		app.headerAdapter,
		app.readChunkSize,
		app.basePath,
		app.targetIdentifier,
	), nil
}
