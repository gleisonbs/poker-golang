package table

import (
	"testing"

	"github.com/gleisonbs/poker-golang/game/table"
)

func TestHandlesCall(t *testing.T) {
	nSeats := 6
	newTable, _ := table.New(nSeats)
	newTable.SitPlayers()
	newTable.DealStartingHands()

	if len(newTable.Players) != nSeats {
		t.Errorf("Number of seats is %v, expected %v", len(newTable.Players), nSeats)
	}
}
