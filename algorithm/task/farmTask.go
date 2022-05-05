package task

import "darkorbitbot/algorithm/controllers"

type FarmTask struct {
	*Task
}

func NewFarmTask() FarmTask {
	controllers := []controllers.IController{
		controllers.NewBasicController(),
	}
	return FarmTask{&Task{controllers}}
}
