package table

import (
	"errors"

	"github.com/gleisonbs/poker-golang/game/deck"
)

var INVALID_SEAT_NUMBER_ERROR = errors.New("Seats must be a value between 2 and 12")

type Table struct {
	Seats int
	Deck  deck.Deck
}

func New(seats int) (Table, error) {
	table := Table{}
	if seats < 2 || seats > 12 {
		return table, INVALID_SEAT_NUMBER_ERROR
	}

	newDeck := deck.New()
	table = Table{Seats: seats, Deck: newDeck}
	return table, nil
}
