package graphs

import (
	"testing"

	"github.com/google/uuid"
	applications_connections "github.com/steve-care-software/steve/engine/applications/chains/interpreters/graphs/connections"
	"github.com/steve-care-software/steve/engine/domain/graphs/connections"
	"github.com/steve-care-software/steve/engine/domain/graphs/connections/links"
)

func TestApplication_Success(t *testing.T) {
	son, _ := uuid.NewRandom()
	father, _ := uuid.NewRandom()
	grandFather, _ := uuid.NewRandom()
	grandGrandFather, _ := uuid.NewRandom()
	noWhereID, _ := uuid.NewRandom()
	anotherNoWhereID, _ := uuid.NewRandom()

	connections := connections.NewConnectionsForTests([]connections.Connection{
		connections.NewConnectionForTests(
			son,
			links.NewLinkWithReverseForTests("father", "son"),
			father,
		),
		connections.NewConnectionForTests(
			son,
			links.NewLinkWithReverseForTests("grand-father", "grand-son"),
			grandFather,
		),
		connections.NewConnectionForTests(
			son,
			links.NewLinkWithReverseForTests("grand-grand-father", "grand-grand-son"),
			grandGrandFather,
		),
		connections.NewConnectionForTests(
			father,
			links.NewLinkWithReverseForTests("father", "son"),
			grandFather,
		),
		connections.NewConnectionForTests(
			father,
			links.NewLinkWithReverseForTests("grand-father", "grand-son"),
			grandGrandFather,
		),
		connections.NewConnectionForTests(
			grandFather,
			links.NewLinkWithReverseForTests("father", "son"),
			grandGrandFather,
		),
		connections.NewConnectionForTests(
			son,
			links.NewLinkWithReverseForTests("son - nowhere", "nowhere - son"),
			noWhereID,
		),
		connections.NewConnectionForTests(
			anotherNoWhereID,
			links.NewLinkWithReverseForTests("another no-where - great-grand-father", "great-grand-father - another no-where"),
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

	application := NewApplication(inMemoryApp)
	retRoute, err := application.Route(grandGrandFather, son)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	possibilities := retRoute.Possibilities()
	if len(possibilities) != 5 {
		t.Errorf("the Route was expected to contain %d possibilities, %d returned", 5, len(possibilities))
		return
	}

	list, err := application.LinkIntersect([]string{"son", "grand-son", "grand-grand-son"})
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if len(list) != 1 {
		t.Errorf("the list was expected to contain %d edge, %d returned", 1, len(list))
		return
	}

	if list[0].String() != son.String() {
		t.Errorf("the son (id: %s) was expected to be contained in the list, %s returned", son.String(), list[0].String())
		return
	}

}
