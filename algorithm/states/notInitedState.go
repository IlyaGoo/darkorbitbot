package states

/** Первоночальное состояние, при включении бота */
type NotInitedState struct {
	state
}

func GetNotInitedState() NotInitedState {
	return NotInitedState{state{"Not inited"}}
}

func (s NotInitedState) Do() {
}

func (s NotInitedState) GetName() string {
	return s.name
}
