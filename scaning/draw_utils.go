package scaning

import (
	"image"
	"image/color"
)

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}

// HLine draws a horizontal line
func HLine(x1, y, x2 int, imgRGBA *image.RGBA, color color.RGBA) {
	for ; x1 <= x2; x1++ {
		imgRGBA.Set(x1, y, color)
	}
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int, imgRGBA *image.RGBA, color color.RGBA) {
	for ; y1 <= y2; y1++ {
		imgRGBA.Set(x, y1, color)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1 int, y1 int, x2 int, y2 int, imgRGBA *image.RGBA, color color.RGBA) {
	HLine(x1, y1, x2, imgRGBA, color)
	HLine(x1, y2, x2, imgRGBA, color)
	VLine(x1, y1, y2, imgRGBA, color)
	VLine(x2, y1, y2, imgRGBA, color)
}
