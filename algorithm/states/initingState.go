package states

/** Второе состояние, при инициализации */
type InitingState struct {
	state
}

func GetInitingState() InitingState {
	return InitingState{state{"Inited"}}
}

func (s InitingState) Do() {
}

func (s InitingState) GetName() string {
	return s.name
}
