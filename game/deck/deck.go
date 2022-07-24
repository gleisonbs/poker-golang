package deck

import (
	"math/rand"
	"time"
)

const (
	Clubs    string = "c"
	Diamonds        = "d"
	Hearts          = "h"
	Spades          = "s"
)

type Card struct {
	Suit  string
	Value string
}

type Deck struct {
	Cards []Card
}

func (d *Deck) AddCard(c Card) {
	d.Cards = append(d.Cards, c)
}

func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Card, len(d.Cards))
	perm := r.Perm(len(d.Cards))
	for i, randIndex := range perm {
		ret[i] = d.Cards[randIndex]
	}
	d.Cards = ret
}

func (d *Deck) Deal(n int) []Card {
	cards := d.Cards[len(d.Cards)-n : len(d.Cards)]
	d.Cards = d.Cards[:len(d.Cards)-n]
	return cards
}

func New() Deck {
	suits := [4]string{Clubs, Diamonds, Hearts, Spades}
	values := [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K"}

	cards := []Card{}
	deck := Deck{Cards: cards}

	for _, suit := range suits {
		for _, value := range values {
			deck.AddCard(Card{Value: value, Suit: suit})
		}
	}

	return deck
}
