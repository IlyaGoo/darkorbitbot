package main

import (
	"darkorbitbot/algorithm"
	"darkorbitbot/config"
	"fmt"

	hook "github.com/robotn/gohook"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Get()
	fmt.Println("Bot started")
	algo := algorithm.NewAlgorithm(cfg)

	go listenEvents(&algo)
	algo.Run()
}

func listenEvents(algo *algorithm.Algorithm) {
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		logrus.Info("Stopped")
		algo.Stop()
		hook.End()
	})

	s := hook.Start()
	<-hook.Process(s)
}
