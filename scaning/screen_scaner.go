package scaning

import (
	"image"

	"github.com/kbinani/screenshot"
	"github.com/vcaesar/gcv"
)

type ImageScaner struct {
	bounds image.Rectangle
}

func NewImageScaner() ImageScaner {
	n := screenshot.NumActiveDisplays()
	bounds := screenshot.GetDisplayBounds(n - 1)
	return ImageScaner{bounds}
}

func gcvCheckTeampleate(canvas image.Image, teamplate *teamplate, ch chan FoundObject) {
	mm := gcv.FindAllImg(teamplate.Image, canvas)

	for _, f := range mm {
		newObject := FoundObject{
			teamplate.name,
			teamplate.color,
			f.Middle.X,
			f.Middle.Y,
			f.TopLeft.X,
			f.TopLeft.Y,
		}
		ch <- newObject
	}
}

func (s *ImageScaner) Scan(teamplates []*teamplate, ch chan FoundObject) image.Image {
	imgRGBA, err := screenshot.CaptureRect(s.bounds)
	if err != nil {
		panic(err)
	}

	canvas := image.Image(imgRGBA)
	for _, t := range teamplates {
		go gcvCheckTeampleate(canvas, t, ch)
	}

	return canvas
}
