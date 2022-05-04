package task

import "darkorbitbot/algorithm/controllers"

type FarmTask struct {
	Task
}

func NewFarmTask() FarmTask {
	controllers := []controllers.IController{
		controllers.NewAttackController(),
	}
	return FarmTask{Task{controllers}}
}
