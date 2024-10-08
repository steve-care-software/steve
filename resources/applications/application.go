package resources

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/steve-care-software/steve/hash"
	"github.com/steve-care-software/steve/resources/domain/contents"
	"github.com/steve-care-software/steve/resources/domain/headers"
	"github.com/steve-care-software/steve/resources/domain/headers/activities"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications/resources"
	"github.com/steve-care-software/steve/resources/domain/headers/activities/commits/modifications/resources/pointers"
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
	readChunkSize        uint64
	basePath             string
	targetIdentifier     string
	loadedPointers       map[string]pointers.Pointer
	header               headers.Header
	currentIdentifier    string
	pBodyIndex           *uint64
	pFile                *os.File
	filePath             string
	updates              map[string]modifications.Modification
	additions            []contents.Content
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
	readChunkSize uint64,
	basePath string,
	targetIdentifier string,
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
		readChunkSize:        readChunkSize,
		targetIdentifier:     targetIdentifier,
		basePath:             basePath,
		loadedPointers:       nil,
		header:               nil,
		currentIdentifier:    "",
		pBodyIndex:           nil,
		pFile:                nil,
		filePath:             "",
		updates:              map[string]modifications.Modification{},
		additions:            []contents.Content{},
	}

	return &out
}

// Begin init the modifications
func (app *application) Init(dbIdentifier string) error {
	// create the path:
	loadedPointers := map[string]pointers.Pointer{}
	file := filepath.Join(app.basePath, dbIdentifier)
	fileInfo, err := app.createIfNotExists(file)
	if err != nil {
		return err
	}

	// open the file:
	bodyIndex := uint64(0)
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
		bodyIndex = *pBodyIndex
		app.header = retHeader
	}

	app.loadedPointers = loadedPointers
	app.currentIdentifier = dbIdentifier
	app.pBodyIndex = &bodyIndex
	app.pFile = pFile
	app.filePath = file
	app.updates = map[string]modifications.Modification{}
	app.additions = []contents.Content{}
	return nil
}

// Retrieve retrieves bytes from an identifier
func (app *application) Retrieve(identifier string) ([]byte, error) {
	if app.loadedPointers == nil {
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

	var lastUpdate contents.Content
	if len(app.additions) > 0 {
		lastUpdate = app.additions[len(app.additions)-1]
	}

	resource, err := app.buildResource(identifier, data, lastUpdate)
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
	app.additions = append(app.additions, content)
	return nil
}

// Save saves data into an identifier
func (app *application) Save(identifier string, data []byte) error {
	var lastUpdate contents.Content
	if len(app.additions) > 0 {
		lastUpdate = app.additions[len(app.additions)-1]
	}

	resource, err := app.buildResource(identifier, data, lastUpdate)
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
	app.additions = append(app.additions, content)
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
		for _, oneResource := range app.additions {
			modification := oneResource.Modification()
			if modification.IsDelete() {
				str := fmt.Sprintf("the resource (identifier: %s) could not be deleted because there is no resource to delete yet", modification.Delete())
				return errors.New(str)
			}

			if modification.IsSave() {
				resourcesList = append(resourcesList, modification.Save())
				continue
			}

			resourcesList = append(resourcesList, modification.Insert())
		}

		resources, err := app.resourcesBuilder.Create().WithList(resourcesList).Now()
		if err != nil {
			return err
		}

		header, err := app.headerBuilder.Create().WithRoot(resources).Now()
		if err != nil {
			return err
		}

		err = app.updateSource(
			app.filePath,
			app.pFile,
			header,
			app.additions,
		)

		if err != nil {
			return err
		}

		// cleanup:
		err = app.Cancel()
		if err != nil {
			return err
		}

		return app.Init(app.currentIdentifier)
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
		commitBuilder.WithParent(app.header.Activity().Head())
	}

	commit, err := commitBuilder.Now()
	if err != nil {
		return err
	}

	commitsList := []commits.Commit{}
	if app.header.HasActivity() {
		commitsList = app.header.Activity().Commits().List()
	}

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

	err = app.updateSource(
		app.filePath,
		app.pFile,
		updatedHeader,
		app.additions,
	)

	if err != nil {
		return err
	}

	// cleanup:
	err = app.Cancel()
	if err != nil {
		return err
	}

	return app.Init(app.currentIdentifier)
}

