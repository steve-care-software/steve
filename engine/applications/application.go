package applications

import (
	"errors"

	lists "github.com/steve-care-software/steve/databases/lists/applications"
	resources "github.com/steve-care-software/steve/databases/resources/applications"
	chains_applications "github.com/steve-care-software/steve/engine/applications/chains"
	"github.com/steve-care-software/steve/engine/domain/chains"
	"github.com/steve-care-software/steve/engine/domain/nfts"
)

type application struct {
	chainApplication    chains_applications.Application
	listApplication     lists.Application
	resourceApplication resources.Application
	nftAdapter          nfts.Adapter
	chainAdapter        chains.Adapter
	rootChainListName   string
}

func createApplication(
	chainApplication chains_applications.Application,
	listApplication lists.Application,
	resourceApplication resources.Application,
	nftAdapter nfts.Adapter,
	chainAdapter chains.Adapter,
	rootChainListName string,
) Application {
	out := application{
		chainApplication:    chainApplication,
		listApplication:     listApplication,
		resourceApplication: resourceApplication,
		nftAdapter:          nftAdapter,
		chainAdapter:        chainAdapter,
		rootChainListName:   rootChainListName,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(input []byte) ([]byte, error) {
	pAmount, err := app.listApplication.Amount(app.rootChainListName)
	if err != nil {
		return nil, err
	}

	chunkSize := uint(20)
	amount := *pAmount
	amountLoops := int(amount/chunkSize) + 1
	for i := 0; i < amountLoops; i++ {
		index := uint(i) * chunkSize
		values, err := app.listApplication.Retrieve(app.rootChainListName, index, chunkSize)
		if err != nil {
			return nil, err
		}

		for _, oneChainHashStr := range values {
			retBytes, err := app.resourceApplication.Retrieve(string(oneChainHashStr))
			if err != nil {
				return nil, err
			}

			retNFT, err := app.nftAdapter.ToNFT(retBytes)
			if err != nil {
				return nil, err
			}

			retChain, err := app.chainAdapter.ToInstance(retNFT)
			if err != nil {
				return nil, err
			}

			retAfterExecBytes, err := app.chainApplication.Execute(retChain, input)
			if err != nil {
				continue
			}

			return retAfterExecBytes, nil
		}
	}

	return nil, errors.New("the provided input could not be understood and interpreted")
}
