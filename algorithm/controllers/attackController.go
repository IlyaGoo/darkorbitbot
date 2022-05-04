package controllers

/** Контроллер, атакующий всех ботов подряд */
type AttackController struct {
}

func NewAttackController() *AttackController {
	return &AttackController{}
}

func (c *AttackController) Run() bool { //todo
	return false
}
