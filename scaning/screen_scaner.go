package scaning

import (
	"context"
	"image"
	"image/draw"
	"log"

	"github.com/andreyvit/locateimage"
	"github.com/kbinani/screenshot"
)

type ImageScaner struct {
	bounds image.Rectangle
}

func NewImageScaner() ImageScaner {
	n := screenshot.NumActiveDisplays()
	bounds := screenshot.GetDisplayBounds(n - 1)
	return ImageScaner{bounds}
}

func checkTeampleate(canvas *image.RGBA, fObject *teamplate) []*FoundObject {
	mm, err := locateimage.All(context.Background(), canvas, fObject.image, fObject.recommendedTolerance)
	if err != nil {
		log.Fatal(err)
	}

	for _, find := range mm {
		Rect(find.Rect.Min.X, find.Rect.Min.Y, find.Rect.Max.X, find.Rect.Max.Y, canvas, fObject.color)
	}

	foundObjects := make([]*FoundObject, len(mm))
	for i, e := range mm {
		newObject := FoundObject{
			fObject,
			e.Rect.Min.X,
			e.Rect.Min.Y,
		}
		foundObjects[i] = &newObject
	}
	return foundObjects
}

func (s *ImageScaner) Scan() (map[string][]*FoundObject, *image.RGBA) {
	imgRGBA, err := screenshot.CaptureRect(s.bounds)
	if err != nil {
		panic(err)
	}

	//kek9 := imaging.AdjustSaturation(imgRGBA, 80)
	//kek10 := imaging.AdjustContrast(kek9, 100)
	img0 := image.Image(imgRGBA)
	b0 := img0.Bounds()
	canvas := image.NewRGBA(image.Rect(0, 0, b0.Dx(), b0.Dy()))
	draw.Draw(canvas, canvas.Bounds(), img0, b0.Min, draw.Src)

	teamplates := getTeamplates()

	m := make(map[string][]*FoundObject)

	for _, t := range teamplates {
		m[t.name] = checkTeampleate(canvas, t)
	}

	return m, canvas
}
