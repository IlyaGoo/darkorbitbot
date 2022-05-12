package controllers

import (
	"darkorbitbot/scaning"
	"fmt"

	"github.com/go-vgo/robotgo"
)

/**
 Контроллер с базовыми потребностями
	Следит, чтобы бот не сдох и если мало хп, то дает по съебам
	Докупает беприпасы(если нет соответсвующего чипа)
	Следит за шаклой врагов на карте на картах, где это вохможно и если он слаб, то тоже дает по съебам
*/
type BasicController struct {
	initInterfaceed bool
}

func NewBasicController() *BasicController {
	return &BasicController{}
}

/** Инициализировать интерфейс */
func (c *BasicController) InitInterface(foundMap map[string][]scaning.FoundObject) {
	c.initInterfaceed = InitInterface(foundMap)
}

func (c *BasicController) Run(foundMap map[string][]scaning.FoundObject) bool {
	// if 1 == 1 {
	// 	x, y := robotgo.GetMousePos()
	// 	fmt.Println("pos: ", x, y)
	// 	return true
	// }

	if len(foundMap["map"]) > 0 {
		return false
	} else if len(foundMap["connect"]) > 0 {
		fmt.Println("Подключаюсь")
		found := foundMap["connect"][0]
		MooveAndClick(found)
		robotgo.MilliSleep(3000)
		return true
	} else if len(foundMap["close"]) > 0 {
		fmt.Println("Закрываю рекламу")
		found := foundMap["close"][0]
		MooveAndClick(found)
		return true
	} else if !c.initInterfaceed {
		fmt.Println("Начинаю инициализацию")
		c.InitInterface(foundMap)
		return true
	} else if len(foundMap["cargo"]) > 0 && len(foundMap["trade"]) > 0 {
		found := foundMap["trade"][0]
		MooveAndClick(found)
		robotgo.MilliSleep(2000)
		return true
	} else if len(foundMap["sell"]) > 0 {
		found := foundMap["sell"][0]
		MooveAndClick(found)
		robotgo.MilliSleep(2000)
		return true
	} else {
		//log.Println("Waiting for initInterface")
	}
	return false
}

func MooveAndClick(found scaning.FoundObject) {
	//robotgo.MouseSleep = 100
	//todo 1.66 коэффицент положения мыши на втором экране
	robotgo.Move(int(float64(found.ClickX)/float64(1.66)), found.ClickY)
	robotgo.Click()
	robotgo.MilliSleep(100)
}
