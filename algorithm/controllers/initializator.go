package controllers

import (
	"darkorbitbot/scaning"
	"fmt"

	"github.com/go-vgo/robotgo"
)

/** Инициализировать интерфейс
Обнаружить все необходимые табло
Переместить на стандартные позиции
Закрыть ненужые табло
*/
func InitInterface(foundMap map[string][]scaning.FoundObject) bool {
	if len(foundMap["minimap_icon"]) > 0 {
		fmt.Println("Перетаскиваю миникарту")
		found := foundMap["minimap_icon"][0]
		startX := int(float64(found.ClickX)/float64(1.66)) + 50
		robotgo.Move(startX, found.ClickY)
		robotgo.MilliSleep(100)
		x := 1200 / 1.66
		robotgo.Toggle("left")
		robotgo.MilliSleep(1000)
		robotgo.Drag(int(x), 800)
		robotgo.MilliSleep(1000)
		robotgo.Toggle("left", "up")

		if len(foundMap["user_icon"]) > 0 {
			fmt.Println("Перетаскиваю миникарту")
			found := foundMap["user_icon"][0]
			startX := int(float64(found.ClickX)/float64(1.66)) + 50
			robotgo.Move(startX, found.ClickY)
			robotgo.MilliSleep(100)
			x := 1200 / 1.66
			robotgo.Toggle("left")
			robotgo.MilliSleep(1000)
			robotgo.Drag(int(x), 800)
			robotgo.MilliSleep(1000)
			robotgo.Toggle("left", "up")
			return true
		}
	}
	return false
}
