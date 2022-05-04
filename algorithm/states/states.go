package states

type IState interface {
	GetName() string
	Do()
}

type state struct {
	name string
}

func (s *state) GetName() string {
	return s.name
}
