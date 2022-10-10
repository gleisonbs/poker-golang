package table

import (
	"errors"

	"github.com/gleisonbs/poker-golang/game/deck"
	"github.com/gleisonbs/poker-golang/game/player"
	"github.com/gleisonbs/poker-golang/game/round"
)

var INVALID_SEAT_NUMBER_ERROR = errors.New("Seats must be a value between 2 and 12")

type TableState struct {
	BetValue  int
	callValue int
}

type Blinds struct {
	Small int
	Big   int
}

type Table struct {
	MaxSeats    int
	Players     []player.Player
	Deck        deck.Deck
	Pot         int
	State       TableState
	Dealer      int
	Blinds      Blinds
	Round       round.Round
	PlayerToAct int
}

func New(maxSeats int) (Table, error) {
	table := Table{}
	if maxSeats < 2 || maxSeats > 12 {
		return table, INVALID_SEAT_NUMBER_ERROR
	}

	newDeck := deck.New()
	table = Table{MaxSeats: maxSeats, Deck: newDeck}
	return table, nil
}

func (t *Table) CollectBlinds() {
	if len(t.Players) < 2 {
		return
	}

	if len(t.Players) == 2 {
		bigBlindPlayerIndex := (t.Dealer + 1) % len(t.Players)
		t.Pot += t.Players[bigBlindPlayerIndex].PostBlind(t.Blinds.Big)
		return
	}

	smallBlindPlayerIndex := (t.Dealer + 1) % len(t.Players)
	bigBlindPlayerIndex := (t.Dealer + 2) % len(t.Players)
	t.Pot += t.Players[bigBlindPlayerIndex].PostBlind(t.Blinds.Big)
	t.Pot += t.Players[smallBlindPlayerIndex].PostBlind(t.Blinds.Small)
}

func (t *Table) GetNextPlayerToAct() {

}

func (t *Table) DealStartingHands() {
	for i, _ := range t.Players {
		playerHand := t.Deck.Deal(2)
		t.Players[i].Hand = playerHand
	}
}

func (t *Table) SitPlayers() {
	for i := 0; i < t.MaxSeats; i++ {
		newPlayer := player.Player{Hand: nil, Stack: 2500}
		t.Players = append(t.Players, newPlayer)
	}

	t.PlayerToAct = len(t.Players) - 1
	if len(t.Players) > 4 {
		t.PlayerToAct = 3
	}
}

func (t *Table) UpdateDealerPosition() {
	t.Dealer = (t.Dealer + 1) % len(t.Players)
}

func (t *Table) Start() {
	t.Deck = deck.New()
	t.Dealer = 0
	t.Blinds = Blinds{Small: 5, Big: 10}
	t.Round.CallValue = t.Blinds.Big
	t.Deck.Shuffle()

	// for _, p := range t.Players {
	// 	playerAction, value := p.GetAction()
	// 	switch playerAction {
	// 	case player.Actions.Bet:
	// 		t.State.BetValue += value
	// 		t.Pot += value
	// 	}
	// }
}
