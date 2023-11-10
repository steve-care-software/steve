package applications

import (
	"errors"
	"fmt"

	command_applications "github.com/steve-care-software/steve/applications/commands"
	"github.com/steve-care-software/steve/domain/blockchains"
	"github.com/steve-care-software/steve/domain/blockchains/blocks"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/frames"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/commands/inputs"
	"github.com/steve-care-software/steve/domain/blockchains/roots"
	"github.com/steve-care-software/steve/domain/blockchains/roots/resolutions"
)

type application struct {
	cmdApplication    command_applications.Application
	blockchainBuilder blockchains.Builder
	blockBuilder      blocks.Builder
	rootBuilder       roots.Builder
	resolutionBuilder resolutions.Builder
	commandsBuilder   commands.Builder
	commandBuilder    commands.CommandBuilder
	inputAdapter      inputs.Adapter
	queue             map[uint][]commands.Command
}

func createApplication(
	cmdApplication command_applications.Application,
	blockchainBuilder blockchains.Builder,
	blockBuilder blocks.Builder,
	rootBuilder roots.Builder,
	resolutionBuilder resolutions.Builder,
	commandsBuilder commands.Builder,
	commandBuilder commands.CommandBuilder,
	inputAdapter inputs.Adapter,
) Application {
	out := application{
		cmdApplication:    cmdApplication,
		blockchainBuilder: blockchainBuilder,
		blockBuilder:      blockBuilder,
		rootBuilder:       rootBuilder,
		resolutionBuilder: resolutionBuilder,
		commandsBuilder:   commandsBuilder,
		commandBuilder:    commandBuilder,
		inputAdapter:      inputAdapter,
		queue:             map[uint][]commands.Command{},
	}

	return &out
}

// Begin creates a context and returns it
func (app *application) Begin() (*uint, error) {
	ctx := uint(len(app.queue))
	app.queue[ctx] = []commands.Command{}
	return &ctx, nil
}

// Exists returns true if the context exists, false otherwise
func (app *application) Exists(context uint) bool {
	if _, ok := app.queue[context]; ok {
		return true
	}

	return false
}

// Init inits the blockchain with a root and path
func (app *application) Init(context uint, root roots.Root, path string) error {
	return nil
}

// Source sources the blockchain context with a path.  The application will retrieve the blockchain from its repository and use it in the provided context
func (app *application) Source(context uint, path string) error {
	return nil
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

	app.queue[context] = append(app.queue[context], cmd)
	if exec.HasOutput() {
		return exec.Output(), nil
	}

	return nil, nil
}

// Queue returns the commands queue
func (app *application) Queue(context uint) (commands.Commands, error) {
	return nil, nil
}

// Commit commits the current commands queue to a block
func (app *application) Commit(context uint, message string) error {
	return nil
}

// Back removes the last command from the queue
func (app *application) Back(context uint) error {
	return nil
}

// Clear removes all the commands from the queue
func (app *application) Clear(context uint) error {
	return nil
}

// Rollback takes the head block's commands, puts them in the current queue then remove the head block from the blockchain
func (app *application) Rollback(context uint) error {
	return nil
}

// Reset removes all blocks from the blockchain, but keep the root
func (app *application) Reset(context uint) error {
	return nil
}
