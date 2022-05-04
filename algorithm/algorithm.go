package algorithm

import (
	"darkorbitbot/scaning"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"os/signal"

	"github.com/go-vgo/robotgo"
)

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
	b := NewBot()

	for run {
		foundMap, scan := scaner.Scan()
		log.Println(foundMap)
		if needSave {
			saveImageOut(scan)
		}

		if b.currentState.GetName() == "" {
			if len(foundMap["map"]) > 0 {
				b.Inited()
			} else if len(foundMap["connect"]) > 0 {
				foundRef := foundMap["connect"][0]
				x, y := foundRef.GetClickPosition()
				robotgo.Move(x, y)
				robotgo.MilliSleep(100)
				robotgo.Click()
			} else {
				log.Println("Waiting for init")
			}
		}

	}
}
