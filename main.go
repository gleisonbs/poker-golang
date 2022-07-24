package main

import (
	"fmt"

	"github.com/gleisonbs/poker-golang/game/deck"
	"github.com/gleisonbs/poker-golang/game/table"
)

func main() {
	newDeck := deck.New()
	fmt.Println(newDeck)
	newDeck.Shuffle()
	fmt.Println()
	fmt.Println(newDeck)
	fmt.Println()
	cards := newDeck.Deal(2)
	fmt.Println(cards)
	fmt.Println(newDeck)

	table.New(6)
}
