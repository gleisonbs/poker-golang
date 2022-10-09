package table

import (
	"testing"

	"github.com/gleisonbs/poker-golang/game/table"
)

func TestNewTable(t *testing.T) {
	nSeats := 9
	newTable, _ := table.New(nSeats)
	newTable.SitPlayers()

	if len(newTable.Players) != nSeats {
		t.Errorf("Number of seats is %v, expected %v", len(newTable.Players), nSeats)
	}

	if len(newTable.Deck.Cards) != 52 {
		t.Errorf("Deck size is %v, expected %v", len(newTable.Deck.Cards), 52)
	}

	nSeats = 6
	newTable, _ = table.New(nSeats)
	newTable.SitPlayers()

	if len(newTable.Players) != nSeats {
		t.Errorf("Number of seats is %v, expected %v", len(newTable.Players), nSeats)
	}
}

func TestNewTableInvalidSeats(t *testing.T) {
	nSeats := -2
	_, err := table.New(nSeats)

	if err != table.INVALID_SEAT_NUMBER_ERROR {
		t.Errorf("Expected to throw error for invalid seat number %v, not throw", nSeats)
	}

	nSeats = 0
	_, err = table.New(nSeats)

	if err != table.INVALID_SEAT_NUMBER_ERROR {
		t.Errorf("Expected to throw error for invalid seat number %v, not throw", nSeats)
	}

	nSeats = 1
	_, err = table.New(nSeats)

	if err != table.INVALID_SEAT_NUMBER_ERROR {
		t.Errorf("Expected to throw error for invalid seat number %v, not throw", nSeats)
	}

	nSeats = 13
	_, err = table.New(nSeats)

	if err != table.INVALID_SEAT_NUMBER_ERROR {
		t.Errorf("Expected to throw error for invalid seat number %v, not throw", nSeats)
	}
}
