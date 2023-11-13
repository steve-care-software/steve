package identities

import (
	"github.com/steve-care-software/steve/applications/commands/identities/connections"
	"github.com/steve-care-software/steve/applications/commands/identities/encryptors"
	"github.com/steve-care-software/steve/applications/commands/identities/identities"
	"github.com/steve-care-software/steve/applications/commands/identities/profiles"
	"github.com/steve-care-software/steve/applications/commands/identities/shares"
	"github.com/steve-care-software/steve/applications/commands/identities/signers"
	"github.com/steve-care-software/steve/applications/commands/shares/dashboards"
	executions "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/executions/identities"
	"github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/frames"
	inputs "github.com/steve-care-software/steve/domain/blockchains/blocks/queues/commands/inputs/identities"
)

type application struct {
	identityApp      identities.Application
	dashboardApp     dashboards.Application
	connectionApp    connections.Application
	profileApp       profiles.Application
	sharesApp        shares.Application
	encryptorApp     encryptors.Application
	signerApp        signers.Application
	executionBuilder executions.Builder
}

func createApplication(
	identityApp identities.Application,
	dashboardApp dashboards.Application,
	connectionApp connections.Application,
	profileApp profiles.Application,
	sharesApp shares.Application,
	encryptorApp encryptors.Application,
	signerApp signers.Application,
	executionBuilder executions.Builder,
) Application {
	out := application{
		identityApp:      identityApp,
		dashboardApp:     dashboardApp,
		connectionApp:    connectionApp,
		profileApp:       profileApp,
		sharesApp:        sharesApp,
		encryptorApp:     encryptorApp,
		signerApp:        signerApp,
		executionBuilder: executionBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(identity inputs.Identity, frame frames.Frame) (executions.Identity, error) {
	builder := app.executionBuilder.Create()
	if identity.IsIdentity() {
		identity := identity.Identity()
		exec, err := app.identityApp.Execute(identity, frame)
		if err != nil {
			return nil, err
		}

		builder.WithIdentity(exec)
	}

	if identity.IsDashboard() {
		dashboard := identity.Dashboard()
		exec, err := app.dashboardApp.Execute(dashboard, frame)
		if err != nil {
			return nil, err
		}

		builder.WithDashboard(exec)
	}

	if identity.IsConnections() {
		connections := identity.Connections()
		exec, err := app.connectionApp.Execute(connections, frame)
		if err != nil {
			return nil, err
		}

		builder.WithConnections(exec)
	}

	if identity.IsProfile() {
		profile := identity.Profile()
		exec, err := app.profileApp.Execute(profile, frame)
		if err != nil {
			return nil, err
		}

		builder.WithProfile(exec)
	}

	if identity.IsShares() {
		shares := identity.Shares()
		exec, err := app.sharesApp.Execute(shares, frame)
		if err != nil {
			return nil, err
		}

		builder.WithShares(exec)
	}

	if identity.IsEncryptor() {
		encryptor := identity.Encryptor()
		exec, err := app.encryptorApp.Execute(encryptor, frame)
		if err != nil {
			return nil, err
		}

		builder.WithEncryptor(exec)
	}

	if identity.IsSigner() {
		signer := identity.Signer()
		exec, err := app.signerApp.Execute(signer, frame)
		if err != nil {
			return nil, err
		}

		builder.WithSigner(exec)
	}

	return builder.Now()
}
