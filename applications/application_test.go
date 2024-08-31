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

	context := contexts.NewContextForTests("family")

	son, _ := uuid.NewRandom()
	father, _ := uuid.NewRandom()
	grandFather, _ := uuid.NewRandom()
	grandGrandFather, _ := uuid.NewRandom()
	noWhereID, _ := uuid.NewRandom()
	anotherNoWhereID, _ := uuid.NewRandom()

	connections := connections.NewConnectionsForTests([]connections.Connection{
		connections.NewConnectionForTests(
			son,
			links.NewLinkWithReverseForTests(context, "son - father", 1.0, "father - son"),
			father,
		),
		connections.NewConnectionForTests(
			son,
			links.NewLinkWithReverseForTests(context, "son - grand-father", 1.0, "grand-father - son"),
			grandFather,
		),
		connections.NewConnectionForTests(
			son,
			links.NewLinkWithReverseForTests(context, "son - grand-grand father", 1.0, "grand-grand father - son"),
			grandGrandFather,
		),
		connections.NewConnectionForTests(
			father,
			links.NewLinkWithReverseForTests(context, "father - grandfather", 1.0, "grandfather, father"),
			grandFather,
		),
		connections.NewConnectionForTests(
			father,
			links.NewLinkWithReverseForTests(context, "father - great-grand-father", 1.0, "great-grand-father - father"),
			grandGrandFather,
		),
		connections.NewConnectionForTests(
			grandFather,
			links.NewLinkWithReverseForTests(context, "grand-father - great-grand-father", 1.0, "great-grand-father - grand-father"),
			grandGrandFather,
		),
		connections.NewConnectionForTests(
			son,
			links.NewLinkWithReverseForTests(context, "son - nowhere", 1.0, "nowhere - son"),
			noWhereID,
		),
		connections.NewConnectionForTests(
			anotherNoWhereID,
			links.NewLinkWithReverseForTests(context, "another no-where - great-grand-father", 1.0, "great-grand-father - another no-where"),
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

	query := queries.NewQueryForTests(grandGrandFather, son)

	application := NewApplication(inMemoryApp)
	retRoute, err := application.Route(query)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	possibilities := retRoute.Possibilities()
	if len(possibilities) != 5 {
		t.Errorf("the Route was expected to contain %d possibilities, %d returned", 5, len(possibilities))
		return
	}

}
