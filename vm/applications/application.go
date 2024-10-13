package applications

type application struct {
	instructions []byte
	params       map[string][]byte
}

func createApplication(
	instructions []byte,
	params map[string][]byte,
) Application {
	out := application{
		instructions: instructions,
		params:       params,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute() ([]byte, error) {
	return nil, nil
}