// Head returns the commit head
func (app *application) Head() (commits.Commit, error) {
	if app.header == nil {
		return nil, errors.New("there is no header, which means that the database cannot be rollbacked (not enough past commits) or it has never been initialized")
	}

	if !app.header.HasActivity() {
		return nil, errors.New("there is no activity, which means that the database cannot be rollbacked (not enough past commits)")
	}

	activity := app.header.Activity()
	head := activity.Head()
	commits := activity.Commits()
	return commits.Fetch(head)
}

// Cancel cancels the modifications
func (app *application) Cancel() error {
	defer app.pFile.Close()
	app.pFile = nil
	app.updates = nil
	app.loadedPointers = nil
	app.pBodyIndex = nil
	return nil
}

// Rollback remove the last commits
func (app *application) Rollback(amount uint) error {
	if app.header == nil {
		return errors.New("there is no header, which means that the database cannot be rollbacked (not enough past commits) or it has never been initialized")
	}

	if !app.header.HasActivity() {
		return errors.New("there is no activity, which means that the database cannot be rollbacked (not enough past commits)")
	}

	currentHeadCommit, err := app.Head()
	if err != nil {
		return err
	}

	commits := app.header.Activity().Commits()
	currentHeadCommitHash := currentHeadCommit.Hash()
	castedAmount := int(amount)
	for i := 0; i < castedAmount; i++ {
		isLast := (i + 1) == castedAmount
		parent := currentHeadCommit.Parent()
		retCommit, err := commits.Fetch(parent)
		if err != nil {
			if !isLast {
				return err
			}

			currentHeadCommitHash = app.header.Root().Hash()
			break
		}

		currentHeadCommitHash = retCommit.Hash()
	}

	return app.RollbackTo(currentHeadCommitHash)
}

// RollbackTo rollbacks to the provided commit hash
func (app *application) RollbackTo(head hash.Hash) error {
	if app.header == nil {
		return errors.New("there is no header, which means that the database cannot be rollbacked (not enough past commits) or it has never been initialized")
	}

	if !app.header.HasActivity() {
		return errors.New("there is no activity, which means that the database cannot be rollbacked (not enough past commits)")
	}

	root := app.header.Root()
	activity := app.header.Activity()
	commits := activity.Commits()
	builder := app.headerBuilder.Create().WithRoot(root)
	if !root.Hash().Compare(head) {
		activity, err := app.activityBuilder.Create().
			WithHead(head).
			WithCommits(commits).
			Now()

		if err != nil {
			return err
		}

		builder.WithActivity(activity)
	}

	retHeader, err := builder.Now()
	if err != nil {
		return err
	}

	err = app.updateSource(
		app.filePath,
		app.pFile,
		retHeader,
		[]contents.Content{},
	)

	if err != nil {
		return err
	}

	// cleanup:
	err = app.Cancel()
	if err != nil {
		return err
	}

	return app.Init(app.currentIdentifier)
}

