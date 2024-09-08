package resources

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/steve-care-software/steve/domain/stores/headers"
	"github.com/steve-care-software/steve/domain/stores/headers/activities/commits/modifications/resources/pointers"
)

type application struct {
	headerAdapter  headers.Adapter
	basePath       string
	loadedPointers map[string]pointers.Pointer
	pBodyIndex     *uint64
	pFile          *os.File
	inserts        map[string][]byte
	saves          map[string][]byte
	deletes        []string
}

func createApplication(
	headerAdapter headers.Adapter,
) Application {
	out := application{
		headerAdapter:  headerAdapter,
		basePath:       "",
		loadedPointers: nil,
		pBodyIndex:     nil,
		pFile:          nil,
		inserts:        map[string][]byte{},
		saves:          map[string][]byte{},
		deletes:        []string{},
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
		retLoadedPointers, pBodyIndex, err := app.readHeader(pFile)
		if err != nil {
			return err
		}

		loadedPointers = retLoadedPointers
		app.pBodyIndex = pBodyIndex
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
		str := fmt.Sprintf("the identifier (name: %s) already exists", identifier)
		return errors.New(str)
	}

	if _, ok := app.inserts[identifier]; ok {
		str := fmt.Sprintf("the identifier (name: %s) has already been inserted in the current session", identifier)
		return errors.New(str)
	}

	app.inserts[identifier] = data
	return nil
}

// Save saves data into an identifier
func (app *application) Save(identifier string, data []byte) {
	app.saves[identifier] = data
}

// Delete deletes an identifier
func (app *application) Delete(identifier string) error {
	if _, ok := app.loadedPointers[identifier]; !ok {
		str := fmt.Sprintf("the identifier (name: %s) does not exists and therefore cannot be deleted", identifier)
		return errors.New(str)
	}

	for _, oneDelete := range app.deletes {
		if oneDelete == identifier {
			str := fmt.Sprintf("the identifier (name: %s) has already been deleted in the current session", oneDelete)
			return errors.New(str)
		}
	}

	app.deletes = append(app.deletes, identifier)
	return nil
}

// Commit commits the modifications
func (app *application) Commit() error {
	return nil
}

// Cancel cancels the modifications
func (app *application) Cancel() error {
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

func (app *application) readHeader(pFile *os.File) (map[string]pointers.Pointer, *uint64, error) {
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
	retMap, err := ins.Map()
	if err != nil {
		return nil, nil, err
	}

	return retMap, &index, nil
}
