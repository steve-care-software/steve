package jsons

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/steve-care-software/steve/domain/stencils/libraries/symbols"
)

type verifyFn = func(symbols symbols.Symbols) error

type testCase struct {
	filePath string
	verifyFn verifyFn
}

func TestExecute_Success(t *testing.T) {
	testCases := []testCase{
		{
			filePath: "./test_files/0.json",
			verifyFn: func(symbols symbols.Symbols) error {
				fmt.Printf("\n%v\n", symbols)
				return nil
			},
		},
	}

	for _, oneTestCase := range testCases {
		content, err := ioutil.ReadFile(oneTestCase.filePath)
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		application := NewJsonApplication()
		resource, err := application.Execute(content)
		if err != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
			return
		}

		fn := oneTestCase.verifyFn
		err = fn(resource)
		if err != nil {
			t.Error(err)
		}
	}

}
