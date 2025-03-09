package player

import (
	"fmt"

	"github.com/gleisonbs/poker-golang/game/deck"
	"github.com/gleisonbs/poker-golang/game/round"
)

type Player struct {
	Hand         []deck.Card
	Stack        int
	MoneyInRound int
}

func (p *Player) Call(value int) int {
	p.Stack -= value
	p.MoneyInRound += value
	return value
}

func (p *Player) Bet(value int) int {
	p.Stack -= value
	p.MoneyInRound += value
	return value
}

func (p *Player) PostBlind(value int) int {
	p.Stack -= value
	p.MoneyInRound += value
	return value
}

func (p *Player) GetAction(round round.Round) (Action, int) {
	value := 10
	allowedActions := p.GetAllowedActions(round)
	fmt.Println(allowedActions)
	p.Bet(value)
	return Actions.Bet, value
}

func (p *Player) GetAllowedActions(round round.Round) []Action {
	if p.MoneyInRound == round.CallValue {
		return []Action{Actions.Fold, Actions.Check, Actions.Bet, Actions.AllIn}
	}

	if p.MoneyInRound < round.CallValue {
		return []Action{Actions.Fold, Actions.Call, Actions.Bet, Actions.AllIn}
	}
	
	return []Action{Actions.Fold, Actions.Call, Actions.Check, Actions.AllIn}
}
