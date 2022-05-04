package algorithm

import "log"

type bot struct {
	currentState *state
	inited       bool
}

func NewBot() bot {
	return bot{&notInitedState, false}
}

func (b *bot) Inited() {
	b.inited = true
	log.Println("Bot inited")
}
