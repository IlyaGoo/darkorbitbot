package algorithm

import (
	"darkorbitbot/algorithm/task"
	"darkorbitbot/config"
	"darkorbitbot/scaning"
	"image"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"
)

/** По сути единица бота
Хранит в себе сущность корабля
Хранит текущий таск
Выполняет его
*/
type Algorithm struct {
	config config.Config
}

func NewAlgorithm(config config.Config) Algorithm {
	return Algorithm{config}
}

func (a *Algorithm) Stop() {
	run = false
}

var run bool = true

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

	b := task.NewFarmTask()

	teamplates := scaning.GetTeamplates()
	foundChannel := make(chan scaning.FoundObject, 30)

	for run {
		foundMap := make(map[string][]scaning.FoundObject)
		scan := scaner.Scan(teamplates, foundChannel)

		res := int(math.Min(10, float64(len(foundChannel))))
		for i := 0; i < res; i++ {
			x, ok := <-foundChannel
			if !ok {
				break
			}
			foundMap[x.Name] = append(foundMap[x.Name], x)
		}

		if a.config.SaveScreen {
			b := scan.Bounds()
			m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
			draw.Draw(m, m.Bounds(), scan, b.Min, draw.Src)

			for _, find := range foundMap {
				for _, found := range find {
					scaning.Rect(found.PosX, found.PosY, found.ClickX, found.ClickY, m, found.GetColor())
				}
			}
			saveImageOut(m)
		}

		b.Run(foundMap)
	}
}
