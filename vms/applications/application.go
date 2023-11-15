package applications

import (
	"time"

	commnd_resources "github.com/steve-care-software/steve/vms/children/commands/domain/resources"
	"github.com/steve-care-software/steve/vms/domain/frames"
	"github.com/steve-care-software/steve/vms/domain/programs"
	"github.com/steve-care-software/steve/vms/domain/resources"
	"github.com/steve-care-software/steve/vms/domain/resources/blocks"
	"github.com/steve-care-software/steve/vms/domain/resources/blocks/queues"
	"github.com/steve-care-software/steve/vms/domain/resources/roots"
	"github.com/steve-care-software/steve/vms/domain/resources/roots/resolutions"
	"github.com/steve-care-software/steve/vms/domain/results"
)

type application struct {
	frameFactory            frames.FrameFactory
	frameBuilder            frames.FrameBuilder
	frameAssignablesBuilder frames.AssignablesBuilder
	resultBuilder           results.Builder
	blockchainBuilder       resources.Builder
	blockchainRepository    resources.Repository
	blockchainService       resources.Service
	blockBuilder            blocks.Builder
	rootBuilder             roots.Builder
	resolutionBuilder       resolutions.Builder
	queueBuilder            queues.Builder
	queueRepository         queues.Repository
	queueService            queues.Service
	commandsBuilder         commnd_resources.Builder
}

func createApplication(
	frameFactory frames.FrameFactory,
	frameBuilder frames.FrameBuilder,
	frameAssignablesBuilder frames.AssignablesBuilder,
	resultBuilder results.Builder,
	blockchainBuilder resources.Builder,
	blockchainRepository resources.Repository,
	blockchainService resources.Service,
	blockBuilder blocks.Builder,
	rootBuilder roots.Builder,
	resolutionBuilder resolutions.Builder,
	queueBuilder queues.Builder,
	queueRepository queues.Repository,
	queueService queues.Service,
	commandsBuilder commnd_resources.Builder,
) Application {
	out := application{
		frameFactory:            frameFactory,
		frameBuilder:            frameBuilder,
		frameAssignablesBuilder: frameAssignablesBuilder,
		resultBuilder:           resultBuilder,
		blockchainBuilder:       blockchainBuilder,
		blockchainRepository:    blockchainRepository,
		blockchainService:       blockchainService,
		blockBuilder:            blockBuilder,
		rootBuilder:             rootBuilder,
		resolutionBuilder:       resolutionBuilder,
		queueBuilder:            queueBuilder,
		queueRepository:         queueRepository,
		queueService:            queueService,
		commandsBuilder:         commandsBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(programm programs.Program, frame frames.Frame) (results.Result, error) {
	innerFrame, err := app.createInnerFrame(programm, frame)
	if err != nil {
		return nil, err
	}

	instructions := programm.Instructions()
	err = app.executeInstructions(instructions, innerFrame)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (app *application) executeInstructions(instructions programs.Instructions, frame frames.Frame) error {
	list := instructions.List()
	for _, oneInstruction := range list {
		err := app.executeInstruction(oneInstruction, frame)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *application) executeInstruction(instruction programs.Instruction, frame frames.Frame) error {
	if instruction.IsAssignment() {

	}

	if instruction.IsDelete() {

	}

	init := instruction.Init()
	return app.executeInit(init)
}

func (app *application) executeAssignment(assignment programs.Assignment, frame frames.Frame) error {
	assignable := assignment.Assignable()
	if assignable.IsBack() {
		variable := assignable.Back()
		return app.executeBack(variable, frame)
	}

	if assignable.IsBegin() {

	}

	if assignable.IsClear() {

	}

	if assignable.IsCommit() {
		commit := assignable.Commit()
		return app.executeCommit(commit, frame)
	}

	if assignable.IsExists() {

	}

	if assignable.IsQueue() {

	}

	if assignable.IsRollback() {

	}

	if assignable.IsTransact() {

	}

	return nil
}

func (app *application) executeCommit(commit programs.Commit, frame frames.Frame) error {
	contextVariable := commit.Context()
	pContext, err := frame.FetchContext(contextVariable)
	if err != nil {
		return err
	}

	message := commit.Message()

	queue, err := app.queueRepository.Retrieve(*pContext)
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
			return app.queueService.Clear(*pContext)
		},
	)
}

func (app *application) executeBack(variable string, frame frames.Frame) error {
	pContext, err := frame.FetchContext(variable)
	if err != nil {
		return err
	}

	queue, err := app.queueRepository.Retrieve(*pContext)
	if err != nil {
		return err
	}

	commandsList := queue.Commands().List()
	if len(commandsList) <= 1 {
		return app.queueService.Clear(*pContext)
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

	return app.queueService.Replace(
		*pContext,
		updatedQueue,
	)
}

func (app *application) executeInit(init programs.Init) error {
	rootIns := init.Root()
	affiliate := rootIns.Affiliate()
	fees := rootIns.Fees()
	resolution, err := app.resolutionBuilder.Create().
		WithAffiliate(affiliate).
		WithFees(fees).
		Now()

	if err != nil {
		return err
	}

	createdOn := time.Now().UTC()
	root, err := app.rootBuilder.Create().WithResolution(resolution).CreatedOn(createdOn).Now()
	if err != nil {
		return err
	}

	blockchain, err := app.blockchainBuilder.Create().
		WithRoot(root).
		Now()

	if err != nil {
		return err
	}

	path := init.Path()
	return app.blockchainService.Insert(
		path,
		blockchain,
	)
}

func (app *application) createInnerFrame(programm programs.Program, outerFrame frames.Frame) (frames.Frame, error) {
	if !programm.HasParameters() {
		return app.frameFactory.Create(), nil
	}

	list := []frames.Assignable{}
	parameters := programm.Parameters()
	for _, oneParameter := range parameters {
		assignable, err := outerFrame.Fetch(oneParameter)
		if err != nil {
			// error, param is not in frame
		}

		list = append(list, assignable)
	}

	assignables, err := app.frameAssignablesBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		return nil, err
	}

	return app.frameBuilder.Create().
		WithAssignables(assignables).
		Now()
}
