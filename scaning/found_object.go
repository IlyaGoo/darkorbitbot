package scaning

import (
	"image/color"
)

type FoundObject struct {
	Name   string
	Color  color.RGBA
	ClickX int
	ClickY int
	PosX   int
	PosY   int
}

func (o FoundObject) GetColor() color.RGBA {
	return o.Color
}
