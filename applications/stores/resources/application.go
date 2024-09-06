package resources

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Retrieve retrieves bytes from an identifier
func (app *application) Retrieve(identifier string) ([]byte, error) {
	return nil, nil
}

// Insert inserts data into an identifier
func (app *application) Insert(identifier string, data []byte) error {
	return nil
}

// Save saves data into an identifier
func (app *application) Save(identifier string, data []byte) error {
	return nil
}

// Delete deletes an identifier
func (app *application) Delete(identifier string) error {
	return nil
}
