package algorithm

import (
	"darkorbitbot/algorithm/task"
	"darkorbitbot/scaning"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"os/signal"
)

/** По сути единица бота
Хранит в себе сущность корабля
Хранит текущий таск
Выполняет его
*/
type Algorithm struct {
}

func saveImage(scaned *image.RGBA, counter int) {
	fmt.Println(counter)
}

var run bool = true
var needSave bool = true

func listenForShotDown(ch <-chan os.Signal) {
	<-ch
	run = false
}

func saveImageOut(scan *image.RGBA) {
	f, err := os.Create("outimage.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = png.Encode(f, scan)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *Algorithm) Run() {
	scaner := scaning.NewImageScaner()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShotDown(ch)
	b := task.NewFarmTask()

	for run {
		foundMap, scan := scaner.Scan()
		log.Println(foundMap)
		if needSave {
			saveImageOut(scan)
		}

		b.Run(foundMap)
	}
}
