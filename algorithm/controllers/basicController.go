package controllers

import (
	"darkorbitbot/scaning"
	"log"

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

func (c *BasicController) Run(foundMap map[string][]scaning.FoundObject) bool {
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
	//robotgo.MouseSleep = 100
	//todo 1.66 коэффицент при движении мыши на втором экране
	robotgo.Move(int(float64(found.ClickX)/float64(1.66)), found.ClickY)
	robotgo.Click()
	robotgo.MilliSleep(100)
}
