package repositories

import (
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/hash"
)

type blockchainRepository struct {
}

func createBlockchainRepository() blockchains.Repository {
	out := blockchainRepository{}
	return &out
}

// List returns the blockchain list
func (app *blockchainRepository) List() []hash.Hash {
	return nil
}

// Retrieve retrieves a blockchain by identifier
func (app *blockchainRepository) Retrieve(identifier hash.Hash) (blockchains.Blockchain, error) {
	return nil, nil
}
