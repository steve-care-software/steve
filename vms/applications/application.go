package applications

import (
	"time"

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
	if instruction.IsInit() {

	}

	if instruction.IsAssignment() {

	}

	if instruction.IsDelete() {

	}

	init := instruction.Init()
	return app.executeInit(init)
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
