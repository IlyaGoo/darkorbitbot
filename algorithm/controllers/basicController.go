package controllers

import (
	"darkorbitbot/scaning"
	"log"

	"darkorbitbot/config"

	"github.com/go-vgo/robotgo"
)

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

func (c *BasicController) Run(foundMap map[string][]scaning.FoundObject) bool { //todo
	if len(foundMap["map"]) > 0 {
		return false
	} else if len(foundMap["connect"]) > 0 {
		found := foundMap["connect"][0]
		MooveAndClick(found)
		return true
	} else if len(foundMap["close"]) > 0 {
		found := foundMap["close"][0]
		MooveAndClick(found)
		return true
	} else {
		log.Println("Waiting for init")
	}
	return false
}

func MooveAndClick(found scaning.FoundObject) {
	robotgo.Move(config.Get().XDif+found.ClickX, config.Get().YDif+found.ClickY)
	robotgo.MilliSleep(100)
	robotgo.Click()
}
