package services

import (
	"encoding/binary"
	"fmt"

	"github.com/steve-care-software/steve/applications/blockchains/databases"
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commits/actions/inserts"
	"github.com/steve-care-software/steve/domain/hash"
)

type blockchainService struct {
	dbApp                databases.Application
	blockchainRepository blockchains.Repository
	blockchainAdapter    blockchains.Adapter
	readChunkSize        uint
}

func createBlockchainService(
	dbApp databases.Application,
	blockchainRepository blockchains.Repository,
	blockchainAdapter blockchains.Adapter,
	readChunkSize uint,
) blockchains.Service {
	out := blockchainService{
		dbApp:                dbApp,
		blockchainRepository: blockchainRepository,
		blockchainAdapter:    blockchainAdapter,
		readChunkSize:        readChunkSize,
	}

	return &out
}

// Save saves a blockchain with inserts
func (app *blockchainService) Save(blockchain blockchains.Blockchain, inserts inserts.Inserts) error {
	identifier := blockchain.Identifier()
	name := identifier.String()
	processName := fmt.Sprintf("%s.process", name)
	err := app.dbApp.New(processName)
	if err != nil {
		return err
	}

	pContext, err := app.dbApp.Open(processName)
	if err != nil {
		return err
	}

	err = app.dbApp.Lock(*pContext)
	if err != nil {
		return err
	}

	defer func() {
		app.dbApp.Unlock(*pContext)
		app.dbApp.Close(*pContext)
	}()

	// convert the blockchain instance to bytes:
	blockchainBytes, err := app.blockchainAdapter.ToBytes(blockchain)
	if err != nil {
		return err
	}

	// convert the length to bytes:
	lengthAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lengthAsBytes, uint64(len(blockchainBytes)))

	// write the blockchain data:
	bytes := []byte{}
	bytes = append(bytes, lengthAsBytes...)
	bytes = append(bytes, blockchainBytes...)
	err = app.dbApp.Write(*pContext, int64(0), bytes)
	if err != nil {
		return err
	}

	// write blockchain bytes:
	err = app.dbApp.Write(*pContext, 0, bytes)
	if err != nil {
		return err
	}

	// merge the data:
	err = app.mergeCurrentBytes(identifier, *pContext, uint(len(bytes)))
	if err != nil {
		return err
	}

	// add the new transaction data:
	list := inserts.List()
	for _, oneInsert := range list {
		resource := oneInsert.Resource()
		bytes := oneInsert.Bytes()
		err = app.dbApp.Write(*pContext, int64(resource.Pointer().Index()), bytes)
		if err != nil {
			return err
		}
	}

	return nil

}

func (app *blockchainService) mergeCurrentBytes(identifier hash.Hash, writeContext uint, writeIndex uint) error {
	// fetch the current database instance:
	retBlockchain, err := app.blockchainRepository.Retrieve(identifier)
	if err != nil {
		return err
	}

	// fetch the read indexes:
	fromIndex := retBlockchain.Pointer().Next()
	toIndex := retBlockchain.Head().Content().Pointer().Delimiter()

	// open the read context:
	name := identifier.String()
	pReadContext, err := app.dbApp.Open(name)
	if err != nil {
		return err
	}

	err = app.dbApp.Lock(*pReadContext)
	if err != nil {
		return err
	}

	defer func() {
		app.dbApp.Close(*pReadContext)
	}()

	currentWriteIndex := writeIndex
	amount := toIndex - fromIndex/uint(app.readChunkSize)
	for i := 0; i < int(amount); i++ {
		// read the chunk:
		index := fromIndex + (uint(i) * app.readChunkSize)
		bytes, err := app.dbApp.Read(*pReadContext, index, uint(app.readChunkSize))
		if err != nil {
			return err
		}

		// write the chunk:
		err = app.dbApp.Write(writeContext, int64(currentWriteIndex), bytes)
		if err != nil {
			return err
		}

		// update the index:
		currentWriteIndex += uint(len(bytes))
	}

	return nil
}

// Delete deletes a blockchain
func (app *blockchainService) Delete(identifier hash.Hash) error {
	return nil
}
