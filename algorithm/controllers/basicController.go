package controllers

/**
 Контроллер с базовыми потребностями
	Следит, чтобы бот не сдох и если мало хп, то дает по съебам
	Докупает беприпасы(если нет соответсвующего чипа)
	Следит за шаклой врагов на карте на картах, где это вохможно и если он слаб, то тоже дает по съебам
*/
type BasicController struct {
}

func NewBasicController() *BasicController {
	return &BasicController{}
}

func (c *BasicController) Run() bool { //todo
	return false
}
