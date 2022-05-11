package algorithm

import (
	"darkorbitbot/algorithm/task"
	"darkorbitbot/scaning"
	"image"
	"image/png"
	"log"
	"math"
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

var run bool = true
var needSave bool = false

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

	teamplates := scaning.GetTeamplates()
	foundChannel := make(chan scaning.FoundObject, 30)

	for run {
		foundMap := make(map[string][]scaning.FoundObject)
		scaner.Scan(teamplates, foundChannel)

		res := int(math.Min(10, float64(len(foundChannel))))
		for i := 0; i < res; i++ {
			x, ok := <-foundChannel
			if !ok {
				break
			}
			foundMap[x.Name] = append(foundMap[x.Name], x)
		}

		log.Println(foundMap)
		// if needSave {
		// 	for _, find := range foundMap {
		// 		for _, found := range find {
		// 			x, y := found.BottomRight()
		// 			scaning.Rect(found.PosX, found.PosY, x, y, scan, found.GetColor())
		// 		}
		// 	}
		// 	saveImageOut(scan)
		// }

		b.Run(foundMap)
	}
}
