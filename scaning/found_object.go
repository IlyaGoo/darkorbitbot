package scaning

import "image/color"

type FoundObject struct {
	Name  string
	Color color.RGBA
	Dx    int
	Dy    int
	PosX  int
	PosY  int
}

var dX = -1280
var dY = 0

func (o *FoundObject) GetClickPosition() (int, int) {
	return o.PosX + o.Dx/2 + dX, o.PosY + o.Dy/2 + dY
}

func (o FoundObject) BottomRight() (int, int) {
	return o.PosX + o.Dx, o.PosY + o.Dy
}

func (o FoundObject) GetColor() color.RGBA {
	return o.Color
}
