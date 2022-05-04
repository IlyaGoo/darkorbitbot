package algorithm

import (
	"darkorbitbot/algorithm/states"
	"log"
)

type bot struct {
	currentState states.IState
	inited       bool
}

func NewBot() bot {
	state := states.GetNotInitedState()
	return bot{state, false}
}

func (b *bot) Inited() {
	b.inited = true
	log.Println("Bot inited")
}
