package scaning

type foundObject struct {
	t    *teamplate
	PosX int
	PosY int
}

var dX = -1280
var dY = 0

func (o *foundObject) GetClickPosition() (int, int) {
	return o.PosX + o.t.image.Rect.Dx()/2 + dX, o.PosY + o.t.image.Rect.Dy()/2 + dY
}
