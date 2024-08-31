package applications

import (
	"testing"

	"github.com/google/uuid"
	applications_connections "github.com/steve-care-software/steve/applications/connections"
	"github.com/steve-care-software/steve/domain/connections"
	"github.com/steve-care-software/steve/domain/connections/links"
	"github.com/steve-care-software/steve/domain/connections/links/contexts"
	"github.com/steve-care-software/steve/domain/queries"
)

func TestApplication_Success(t *testing.T) {

	contexts := contexts.NewContextsForTests([]contexts.Context{
		contexts.NewContextForTests("family"),
		contexts.NewContextForTests("english"),
	})

	son, _ := uuid.NewRandom()
	father, _ := uuid.NewRandom()
	grandFather, _ := uuid.NewRandom()
	grandGrandFather, _ := uuid.NewRandom()
	noWhereID, _ := uuid.NewRandom()
	anotherNoWhereID, _ := uuid.NewRandom()

	connections := connections.NewConnectionsForTests([]connections.Connection{
		connections.NewConnectionForTests(
			son,
			links.NewLinkForTests(contexts, "son - father", false, 1.0),
			father,
		),
		connections.NewConnectionForTests(
			son,
			links.NewLinkForTests(contexts, "son - grand-father", false, 1.0),
			grandFather,
		),
		connections.NewConnectionForTests(
			son,
			links.NewLinkForTests(contexts, "son - grand-grand father", false, 1.0),
			grandGrandFather,
		),
		connections.NewConnectionForTests(
			father,
			links.NewLinkForTests(contexts, "father - grandfather", false, 1.0),
			grandFather,
		),
		connections.NewConnectionForTests(
			father,
			links.NewLinkForTests(contexts, "father - great-grand-father", false, 1.0),
			grandGrandFather,
		),
		connections.NewConnectionForTests(
			grandFather,
			links.NewLinkForTests(contexts, "grand-father - great-grand-father", false, 1.0),
			grandGrandFather,
		),
		connections.NewConnectionForTests(
			son,
			links.NewLinkForTests(contexts, "son - nowhere", false, 1.0),
			noWhereID,
		),
		connections.NewConnectionForTests(
			anotherNoWhereID,
			links.NewLinkForTests(contexts, "another no-where - great-grand-father", false, 1.0),
			grandGrandFather,
		),
	})

	inMemoryApp, err := applications_connections.NewInMemoryBuilder().
		Create().
		WithConnections(connections).
		Now()

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	query := queries.NewQueryForTests(son, grandGrandFather)

	application := NewApplication(inMemoryApp)
	retRoute, err := application.Route(query)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	possibilities := retRoute.Possibilities()
	if len(possibilities) != 4 {
		t.Errorf("the Route was expected to contain %d possibilities, %d returned", 4, len(possibilities))
		return
	}

}
