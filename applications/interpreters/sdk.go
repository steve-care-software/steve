package interpreters

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/steve-care-software/steve/applications/stackframes"
	"github.com/steve-care-software/steve/applications/stackframes/cursors"
	"github.com/steve-care-software/steve/domain/programs"
	"github.com/steve-care-software/steve/domain/programs/grammars"
	"github.com/steve-care-software/steve/domain/programs/instructions"
	"github.com/steve-care-software/steve/domain/stacks"
)

// SyscallFn represents the syscall func
type SyscallFn func(map[string][]byte) error

// NewApplication creates a new application
func NewApplication() Application {
	cursorApp := cursors.NewApplication()
	stackframeApp, err := stackframes.NewFactory().Create()
	if err != nil {
		panic(err)
	}

	elementsAdapter := instructions.NewElementsAdapter()
	programParserAdapter := programs.NewParserAdapter()
	return createApplication(
		stackframeApp,
		elementsAdapter,
		programParserAdapter,
		map[string]SyscallFn{
			"math_operation_arithmetic_add": func(params map[string][]byte) error {
				if firstBytes, ok := params["first"]; ok {
					if secondBytes, ok := params["second"]; ok {
						first, _ := big.NewInt(0).SetString(string(firstBytes), 0)
						if first == nil {
							return errors.New("the values passed to the first paramter could not be casted to an int")
						}

						second, _ := big.NewInt(0).SetString(string(secondBytes), 0)
						if second == nil {
							return errors.New("the values passed to the second paramter could not be casted to an int")
						}

						value := first.Add(first, second)
						fmt.Printf("\n%s, %s, %d\n", params["first"], params["second"], value.Int64())
						return nil
					}

					return errors.New("the second parameter could not be found")
				}

				return errors.New("the first parameter could not be found")
			},
			"cursor_push": func(params map[string][]byte) error {
				if valueStrBytes, ok := params["value"]; ok {
					if kindStrBytes, ok := params["kind"]; ok {
						kind, err := strconv.Atoi(string(kindStrBytes))
						if err != nil {
							return err
						}

						return cursorApp.PushAsStringBytes(valueStrBytes, uint8(kind))
					}

					return errors.New("the kind parameter could not be found")
				}

				return errors.New("the value parameter could not be found")
			},
		},
	)
}

// Application represents the interpreter application
type Application interface {
	// Execute interprets the input and returns the stack
	Execute(program programs.Program) (stacks.Stack, error)

	// Suites executes all the test suites of the grammar
	Suites(grammar grammars.Grammar) error
}
