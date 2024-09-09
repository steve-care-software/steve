package resources

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/steve-care-software/steve/domain/stores/contents"
	"github.com/steve-care-software/steve/domain/stores/headers"
	"github.com/steve-care-software/steve/domain/stores/headers/activities"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources/pointers"
)

type application struct {
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
	basePath             string
	loadedPointers       map[string]pointers.Pointer
	header               headers.Header
	pBodyIndex           *uint64
	pFile                *os.File
	updates              map[string]modifications.Modification
	additions            map[string]contents.Content
}

func createApplication(
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
) Application {
	out := application{
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
		basePath:             "",
		loadedPointers:       nil,
		header:               nil,
		pBodyIndex:           nil,
		pFile:                nil,
		updates:              map[string]modifications.Modification{},
		additions:            map[string]contents.Content{},
	}

	return &out
}

// Begin init the modifications
func (app *application) Init(dbIdentifier string) error {
	// create the path:
	loadedPointers := map[string]pointers.Pointer{}
	file := filepath.Join(app.basePath, dbIdentifier)
	fileInfo, err := os.Stat(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err := os.MkdirAll(file, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	// open the file:
	pFile, err := os.Open(file)
	if err != nil {
		return err
	}

	if fileInfo.Size() > 0 {
		retHeader, pBodyIndex, err := app.readHeader(pFile)
		if err != nil {
			return err
		}

		retLoadedPointers, err := retHeader.Map()
		if err != nil {
			return err
		}

		loadedPointers = retLoadedPointers
		app.pBodyIndex = pBodyIndex
		app.header = retHeader
	}

	app.loadedPointers = loadedPointers
	bodyIndex := uint64(0)
	app.pBodyIndex = &bodyIndex
	app.pFile = pFile
	return nil
}

// Retrieve retrieves bytes from an identifier
func (app *application) Retrieve(identifier string) ([]byte, error) {
	if app.loadedPointers != nil {
		return nil, errors.New(appNotInitErr)
	}

	if ptr, ok := app.loadedPointers[identifier]; ok {
		return app.readPointer(ptr)
	}

	str := fmt.Sprintf("the identifier (%s) does not exists", identifier)
	return nil, errors.New(str)
}

// Insert inserts data into an identifier
func (app *application) Insert(identifier string, data []byte) error {
	if _, ok := app.loadedPointers[identifier]; ok {
		str := fmt.Sprintf("the resource (identifier: %s) already exists", identifier)
		return errors.New(str)
	}

	if _, ok := app.updates[identifier]; ok {
		str := fmt.Sprintf(alreadyModifiedErrPattern, identifier)
		return errors.New(str)
	}

	resource, err := app.buildResource(identifier, data)
	if err != nil {
		return err
	}

	modification, err := app.modificationBuilder.Create().
		WithInsert(resource).
		Now()

	if err != nil {
		return err
	}

	content, err := app.contentBuilder.Create().
		WithData(data).
		WithModification(modification).
		Now()

	if err != nil {
		return err
	}

	app.updates[identifier] = modification
	app.additions[identifier] = content
	return nil
}

// Save saves data into an identifier
func (app *application) Save(identifier string, data []byte) error {
	resource, err := app.buildResource(identifier, data)
	if err != nil {
		return err
	}

	modification, err := app.modificationBuilder.Create().
		WithSave(resource).
		Now()

	if err != nil {
		return err
	}

	content, err := app.contentBuilder.Create().
		WithData(data).
		WithModification(modification).
		Now()

	if err != nil {
		return err
	}

	app.updates[identifier] = modification
	app.additions[identifier] = content
	return nil
}

// Delete deletes an identifier
func (app *application) Delete(identifier string) error {
	if _, ok := app.loadedPointers[identifier]; !ok {
		str := fmt.Sprintf("the resource (identifier: %s) does not exists and therefore cannot be deleted", identifier)
		return errors.New(str)
	}

	if _, ok := app.updates[identifier]; ok {
		str := fmt.Sprintf(alreadyModifiedErrPattern, identifier)
		return errors.New(str)
	}

	modification, err := app.modificationBuilder.Create().
		WithDelete(identifier).
		Now()

	if err != nil {
		return err
	}

	app.updates[identifier] = modification
	return nil
}

// Commit commits the modifications
func (app *application) Commit() error {
	if len(app.updates) <= 0 {
		return errors.New("there is no update to commit")
	}

	if app.header == nil {
		resourcesList := []resources.Resource{}
		for _, oneResource := range app.updates {
			if oneResource.IsDelete() {
				str := fmt.Sprintf("the resource (identifier: %s) could not be deleted because there is no resource to delete yet", oneResource.Delete())
				return errors.New(str)
			}

			if oneResource.IsSave() {
				resourcesList = append(resourcesList, oneResource.Save())
				continue
			}

			resourcesList = append(resourcesList, oneResource.Insert())
		}

		resources, err := app.resourcesBuilder.Create().WithList(resourcesList).Now()
		if err != nil {
			return err
		}

		header, err := app.headerBuilder.Create().WithRoot(resources).Now()
		if err != nil {
			return err
		}

		return app.updateSource(app.pFile, header, app.additions)
	}

	modificationsList := []modifications.Modification{}
	for _, oneModification := range app.updates {
		modificationsList = append(modificationsList, oneModification)
	}

	modifications, err := app.modificationsBuilder.Create().
		WithList(modificationsList).
		Now()

	if err != nil {
		return err
	}

	commitBuilder := app.commitBuilder.Create().
		WithModifications(modifications).
		WithParent(app.header.Root().Hash())

	if app.header.HasActivity() {
		commitBuilder.WithParent(app.header.Activity().Hash())
	}

	commit, err := commitBuilder.Now()
	if err != nil {
		return err
	}

	commitsList := app.header.Activity().Commits().List()
	commitsList = append(commitsList, commit)
	updatedCommits, err := app.commitsBuilder.Create().WithList(commitsList).Now()
	if err != nil {
		return err
	}

	activity, err := app.activityBuilder.Create().
		WithHead(commit.Hash()).
		WithCommits(updatedCommits).
		Now()

	if err != nil {
		return err
	}

	root := app.header.Root()
	updatedHeader, err := app.headerBuilder.Create().
		WithRoot(root).
		WithActivity(activity).
		Now()

	if err != nil {
		return err
	}

	return app.updateSource(app.pFile, updatedHeader, app.additions)
}

// Cancel cancels the modifications
func (app *application) Cancel() error {
	err := app.pFile.Close()
	if err != nil {
		return err
	}

	app.pFile = nil
	app.updates = nil
	app.loadedPointers = nil
	app.pBodyIndex = nil
	return nil
}

// Rollback remove the last commits
func (app *application) Rollback(amount uint) error {
	return nil
}

func (app *application) readPointer(pointer pointers.Pointer) ([]byte, error) {
	index := pointer.Index() + uint(*app.pBodyIndex)
	length := pointer.Length()

	// seek to the index
	_, err := app.pFile.Seek(int64(index), 0)
	if err != nil {
		return nil, err
	}

	// read the pointer:
	data := make([]byte, length)
	_, err = app.pFile.Read(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (app *application) readHeader(pFile *os.File) (headers.Header, *uint64, error) {
	// seek at the beginning:
	_, err := pFile.Seek(0, 0)
	if err != nil {
		return nil, nil, err
	}

	// read the header length:
	sizeInBytes := make([]byte, pointers.Uint64Size)
	_, err = pFile.Read(sizeInBytes)
	if err != nil {
		return nil, nil, err
	}

	pHeaderSize, err := pointers.BytesToUint64(sizeInBytes)
	if err != nil {
		return nil, nil, err
	}

	headerBytes := make([]byte, *pHeaderSize)
	if err != nil {
		return nil, nil, err
	}

	_, err = pFile.Read(headerBytes)
	if err != nil {
		return nil, nil, err
	}

	ins, _, err := app.headerAdapter.ToInstance(headerBytes)
	if err != nil {
		return nil, nil, err
	}

	index := pointers.Uint64Size + *pHeaderSize
	return ins, &index, nil
}

func (app *application) buildResource(identifier string, data []byte) (resources.Resource, error) {
	nextIndex := 0
	if app.header != nil {
		nextIndex = int(app.header.NextPointerIndex())
	}

	pointer, err := app.pointerBuilder.Create().
		WithIndex(uint(nextIndex)).
		WithLength(uint(len(data))).
		Now()

	if err != nil {
		return nil, err
	}

	return app.resourceBuilder.Create().
		WithIdentifier(identifier).
		WithPointer(pointer).
		Now()
}

func (app *application) updateSource(
	pFile *os.File,
	header headers.Header,
	additions map[string]contents.Content,
) error {
	return nil
}
