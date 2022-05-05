package task

import (
	"darkorbitbot/algorithm/controllers"
	"darkorbitbot/scaning"
)

/**
Таск - это по сути текущая цель бота
Как пример:
	-фарм(включает и бои и сбор коробок и прокачку)
	-сбор коробок(чисто безопасный сбор без боёв)
	-патрулирование низов
	-пвп(самостоятельное)
	-следование за ведущим(живым игроком)
	-фарм кредитов через скайлэб(чисто прокачка скайлеба, продажа ресурсов, отправка)
	Можно продолжать очень много, например, прохождение ворот, хотя и это уйдет в фарм
Имеет набор контроллеров, которые выполняются последовательно(если передают управление дальше)
Например, контроллер боя выше контроллера сбора коробок и если мы деремся,
то не передаем управление сборке коробок, потому камон, мы пиздимся, какие коробки вообще
*/
type Task struct {
	controllers []controllers.IController
}

func (t *Task) Run(foundMap map[string][]*scaning.FoundObject) {
	for _, c := range t.controllers {
		breaked := c.Run(foundMap)
		if breaked {
			break
		}
	}
}

func (t *Task) GetControllers() []controllers.IController {
	return t.controllers
}

type ITask interface {
	GetControllers() []controllers.IController
}
