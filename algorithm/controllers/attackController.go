package controllers

import "darkorbitbot/scaning"

/** Контроллер, атакующий всех ботов подряд
Упрощенный алгоритм:
-Ведем ли мы сейчас бой
	-Если да, то обрабатываем ведение боя, не передаем управление следующему контроллеру
	-Если нет, то ищем, нет ли врагов, с которыми можно начать бой
		-Нашли врага, атакуем, не передаем управление следующему контроллеру
		-Не нашли, переали управление следующему контроллеру
*/
type AttackController struct {
}

func NewAttackController() *AttackController {
	return &AttackController{}
}

func (c *AttackController) Run(foundMap map[string][]scaning.FoundObject) bool { //todo
	return false
}
