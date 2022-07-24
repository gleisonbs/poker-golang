package deck

import (
	"testing"

	"github.com/gleisonbs/poker-golang/game/deck"
)

func TestNewDeck(t *testing.T) {
	newDeck := deck.New()
	deckSize := 52
	if len(newDeck.Cards) != deckSize {
		t.Errorf("Deck has %v cards, expected %v", len(newDeck.Cards), deckSize)
	}
}

func TestDealFromDeck(t *testing.T) {
	cardsToDeal := 2
	deckSize := 52
	newDeck := deck.New()
	cards := newDeck.Deal(cardsToDeal)
	if (len(cards)) != cardsToDeal {
		t.Errorf("Cards dealt is %v, expected %v", len(cards), cardsToDeal)
	}

	if (len(newDeck.Cards)) != deckSize-cardsToDeal {
		t.Errorf("Deck size is %v, expected %v", len(newDeck.Cards), deckSize-cardsToDeal)
	}
}
