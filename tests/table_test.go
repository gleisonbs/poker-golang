package table

import (
	"testing"

	"github.com/gleisonbs/poker-golang/game/table"
)

func TestNewTable(t *testing.T) {
	nSeats := 9
	newTable, _ := table.New(nSeats)
	newTable.Start()
	newTable.SitPlayers()

	if len(newTable.Players) != nSeats {
		t.Errorf("Number of seats is %v, expected %v", len(newTable.Players), nSeats)
	}

	if len(newTable.Deck.Cards) != 52 {
		t.Errorf("Deck size is %v, expected %v", len(newTable.Deck.Cards), 52)
	}

	nSeats = 6
	newTable, _ = table.New(nSeats)
	newTable.Start()
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

func TestTableCorectlyCollectsBlinds(t *testing.T) {
	// WHEN THERE ARE ONLY 2 PLAYERS
	nSeats := 2
	newTable, _ := table.New(nSeats)
	newTable.Start()
	newTable.SitPlayers()
	newTable.CollectBlinds()

	if newTable.Players[0].Stack != 2500 {
		t.Errorf("Expected to not collect Big Blind from dealer")
	}

	if newTable.Players[1].Stack != 2490 {
		t.Errorf("Expected to collect Big Blind from seat number %v", 2)
	}

	// WHEN THERE ARE ONLY 2 PLAYERS AND THE DEALER IS THE SECOND PLAYER
	newTable, _ = table.New(nSeats)
	newTable.Start()
	newTable.SitPlayers()
	newTable.UpdateDealerPosition()
	newTable.CollectBlinds()

	if newTable.Players[0].Stack != 2490 {
		t.Errorf("Expected to collect Big Blind from seat number %v", 1)
	}

	if newTable.Players[1].Stack != 2500 {
		t.Errorf("Expected to not collect Big Blind from dealer")
	}

	// WHEN THERE ARE ONLY 3 PLAYERS
	nSeats = 3
	newTable, _ = table.New(nSeats)
	newTable.Start()
	newTable.SitPlayers()
	newTable.CollectBlinds()

	if newTable.Players[0].Stack != 2500 {
		t.Errorf("Expected to not collect Big Blind from dealer")
	}

	if newTable.Players[1].Stack != 2495 {
		t.Errorf("Expected to collect Small Blind from seat number %v", 2)
	}

	if newTable.Players[2].Stack != 2490 {
		t.Errorf("Expected to collect Big Blind from seat number %v", 3)
	}

	// WHEN THERE ARE ONLY 3 PLAYERS AND THE DEALER IS THE SECOND PLAYER
	nSeats = 3
	newTable, _ = table.New(nSeats)
	newTable.Start()
	newTable.SitPlayers()
	newTable.UpdateDealerPosition()
	newTable.CollectBlinds()

	if newTable.Players[0].Stack != 2490 {
		t.Errorf("Expected to collect Big Blind from seat number %v", 3)
	}

	if newTable.Players[1].Stack != 2500 {
		t.Errorf("Expected to not collect Big Blind from dealer")
	}

	if newTable.Players[2].Stack != 2495 {
		t.Errorf("Expected to collect Small Blind from seat number %v", 2)
	}

	// WHEN THERE ARE ONLY 3 PLAYERS AND THE DEALER IS THE THIRD PLAYER
	nSeats = 3
	newTable, _ = table.New(nSeats)
	newTable.Start()
	newTable.SitPlayers()
	newTable.UpdateDealerPosition()
	newTable.UpdateDealerPosition()
	newTable.CollectBlinds()

	if newTable.Players[0].Stack != 2495 {
		t.Errorf("Expected to collect Small Blind from seat number %v", 2)
	}

	if newTable.Players[1].Stack != 2490 {
		t.Errorf("Expected to collect Big Blind from seat number %v", 3)
	}

	if newTable.Players[2].Stack != 2500 {
		t.Errorf("Expected to not collect Big Blind from dealer")
	}
}
