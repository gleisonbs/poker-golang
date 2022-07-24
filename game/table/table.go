package table

import (
	"github.com/gleisonbs/poker-golang/game/deck"
)

type Table struct {
	Seats int
	Deck  deck.Deck
}

func New(seats int) Table {
	newDeck := deck.New()
	table := Table{Seats: seats, Deck: newDeck}
	return table
}
