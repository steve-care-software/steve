package applications

import (
	"errors"
	"fmt"
	"time"

	command_applications "github.com/steve-care-software/steve/applications/commands"
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
	"github.com/steve-care-software/steve/domain/blockchains/roots/resolutions"
)

type application struct {
	cmdApplication       command_applications.Application
	blockchainBuilder    blockchains.Builder
	blockchainRepository blockchains.Repository
	blockchainService    blockchains.Service
	blockBuilder         blocks.Builder
	rootBuilder          roots.Builder
	resolutionBuilder    resolutions.Builder
	queueBuilder         queues.Builder
	queueRepository      queues.Repository
	queueService         queues.Service
	commandsBuilder      commands.Builder
	commandBuilder       commands.CommandBuilder
	inputAdapter         inputs.Adapter
	queue                map[uint][]commands.Command
}

func createApplication(
	cmdApplication command_applications.Application,
	blockchainBuilder blockchains.Builder,
	blockchainRepository blockchains.Repository,
	blockchainService blockchains.Service,
	blockBuilder blocks.Builder,
	rootBuilder roots.Builder,
	resolutionBuilder resolutions.Builder,
	queueBuilder queues.Builder,
	queueRepository queues.Repository,
	queueService queues.Service,
	commandsBuilder commands.Builder,
	commandBuilder commands.CommandBuilder,
	inputAdapter inputs.Adapter,
) Application {
	out := application{
		cmdApplication:       cmdApplication,
		blockchainBuilder:    blockchainBuilder,
		blockchainRepository: blockchainRepository,
		blockchainService:    blockchainService,
		blockBuilder:         blockBuilder,
		rootBuilder:          rootBuilder,
		resolutionBuilder:    resolutionBuilder,
		queueBuilder:         queueBuilder,
		queueRepository:      queueRepository,
		queueService:         queueService,
		commandsBuilder:      commandsBuilder,
		commandBuilder:       commandBuilder,
		inputAdapter:         inputAdapter,
		queue:                map[uint][]commands.Command{},
	}

	return &out
}

// Init inits the blockchain with a root and path
func (app *application) Init(root roots.Root, path string) error {
	blockchain, err := app.blockchainBuilder.Create().
		WithRoot(root).
		Now()

	if err != nil {
		return err
	}

	return app.blockchainService.Insert(
		path,
		blockchain,
	)
}

// Begin creates a context using the blockchain path and returns it
func (app *application) Begin(path string) (*uint, error) {
	return app.queueRepository.Init(path)
}

// Exists returns true if the context exists, false otherwise
func (app *application) Exists(context uint) bool {
	_, err := app.queueRepository.Retrieve(context)
	return err != nil
}

// Execute executes a command, using the passed frame and context then returns the result
func (app *application) Execute(context uint, input []byte, frame frames.Frame) ([]byte, error) {
	if !app.Exists(context) {
		str := fmt.Sprintf("the provided context (%d) does not exists", context)
		return nil, errors.New(str)
	}

	inputIns, err := app.inputAdapter.ToInput(input)
	if err != nil {
		return nil, err
	}

	exec, err := app.cmdApplication.Execute(inputIns, frame)
	if err != nil {
		return nil, err
	}

	cmd, err := app.commandBuilder.Create().
		WithExecution(exec).
		WithInput(inputIns).
		WithPreviousFrame(frame).
		Now()

	if err != nil {
		return nil, err
	}

	err = app.queueService.Append(context, cmd)
	if err != nil {
		return nil, err
	}

	if exec.HasOutput() {
		return exec.Output(), nil
	}

	return nil, nil
}

// Queue returns the commands queue
func (app *application) Queue(context uint) (queues.Queue, error) {
	return app.queueRepository.Retrieve(context)
}

// Commit commits the current commands queue to a block
func (app *application) Commit(context uint, message string) error {
	queue, err := app.queueRepository.Retrieve(context)
	if err != nil {
		return err
	}

	path := queue.Path()
	blockchain, err := app.blockchainRepository.Retrieve(path)
	if err != nil {
		return err
	}

	createdOn := time.Now().UTC()
	commands := queue.Commands()
	rootHash := blockchain.Root().Hash()
	blockBuilder := app.blockBuilder.Create().
		WithMessage(message).
		WithCommands(commands).
		CreatedOn(createdOn).
		WithParent(rootHash)

	if blockchain.HasHead() {
		headHash := blockchain.Head().Hash()
		blockBuilder.WithParent(headHash)
	}

	block, err := blockBuilder.Now()
	if err != nil {
		return err
	}

	return app.blockchainService.Chain(
		blockchain,
		block,
		func() error {
			return app.Clear(context)
		},
	)
}

// Back removes the last command from the queue
func (app *application) Back(context uint) error {
	queue, err := app.queueRepository.Retrieve(context)
	if err != nil {
		return err
	}

	commandsList := queue.Commands().List()
	if len(commandsList) <= 1 {
		return app.Clear(context)
	}

	updatedCommands, err := app.commandsBuilder.Create().
		WithList(commandsList[:len(commandsList)-1]).
		Now()

	if err != nil {
		return err
	}

	path := queue.Path()
	updatedQueue, err := app.queueBuilder.Create().
		WithPath(path).Create().
		WithCommands(updatedCommands).
		Now()

	if err != nil {
		return err
	}

	return app.queueService.Replace(context, updatedQueue)
}

// Clear removes all the commands from the queue
func (app *application) Clear(context uint) error {
	return app.queueService.Clear(context)
}

// Rollback takes the head block's commands, puts them in the current queue then remove the head block from the blockchain
func (app *application) Rollback(context uint) error {
	queue, err := app.queueRepository.Retrieve(context)
	if err != nil {
		return err
	}

	path := queue.Path()
	blockchain, err := app.blockchainRepository.Retrieve(path)
	if err != nil {
		return err
	}

	if !blockchain.HasHead() {
		str := fmt.Sprintf("the blockchain (path: %s) cannot rollback beause it contains no head block", path)
		return errors.New(str)
	}

	head := blockchain.Head()
	headCommands := head.Commands()
	headQueue, err := app.queueBuilder.Create().WithCommands(headCommands).WithPath(path).Now()
	if err != nil {
		return err
	}

	return app.blockchainService.Shrink(blockchain, func() error {
		return app.queueService.Replace(context, headQueue)
	})
}