func (app *application) readPointer(pointer pointers.Pointer) ([]byte, error) {
	index := pointer.Index() + uint(*app.pBodyIndex)
	length := pointer.Length()

	// seek to the index
	_, err := app.pFile.Seek(int64(index), io.SeekStart)
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
	_, err := pFile.Seek(0, io.SeekStart)
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

func (app *application) buildResource(identifier string, data []byte, lastUpdate contents.Content) (resources.Resource, error) {
	pNextIndex, err := app.nextIndex(lastUpdate)
	if err != nil {
		return nil, err
	}

	pointer, err := app.pointerBuilder.Create().
		WithIndex(uint(*pNextIndex)).
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

func (app *application) nextIndex(lastUpdate contents.Content) (*uint, error) {
	if lastUpdate != nil {
		modification := lastUpdate.Modification()
		if modification.IsInsert() {
			pointer := modification.Insert().Pointer()
			value := uint(pointer.Index() + pointer.Length())
			return &value, nil
		}

		if modification.IsSave() {
			pointer := modification.Save().Pointer()
			value := uint(pointer.Index() + pointer.Length())
			return &value, nil
		}
	}

	nextIndex := uint(0)
	if app.header != nil {
		pNextPointerIndex, err := app.header.NextPointerIndex()
		if err != nil {
			return nil, err
		}

		nextIndex = uint(*pNextPointerIndex)
	}

	return &nextIndex, nil
}

func (app *application) updateSource(
	originalFilePath string,
	pOriginalFile *os.File,
	header headers.Header,
	additions []contents.Content,
) error {
	// create the header bytes:
	headerBytes, err := app.headerToBytes(header)
	if err != nil {
		return err
	}

	// open the temporary file:
	file := filepath.Join(app.basePath, app.targetIdentifier)
	_, err = app.createIfNotExists(file)
	if err != nil {
		return err
	}

	// open the file:
	pTmpFile, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, fs.ModeAppend)
	if err != nil {
		return err
	}

	// close the file at the end:
	defer pTmpFile.Close()

	//seek to the beginning:
	_, err = pTmpFile.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	// write the header bytes:
	_, err = pTmpFile.Write(headerBytes)
	if err != nil {
		return err
	}

	// copy the body:
	err = app.copyBody(pOriginalFile, pTmpFile, *app.pBodyIndex)
	if err != nil {
		return err
	}

	// close the original file:
	err = pOriginalFile.Close()
	if err != nil {
		return err
	}

	// write the additions to the body:
	for _, oneAddition := range additions {
		data := oneAddition.Data()
		_, writeErr := pTmpFile.Write(data)
		if writeErr != nil {
			return fmt.Errorf("failed to write to destination file: %w", writeErr)
		}
	}

	// rename the file:
	return os.Rename(file, originalFilePath)
}

func (app *application) copyBody(
	pSourceFile *os.File,
	pTargetFile *os.File,
	indexInSource uint64,
) error {
	//seek to the body:
	_, err := pSourceFile.Seek(int64(indexInSource), io.SeekStart)
	if err != nil {
		return err
	}

	// Create a buffer to hold the chunks
	buf := make([]byte, app.readChunkSize)

	// Read and write in chunks
	for {
		// Read a chunk from the source file
		n, readErr := pSourceFile.Read(buf)
		if readErr != nil && readErr != io.EOF {
			return fmt.Errorf("failed to read from source file: %w", readErr)
		}

		// Break if we've reached the end of the file
		if n == 0 {
			break
		}

		// Write the chunk to the destination file
		_, writeErr := pTargetFile.Write(buf[:n])
		if writeErr != nil {
			return fmt.Errorf("failed to write to destination file: %w", writeErr)
		}
	}

	return nil
}

func (app *application) headerToBytes(header headers.Header) ([]byte, error) {
	headerBytes, err := app.headerAdapter.ToBytes(header)
	if err != nil {
		return nil, err
	}

	length := uint64(len(headerBytes))
	lengthBytes := pointers.Uint64ToBytes(length)
	headerWithBytes := []byte{}
	headerWithBytes = append(headerWithBytes, lengthBytes...)
	return append(headerWithBytes, headerBytes...), nil
}

func (app *application) createIfNotExists(
	filePath string,
) (os.FileInfo, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			dirPath := filepath.Dir(filePath)
			err := os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				return nil, err
			}

			pCreatedFile, err := os.Create(filePath)
			if err != nil {
				return nil, err
			}

			err = pCreatedFile.Close()
			if err != nil {
				return nil, err
			}
		}

		retFileInfo, err := os.Stat(filePath)
		if err != nil {
			return nil, err
		}

		fileInfo = retFileInfo
	}

	return fileInfo, nil
}
