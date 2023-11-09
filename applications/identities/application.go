package identities

import (
	"github.com/steve-care-software/steve/applications/identities/connections"
	"github.com/steve-care-software/steve/applications/identities/encryptors"
	"github.com/steve-care-software/steve/applications/identities/identities"
	"github.com/steve-care-software/steve/applications/identities/profiles"
	"github.com/steve-care-software/steve/applications/identities/shares"
	"github.com/steve-care-software/steve/applications/identities/signers"
	"github.com/steve-care-software/steve/applications/shares/dashboards"
	executions "github.com/steve-care-software/steve/domain/commands/executions/identities"
	inputs "github.com/steve-care-software/steve/domain/commands/inputs/identities"
	"github.com/steve-care-software/steve/domain/stacks"
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
func (app *application) Execute(identity inputs.Identity, stack stacks.Stack) (executions.Identity, error) {
	builder := app.executionBuilder.Create()
	if identity.IsIdentity() {
		identity := identity.Identity()
		exec, err := app.identityApp.Execute(identity, stack)
		if err != nil {
			return nil, err
		}

		builder.WithIdentity(exec)
	}

	if identity.IsDashboard() {
		dashboard := identity.Dashboard()
		exec, err := app.dashboardApp.Execute(dashboard, stack)
		if err != nil {
			return nil, err
		}

		builder.WithDashboard(exec)
	}

	if identity.IsConnections() {
		connections := identity.Connections()
		exec, err := app.connectionApp.Execute(connections, stack)
		if err != nil {
			return nil, err
		}

		builder.WithConnections(exec)
	}

	if identity.IsProfile() {
		profile := identity.Profile()
		exec, err := app.profileApp.Execute(profile, stack)
		if err != nil {
			return nil, err
		}

		builder.WithProfile(exec)
	}

	if identity.IsShares() {
		shares := identity.Shares()
		exec, err := app.sharesApp.Execute(shares, stack)
		if err != nil {
			return nil, err
		}

		builder.WithShares(exec)
	}

	if identity.IsEncryptor() {
		encryptor := identity.Encryptor()
		exec, err := app.encryptorApp.Execute(encryptor, stack)
		if err != nil {
			return nil, err
		}

		builder.WithEncryptor(exec)
	}

	if identity.IsSigner() {
		signer := identity.Signer()
		exec, err := app.signerApp.Execute(signer, stack)
		if err != nil {
			return nil, err
		}

		builder.WithSigner(exec)
	}

	return builder.Now()
}
