package player

import (
	"github.com/gleisonbs/poker-golang/game/deck"
)

type Player struct {
	Hand  []deck.Card
	Stack int
}

func (p *Player) Call(value int) int {
	p.Stack -= value
	return value
}

func (p *Player) Bet(value int) int {
	p.Stack -= value
	return value
}

func (p *Player) PostBlind(value int) int {
	p.Stack -= value
	return value
}

func (p *Player) GetAction() (Action, int) {
	value := 10
	p.Bet(value)
	return Actions.Bet, value
}

func allowedActions(p *Player) []Action {
	return []Action{Actions.Call, Actions.Check}
}
